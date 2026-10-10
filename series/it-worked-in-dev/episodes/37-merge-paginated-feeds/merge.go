package mergepaginatedfeeds

import (
	"cmp"
	"container/heap"
	"slices"
)

// The home timeline: for every account you follow, the posts service returns
// that account's latest posts, newest first. The timeline shows one page of
// them all together, newest first.

// Post is one post. IDs are unique, so they settle a tie in time.
type Post struct {
	ID int64
	At int64 // unix ms
}

// newer orders newest first, then by ID, so every two posts compare unequal.
func newer(a, b Post) int {
	if c := cmp.Compare(b.At, a.At); c != 0 {
		return c
	}
	return cmp.Compare(b.ID, a.ID)
}

// PageBySort is what I would write: put every feed in one slice, sort it,
// take the first page.
func PageBySort(feeds [][]Post, n int) []Post {
	var all []Post
	for _, f := range feeds {
		all = append(all, f...)
	}
	slices.SortFunc(all, newer)
	return all[:min(n, len(all))]
}

// PageByFolding uses the merge every feed is already sorted for: merge the
// first two, merge the result with the third, and so on.
func PageByFolding(feeds [][]Post, n int) []Post {
	var out []Post
	for _, f := range feeds {
		out = merge(out, f)
	}
	return out[:min(n, len(out))]
}

// merge is the two-list merge: take the newer front, advance that side.
func merge(a, b []Post) []Post {
	out := make([]Post, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if newer(a[i], b[j]) <= 0 {
			out = append(out, a[i])
			i++
		} else {
			out = append(out, b[j])
			j++
		}
	}
	out = append(out, a[i:]...)
	return append(out, b[j:]...)
}

// fronts is a heap of one position per feed, newest post at the root.
type fronts struct {
	feeds [][]Post
	at    []cursor
}

type cursor struct{ feed, i int }

func (h *fronts) post(c cursor) Post { return h.feeds[c.feed][c.i] }
func (h *fronts) Len() int           { return len(h.at) }
func (h *fronts) Less(i, j int) bool { return newer(h.post(h.at[i]), h.post(h.at[j])) < 0 }
func (h *fronts) Swap(i, j int)      { h.at[i], h.at[j] = h.at[j], h.at[i] }
func (h *fronts) Push(x any)         { h.at = append(h.at, x.(cursor)) }
func (h *fronts) Pop() any {
	x := h.at[len(h.at)-1]
	h.at = h.at[:len(h.at)-1]
	return x
}

// PageByHeap keeps only the front of each feed in a heap. The newest post
// overall is the newest of the fronts; taking it moves that one feed along.
// It stops when the page is full, and the rest of every feed is never read.
func PageByHeap(feeds [][]Post, n int) []Post {
	h := &fronts{feeds: feeds, at: make([]cursor, 0, len(feeds))}
	for f, posts := range feeds {
		if len(posts) > 0 {
			h.at = append(h.at, cursor{f, 0})
		}
	}
	heap.Init(h) // all fronts at once: linear, not k pushes
	out := make([]Post, 0, n)
	for len(out) < n && h.Len() > 0 {
		c := h.at[0]
		out = append(out, h.post(c))
		if c.i+1 < len(feeds[c.feed]) {
			h.at[0] = cursor{c.feed, c.i + 1} // that feed's next post takes its place
			heap.Fix(h, 0)
		} else {
			heap.Pop(h)
		}
	}
	return out
}
