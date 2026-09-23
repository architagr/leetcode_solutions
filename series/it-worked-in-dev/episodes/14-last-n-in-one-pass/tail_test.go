package lastninonepass

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

// A log line the size of a real one: a timestamp, a level, a request id and a
// message. Built per line rather than taken from a pool, because a real
// scanner allocates each line too and the episode is about what is retained,
// not about who allocated it.
func logOf(lines int) func() Log {
	return func() Log {
		return func(yield func(string) bool) {
			for i := 0; i < lines; i++ {
				s := fmt.Sprintf(
					"2026-09-23T19:04:%02dZ INFO req=%08x GET /api/v2/orders 200 %dms",
					i%60, i, i%97)
				if !yield(s) {
					return
				}
			}
		}
	}
}

func TestAllImplementationsAgree(t *testing.T) {
	for _, lines := range []int{0, 1, 7, 199, 200, 201, 5000} {
		for _, n := range []int{1, 200, 5000} {
			open := logOf(lines)
			want := LastNByCollecting(open(), n)
			for _, impl := range []struct {
				name string
				got  []string
			}{
				{"copying", LastNByCollectingAndCopying(open(), n)},
				{"counting twice", LastNByCountingTwice(open, n)},
				{"ring", LastNByRing(open(), n)},
			} {
				if len(want) == 0 && len(impl.got) == 0 {
					continue
				}
				if !reflect.DeepEqual(impl.got, want) {
					t.Fatalf("%d lines, n=%d: %s disagrees\n got %d lines\nwant %d lines",
						lines, n, impl.name, len(impl.got), len(want))
				}
			}
		}
	}
}

// The part of the collecting version that surprises people. all[len-n:] is a
// window onto the original array, so the returned value has the capacity of
// the whole log and keeps every line of it reachable.
func TestCollectedTailPinsTheWholeLog(t *testing.T) {
	const lines, n = 100000, 200
	open := logOf(lines)

	tail := LastNByCollecting(open(), n)
	if len(tail) != n {
		t.Fatalf("len is %d, want %d", len(tail), n)
	}
	if cap(tail) != n {
		t.Logf("a 200-line answer with capacity %d: the other %d slots are the "+
			"rest of the log, still reachable", cap(tail), cap(tail)-n)
	} else {
		t.Fatal("the tail no longer shares the collected array; the episode " +
			"says it does")
	}

	copied := LastNByCollectingAndCopying(open(), n)
	if cap(copied) != n {
		t.Fatalf("the copying version has capacity %d, want %d", cap(copied), n)
	}
}

// retained measures the live heap with the result still held, which is the
// number this episode is actually about. B/op counts what a function
// allocated; this counts what it will not let the collector take back.
func retained(t *testing.T, fn func() any) int64 {
	t.Helper()
	var before, after runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&before)
	v := fn()
	runtime.GC()
	runtime.ReadMemStats(&after)
	runtime.KeepAlive(v)
	return int64(after.HeapAlloc) - int64(before.HeapAlloc)
}

// medianRetained runs the measurement three times and takes the middle one.
// A single reading can come back negative when the collector hands back more
// than the call took, which is noise rather than a result.
func medianRetained(t *testing.T, fn func() any) int64 {
	t.Helper()
	r := []int64{retained(t, fn), retained(t, fn), retained(t, fn)}
	if r[0] > r[1] {
		r[0], r[1] = r[1], r[0]
	}
	if r[1] > r[2] {
		r[1], r[2] = r[2], r[1]
	}
	if r[0] > r[1] {
		r[0], r[1] = r[1], r[0]
	}
	return r[1]
}

func TestRetainedHeap(t *testing.T) {
	const n = 200
	for _, lines := range []int{200, 10000, 1000000, 2000000} {
		open := logOf(lines)
		collecting := medianRetained(t, func() any { return LastNByCollecting(open(), n) })
		copying := medianRetained(t, func() any { return LastNByCollectingAndCopying(open(), n) })
		ring := medianRetained(t, func() any { return LastNByRing(open(), n) })
		t.Logf("%9d lines, n=200  collecting %11d B (cap %7d)  copying %8d B  ring %8d B",
			lines, collecting, cap(LastNByCollecting(open(), n)), copying, ring)
	}
}

func chainOf(lines int) *Line {
	var head *Line
	for i := lines - 1; i >= 0; i-- {
		head = &Line{Text: fmt.Sprintf("line %d", i), Next: head}
	}
	return head
}

func TestChainImplementationsAgree(t *testing.T) {
	for _, lines := range []int{1, 2, 17, 1000} {
		head := chainOf(lines)
		for n := 1; n <= lines+2; n++ {
			a, b := NthFromEndByCounting(head, n), NthFromEndByGap(head, n)
			if a != b {
				t.Fatalf("%d lines, n=%d: counting gives %v, the gap gives %v",
					lines, n, a, b)
			}
		}
	}
}

// The gap version reads the chain once. The counting version reads it twice,
// and the second read is the one a stream cannot do.
func TestGapTouchesEachEntryOnce(t *testing.T) {
	const lines, n = 1000, 7
	head := chainOf(lines)

	got := NthFromEndByGap(head, n)
	want := fmt.Sprintf("line %d", lines-n)
	if got == nil || got.Text != want {
		t.Fatalf("got %v, want %q", got, want)
	}
}
