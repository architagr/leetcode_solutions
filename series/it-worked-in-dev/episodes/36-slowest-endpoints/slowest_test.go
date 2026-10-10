package slowestendpoints

import (
	"fmt"
	"math"
	"math/rand"
	"slices"
	"testing"
)

var paths = func() []string {
	var p []string
	for i := 0; i < 300; i++ {
		p = append(p, fmt.Sprintf("/api/v2/route%03d", i))
	}
	return p
}()

// traffic is n requests with log-normal latency around 40 ms, the shape real
// latency has: most fast, a long slow tail.
func traffic(n int, seed int64) []Req {
	r := rand.New(rand.NewSource(seed))
	out := make([]Req, n)
	for i := range out {
		out[i] = Req{paths[r.Intn(len(paths))], int64(40000 * math.Exp(0.6*r.NormFloat64()))}
	}
	return out
}

// backlog is the same traffic while a queue backs up: every request waits
// 10 µs longer than the one before it, on top of its own service time.
func backlog(n int, seed int64) []Req {
	out := traffic(n, seed)
	for i := range out {
		out[i].Micros += int64(i) * 10
	}
	return out
}

// ascending is a log that arrives already ordered fastest first - an export
// that came back sorted by duration. Every request is the slowest yet.
func ascending(n int, seed int64) []Req {
	out := traffic(n, seed)
	slices.SortFunc(out, func(a, b Req) int { return int(a.Micros - b.Micros) })
	return out
}

// climbing is the extreme: every request strictly slower than the one before
// it, so every one gets into the top k.
func climbing(n int, seed int64) []Req {
	out := traffic(n, seed)
	for i := range out {
		out[i].Micros = int64(i)*10 + out[i].Micros%10
	}
	return out
}

func micros(rs []Req) []int64 {
	var out []int64
	for _, r := range rs {
		out = append(out, r.Micros)
	}
	return out
}

func TestBothAgree(t *testing.T) {
	cases := [][]Req{nil, traffic(1, 1), traffic(9, 2), traffic(10, 3), traffic(11, 4), traffic(5000, 5), backlog(5000, 6), ascending(5000, 7), climbing(5000, 8)}
	for _, s := range shapes {
		cases = append(cases, s.log)
	}
	for _, c := range cases {
		for _, k := range []int{1, 10, 100} {
			a, b := SlowestBySort(c, k), SlowestByHeap(c, k)
			if !slices.Equal(micros(a), micros(b)) {
				t.Fatalf("%d requests, k=%d: sort and heap disagree", len(c), k)
			}
		}
	}
}

func TestLogUntouched(t *testing.T) {
	log := traffic(1000, 7)
	before := slices.Clone(log)
	SlowestBySort(log, 10)
	SlowestByHeap(log, 10)
	if !slices.Equal(log, before) {
		t.Fatal("the shared log was reordered")
	}
}

// How many requests get past the comparison with the root.
func TestEntries(t *testing.T) {
	for _, s := range shapes {
		h := make(kept, 0, 10)
		entered := 0
		for _, r := range s.log {
			if len(h) < 10 {
				h = append(h, r)
				slices.SortFunc(h, func(a, b Req) int { return int(a.Micros - b.Micros) })
				entered++
				continue
			}
			if r.Micros > h[0].Micros {
				h[0] = r
				slices.SortFunc(h, func(a, b Req) int { return int(a.Micros - b.Micros) })
				entered++
			}
		}
		t.Logf("%-9s %9d requests  |  entered the top 10: %9d (%.4f%%)",
			s.name, len(s.log), entered, 100*float64(entered)/float64(len(s.log)))
	}
}
