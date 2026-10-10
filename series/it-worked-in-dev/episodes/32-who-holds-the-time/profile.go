package whoholdsthetime

// A trace from an instrumented service: every function call logs a start and
// an end, with a nanosecond timestamp. Calls nest. The flame-graph panel needs
// each function's self time: how long it was actually running, not counting
// the functions it called. Timestamps are half-open - a call that ends at t
// stopped running at t, and its caller resumes at t.

// Event is one line of the trace.
type Event struct {
	Fn    int   // function id
	Start bool  // true for a call, false for its return
	At    int64 // ns since the trace began
}

type span struct {
	fn         int
	start, end int64
	depth      int
}

// SelfBySpans is what I would write: turn the trace into calls with a start,
// an end and a depth, then take each call's duration minus the durations of
// the calls made directly from it.
func SelfBySpans(n int, events []Event) []int64 {
	spans := make([]span, 0, len(events)/2)
	open := make([]int, 0, 64) // indices into spans of calls not yet returned
	for _, e := range events {
		if e.Start {
			open = append(open, len(spans))
			spans = append(spans, span{fn: e.Fn, start: e.At, depth: len(open) - 1})
			continue
		}
		spans[open[len(open)-1]].end = e.At
		open = open[:len(open)-1]
	}

	out := make([]int64, n)
	for i, s := range spans {
		self := s.end - s.start
		// spans are in start order, so everything this call made comes
		// straight after it, and stops at the first span starting after its end.
		for j := i + 1; j < len(spans) && spans[j].start < s.end; j++ {
			if spans[j].depth == s.depth+1 { // only direct callees
				self -= spans[j].end - spans[j].start
			}
		}
		out[s.fn] += self
	}
	return out
}

// SelfByStack replays the trace once. Between two consecutive events exactly
// one function is running - the one on top of the call stack - so that gap is
// its self time, and nobody else's.
func SelfByStack(n int, events []Event) []int64 {
	out := make([]int64, n)
	stack := make([]int, 0, 64) // function ids, innermost on top
	var last int64
	for _, e := range events {
		if len(stack) > 0 {
			out[stack[len(stack)-1]] += e.At - last // the top was running since last
		}
		last = e.At
		if e.Start {
			stack = append(stack, e.Fn)
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return out
}
