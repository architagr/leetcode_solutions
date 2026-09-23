package mergetwosortedfeeds

import (
	"math/rand"
	"reflect"
	"testing"
)

// feeds builds two already-sorted result sets whose timestamps overlap.
//
// tick is the gap between consecutive events within one feed: tick=1 with
// many events means a lot of equal timestamps across the two feeds, which is
// what a bulk insert or a coarse clock actually produces.
func feeds(n, m int, tick int64, seed int64) ([]Event, []Event) {
	r := rand.New(rand.NewSource(seed))
	own := make([]Event, n)
	var at int64
	for i := range own {
		at += r.Int63n(tick + 1)
		own[i] = Event{At: at, Source: "own", Seq: i}
	}
	followed := make([]Event, m)
	at = 0
	for i := range followed {
		at += r.Int63n(tick + 1)
		followed[i] = Event{At: at, Source: "followed", Seq: i}
	}
	return own, followed
}

func isSorted(t *testing.T, got []Event) {
	t.Helper()
	for i := 1; i < len(got); i++ {
		if got[i-1].At > got[i].At {
			t.Fatalf("out of order at %d: %d then %d", i, got[i-1].At, got[i].At)
		}
	}
}

// Every implementation has to produce the same multiset in the same
// timestamp order. Where they differ is only in how ties are arranged, so
// this compares the timestamp sequence rather than the events.
func TestAllImplementationsAgreeOnOrder(t *testing.T) {
	for seed := int64(0); seed < 500; seed++ {
		n := int(seed%37) + 1
		m := int(seed%23) + 1
		own, followed := feeds(n, m, 3, seed)

		want := MergeByWalking(own, followed)
		isSorted(t, want)
		if len(want) != n+m {
			t.Fatalf("seed %d: merged %d events, expected %d", seed, len(want), n+m)
		}

		for _, impl := range []struct {
			name string
			fn   func(a, b []Event) []Event
		}{
			{"sort.Slice", MergeBySorting},
			{"sort.SliceStable", MergeByStableSorting},
			{"slices.SortFunc", MergeBySortFunc},
		} {
			got := impl.fn(own, followed)
			isSorted(t, got)
			for i := range got {
				if got[i].At != want[i].At {
					t.Fatalf("seed %d: %s disagrees at %d", seed, impl.name, i)
				}
			}
		}
	}
}

// The merge does not move the inputs. A caller reusing the two result sets
// for anything else - a count, a second render - must see them untouched.
func TestInputsAreNotMutated(t *testing.T) {
	own, followed := feeds(200, 200, 2, 7)
	a := append([]Event(nil), own...)
	b := append([]Event(nil), followed...)
	MergeByWalking(own, followed)
	if !reflect.DeepEqual(own, a) || !reflect.DeepEqual(followed, b) {
		t.Fatal("the merge wrote into its inputs")
	}
}

// The merge's tie-break is a decision, not an accident: equal timestamps put
// your own event first and keep each feed's own order intact.
func TestWalkingIsStable(t *testing.T) {
	own := []Event{{At: 5, Source: "own", Seq: 0}, {At: 5, Source: "own", Seq: 1}}
	followed := []Event{{At: 5, Source: "followed", Seq: 0}, {At: 5, Source: "followed", Seq: 1}}
	got := MergeByWalking(own, followed)
	want := []Event{own[0], own[1], followed[0], followed[1]}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ties rearranged:\n got %v\nwant %v", got, want)
	}
}

// What sort.Slice actually does to equal timestamps, measured rather than
// asserted in prose. tick=1 gives roughly one event in two sharing a
// timestamp with another, which is what a coarse clock or a bulk insert
// produces.
//
// If Go's sort ever starts agreeing with the stable order this fails, and
// the episode's stability argument gets rewritten rather than quietly
// surviving. That is what it is pinned for.
func TestUnstableSortRearrangesTies(t *testing.T) {
	own, followed := feeds(200, 200, 1, 3)
	stable := MergeByWalking(own, followed)
	unstable := MergeBySorting(own, followed)
	moved := 0
	for i := range stable {
		if stable[i] != unstable[i] {
			moved++
		}
	}
	if moved == 0 {
		t.Fatal("sort.Slice preserved input order here; the episode says it does not")
	}
	t.Logf("400 events, half of them sharing a timestamp with another: "+
		"sort.Slice puts %d in a different place than the merge does", moved)
}

// The one that costs a bug report. One new event arrives, and events that
// have nothing to do with it swap places on the page - because the sort's
// pivot choices depend on the length of the slice, not on the new event.
func TestOneNewEventReordersUnrelatedItems(t *testing.T) {
	own, followed := feeds(200, 200, 1, 3)
	before := MergeBySorting(own, followed)

	// Somebody you follow posts once more, later than everything so far.
	extra := append(append([]Event(nil), followed...),
		Event{At: 1 << 40, Source: "followed", Seq: len(followed)})
	after := MergeBySorting(own, extra)

	moved := 0
	for i := 0; i < len(before); i++ {
		if before[i] != after[i] {
			moved++
		}
	}
	if moved == 0 {
		t.Fatal("this input survived the extra event; the episode says it does not")
	}
	t.Logf("one event appended after everything else: %d of the %d rows "+
		"in front of it changed position", moved, len(before))

	// The merge does not have the property at all.
	mb := MergeByWalking(own, followed)
	ma := MergeByWalking(own, extra)
	for i := range mb {
		if mb[i] != ma[i] {
			t.Fatalf("the merge moved row %d when an event was appended", i)
		}
	}
}

// An empty feed is the common case on a new account, and it is the input
// that breaks a merge written without the tail copy.
func TestEmptyFeeds(t *testing.T) {
	own, followed := feeds(5, 0, 2, 11)
	if got := MergeByWalking(own, followed); len(got) != 5 {
		t.Fatalf("empty followed feed: got %d events", len(got))
	}
	if got := MergeByWalking(nil, own); len(got) != 5 {
		t.Fatalf("empty own feed: got %d events", len(got))
	}
	if got := MergeByWalking(nil, nil); len(got) != 0 {
		t.Fatalf("two empty feeds: got %d events", len(got))
	}
}
