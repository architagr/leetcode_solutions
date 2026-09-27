package nextbiggerspike

import (
	"math"
	"math/rand"
	"reflect"
	"testing"
)

// daily is n periods of traffic with a daily cycle (period per day), weekly
// swing, a slow trend and noise. trend is the change per day.
func daily(n, perDay int, trend float64, seed int64) []int {
	r := rand.New(rand.NewSource(seed))
	out := make([]int, n)
	for i := range out {
		day := float64(i) / float64(perDay)
		v := 10000 + 6000*math.Sin(2*math.Pi*day) + 1500*math.Sin(2*math.Pi*day/7) + trend*day
		out[i] = int(v) + r.Intn(800)
	}
	return out
}

// launch is traffic after a launch: a big first week that decays towards a
// steady base, with the same daily cycle and noise on top. Every day's peak is
// higher than anything that follows it, so it is never beaten.
func launch(n, perDay int, seed int64) []int {
	r := rand.New(rand.NewSource(seed))
	out := make([]int, n)
	for i := range out {
		day := float64(i) / float64(perDay)
		level := 2000 + 20000*math.Exp(-day/20)
		out[i] = int(level*(1+0.6*math.Sin(2*math.Pi*day))) + r.Intn(400)
	}
	return out
}

func TestBothAgree(t *testing.T) {
	cases := [][]int{nil, {5}, {1, 2}, {2, 1}, {3, 3, 3}, {73, 74, 75, 71, 69, 72, 76, 73},
		daily(3000, 24, 5, 1), daily(3000, 24, -20, 2)}
	for _, s := range shapes {
		cases = append(cases, s.load)
	}
	for _, c := range cases {
		if a, b := WaitByScan(c), WaitByStack(c); !reflect.DeepEqual(a, b) {
			t.Fatalf("%d periods: scan and stack disagree", len(c))
		}
	}
}

func TestSteps(t *testing.T) {
	for _, s := range shapes {
		steps := 0
		for i := range s.load {
			j := i + 1
			for ; j < len(s.load) && s.load[j] <= s.load[i]; j++ {
			}
			steps += j - i - 1
		}
		never := 0
		for _, w := range WaitByStack(s.load) {
			if w == 0 {
				never++
			}
		}
		t.Logf("%-14s %7d periods, %6d never beaten  |  forward steps the scan takes: %12d",
			s.name, len(s.load), never, steps)
	}
}
