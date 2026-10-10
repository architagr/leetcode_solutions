package whoholdsthetime

import (
	"math/rand"
	"reflect"
	"testing"
)

// walk is a trace of calls calls. At every step the current function either
// calls another (probability down, while under maxDepth) or returns. Gaps
// between events are 20 to 2,000 ns. Recursion is common, so function ids
// repeat across depths.
func walk(calls, maxDepth int, down float64, fns int, seed int64) []Event {
	r := rand.New(rand.NewSource(seed))
	var out []Event
	var at int64
	var stack []int
	tick := func() int64 { at += 20 + r.Int63n(1980); return at }
	made := 0
	for made < calls {
		if len(stack) == 0 || (len(stack) < maxDepth && r.Float64() < down) {
			fn := r.Intn(fns)
			stack = append(stack, fn)
			out = append(out, Event{fn, true, tick()})
			made++
			continue
		}
		out = append(out, Event{stack[len(stack)-1], false, tick()})
		stack = stack[:len(stack)-1]
	}
	for len(stack) > 0 {
		out = append(out, Event{stack[len(stack)-1], false, tick()})
		stack = stack[:len(stack)-1]
	}
	return out
}

// chain is one function recursing depth times and unwinding.
func chain(depth int, seed int64) []Event {
	r := rand.New(rand.NewSource(seed))
	var out []Event
	var at int64
	for i := 0; i < depth; i++ {
		at += 20 + r.Int63n(1980)
		out = append(out, Event{i % 3, true, at})
	}
	for i := depth - 1; i >= 0; i-- {
		at += 20 + r.Int63n(1980)
		out = append(out, Event{i % 3, false, at})
	}
	return out
}

func maxDepth(events []Event) int {
	d, m := 0, 0
	for _, e := range events {
		if e.Start {
			d++
			if d > m {
				m = d
			}
		} else {
			d--
		}
	}
	return m
}

// The example from the write-up: main calls parse, parse calls lex twice,
// then main calls render.
var example = []Event{
	{0, true, 0}, {1, true, 2}, {2, true, 3}, {2, false, 5}, {2, true, 6},
	{2, false, 9}, {1, false, 10}, {3, true, 11}, {3, false, 14}, {0, false, 15},
}

func TestExample(t *testing.T) {
	want := []int64{4, 3, 5, 3} // main, parse, lex, render
	for name, fn := range map[string]func(int, []Event) []int64{"spans": SelfBySpans, "stack": SelfByStack} {
		if got := fn(4, example); !reflect.DeepEqual(got, want) {
			t.Fatalf("%s: got %v, want %v", name, got, want)
		}
	}
}

// Subtracting every call inside a call, not only the direct ones, takes the
// grandchildren off twice.
func TestSubtractingEverythingDoubleCounts(t *testing.T) {
	// main 0..15 contains parse 2..10, lex 3..5, lex 6..9, render 11..14:
	// 15 - (8 + 2 + 3 + 3) = -1, where the right answer is 4.
	if got := int64(15 - (8 + 2 + 3 + 3)); got >= 0 {
		t.Fatalf("expected the naive subtraction to go negative, got %d", got)
	}
}

func TestBothAgree(t *testing.T) {
	cases := [][]Event{nil, example, chain(1, 1), chain(50, 2), walk(500, 6, 0.5, 5, 3)}
	for _, s := range shapes {
		cases = append(cases, s.events)
	}
	for _, c := range cases {
		a, b := SelfBySpans(8, c), SelfByStack(8, c)
		if !reflect.DeepEqual(a, b) {
			t.Fatalf("%d events: spans and stack disagree", len(c))
		}
		// Every nanosecond something was running is charged exactly once.
		var total, busy int64
		for _, v := range b {
			total += v
		}
		depth := 0
		for i, e := range c {
			if depth > 0 {
				busy += e.At - c[i-1].At
			}
			if e.Start {
				depth++
			} else {
				depth--
			}
		}
		if total != busy {
			t.Fatalf("self times sum to %d, busy time is %d", total, busy)
		}
	}
}

func TestSteps(t *testing.T) {
	for _, s := range shapes {
		var spans []span
		var open []int
		for _, e := range s.events {
			if e.Start {
				open = append(open, len(spans))
				spans = append(spans, span{start: e.At, depth: len(open) - 1})
			} else {
				spans[open[len(open)-1]].end = e.At
				open = open[:len(open)-1]
			}
		}
		steps := 0
		for i, sp := range spans {
			for j := i + 1; j < len(spans) && spans[j].start < sp.end; j++ {
				steps++
			}
		}
		t.Logf("%-10s %7d calls, max depth %5d  |  spans the inner loop reads: %11d",
			s.name, len(spans), maxDepth(s.events), steps)
	}
}
