package mergetwosortedfeeds

import (
	"slices"
	"sort"
)

// Event is one row of an activity feed.
//
// Two queries produce them - your own activity, and the activity of everyone
// you follow - and both queries end in ORDER BY at. So each slice arrives
// already sorted. Neither is sorted against the other, which is the only
// thing the page still has to do.
type Event struct {
	At     int64  // when it happened, and what the database ordered by
	Source string // which of the two queries produced it
	Seq    int    // position within that query's result set
}

// MergeBySorting is the version I have written more than once: concatenate
// the two result sets and hand the whole thing to sort.
//
// It is one line of intent and it is correct for any two inputs, sorted or
// not, which is exactly why it survives review.
func MergeBySorting(own, followed []Event) []Event {
	out := make([]Event, 0, len(own)+len(followed))
	out = append(out, own...)
	out = append(out, followed...)
	sort.Slice(out, func(i, j int) bool { return out[i].At < out[j].At })
	return out
}

// MergeByStableSorting is the same thing with the tie-breaking fixed, which
// is what you reach for the first time the feed reorders itself on refresh.
func MergeByStableSorting(own, followed []Event) []Event {
	out := make([]Event, 0, len(own)+len(followed))
	out = append(out, own...)
	out = append(out, followed...)
	sort.SliceStable(out, func(i, j int) bool { return out[i].At < out[j].At })
	return out
}

// MergeBySortFunc is the modern spelling. slices.SortFunc is generic, so it
// swaps Events directly instead of going through the reflect-based swapper
// sort.Slice builds. Same algorithm, much less overhead per swap - benchmarked
// here so the episode is not beating a version nobody writes any more.
func MergeBySortFunc(own, followed []Event) []Event {
	out := make([]Event, 0, len(own)+len(followed))
	out = append(out, own...)
	out = append(out, followed...)
	slices.SortFunc(out, func(a, b Event) int {
		switch {
		case a.At < b.At:
			return -1
		case a.At > b.At:
			return 1
		}
		return 0
	})
	return out
}

// MergeByWalking takes the smaller of the two fronts, over and over.
func MergeByWalking(own, followed []Event) []Event {
	out := make([]Event, 0, len(own)+len(followed))
	i, j := 0, 0
	for i < len(own) && j < len(followed) {
		// <= rather than < is the whole tie-break policy: on an equal
		// timestamp your own event goes first, every time, on every machine.
		if own[i].At <= followed[j].At {
			out = append(out, own[i])
			i++
		} else {
			out = append(out, followed[j])
			j++
		}
	}
	// One side is empty. What is left in the other is already in order and
	// already later than everything placed, so it is copied in one move
	// rather than compared element by element.
	out = append(out, own[i:]...)
	out = append(out, followed[j:]...)
	return out
}
