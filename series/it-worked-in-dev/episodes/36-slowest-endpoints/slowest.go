package slowestendpoints

import (
	"cmp"
	"container/heap"
	"slices"
)

// The incident page opens on the ten slowest requests of the last hour: which
// endpoint, how long. The request log is shared with every other panel, so
// nothing here may reorder it.

// Req is one line of the request log.
type Req struct {
	Path   string
	Micros int64
}

func slowestFirst(a, b Req) int { return cmp.Compare(b.Micros, a.Micros) }

// SlowestBySort is what I would write: copy the log, sort it slowest first,
// keep the first k.
func SlowestBySort(log []Req, k int) []Req {
	all := slices.Clone(log) // the log is shared; sort a copy
	slices.SortFunc(all, slowestFirst)
	return all[:min(k, len(all))]
}

// kept is a min-heap on latency: the fastest of the requests kept so far is
// at the root, because it is the one the next slow request replaces.
type kept []Req

func (h kept) Len() int           { return len(h) }
func (h kept) Less(i, j int) bool { return h[i].Micros < h[j].Micros }
func (h kept) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *kept) Push(x any)        { *h = append(*h, x.(Req)) }
func (h *kept) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

// SlowestByHeap keeps only the k slowest seen so far. A request has to beat
// the fastest of those to get in - one comparison with the root - and most
// requests in an hour do not.
func SlowestByHeap(log []Req, k int) []Req {
	h := make(kept, 0, k)
	for _, r := range log {
		if len(h) < k {
			heap.Push(&h, r)
			continue
		}
		if r.Micros > h[0].Micros { // beats the weakest of the k
			h[0] = r
			heap.Fix(&h, 0)
		}
	}
	slices.SortFunc(h, slowestFirst) // k of them, for display
	return h
}
