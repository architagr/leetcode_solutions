package mergetwosortedfeeds

import (
	"slices"
	"sort"
	"testing"
)

// Counting comparisons, because "the sort is redoing work" is a claim and
// this is the measurement behind it. The merge's count is bounded by
// n+m-1 by construction: every comparison places exactly one event.
func comparisons(own, followed []Event) (slice, stable, sortFunc, merge int) {
	cat := func() []Event {
		out := make([]Event, 0, len(own)+len(followed))
		out = append(out, own...)
		return append(out, followed...)
	}

	a := cat()
	sort.Slice(a, func(i, j int) bool { slice++; return a[i].At < a[j].At })

	b := cat()
	sort.SliceStable(b, func(i, j int) bool { stable++; return b[i].At < b[j].At })

	c := cat()
	slices.SortFunc(c, func(x, y Event) int {
		sortFunc++
		switch {
		case x.At < y.At:
			return -1
		case x.At > y.At:
			return 1
		}
		return 0
	})

	i, j := 0, 0
	for i < len(own) && j < len(followed) {
		merge++
		if own[i].At <= followed[j].At {
			i++
		} else {
			j++
		}
	}
	return
}

func TestComparisonCounts(t *testing.T) {
	for _, s := range shapes {
		own, followed := feeds(s.n, s.m, 1, 1)
		sl, st, sf, mg := comparisons(own, followed)
		n := len(own) + len(followed)
		t.Logf("%-18s %7d events  sort.Slice %10d  SliceStable %9d  "+
			"slices.SortFunc %10d  merge %8d  (merge is %.1f%% of n)",
			s.name, n, sl, st, sf, mg, 100*float64(mg)/float64(n))
		if mg >= n {
			t.Fatalf("%s: the merge made %d comparisons for %d events, "+
				"which is not linear", s.name, mg, n)
		}
	}
}
