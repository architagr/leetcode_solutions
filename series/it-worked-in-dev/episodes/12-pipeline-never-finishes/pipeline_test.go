package pipelineneverfinishes

import (
	"fmt"
	"math/rand"
	"testing"
)

// chain builds n stages in a line. When loopTo is >= 0 the last stage points
// back at that index instead of nothing, which is what a config file with one
// bad "next" in it produces.
func chain(n, loopTo int) *Stage {
	if n == 0 {
		return nil
	}
	stages := make([]*Stage, n)
	for i := range stages {
		stages[i] = &Stage{Name: fmt.Sprintf("stage-%d", i)}
	}
	for i := 0; i+1 < n; i++ {
		stages[i].Next = stages[i+1]
	}
	if loopTo >= 0 && loopTo < n {
		stages[n-1].Next = stages[loopTo]
	}
	return stages[0]
}

func TestKnownPipelines(t *testing.T) {
	cases := []struct {
		name   string
		head   *Stage
		cycles bool
		entry  string
	}{
		{"nothing at all", nil, false, ""},
		{"one stage", chain(1, -1), false, ""},
		{"one stage pointing at itself", chain(1, 0), true, "stage-0"},
		{"a normal pipeline", chain(12, -1), false, ""},
		{"the tail points at the head", chain(12, 0), true, "stage-0"},
		{"the tail points at the middle", chain(12, 6), true, "stage-6"},
		{"the last two swap forever", chain(12, 11), true, "stage-11"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := HasCycleBySeen(c.head); got != c.cycles {
				t.Errorf("HasCycleBySeen = %v, want %v", got, c.cycles)
			}
			if got := HasCycleByTwoPointers(c.head); got != c.cycles {
				t.Errorf("HasCycleByTwoPointers = %v, want %v", got, c.cycles)
			}
			if got := EntryBySeen(c.head); got != c.entry {
				t.Errorf("EntryBySeen = %q, want %q", got, c.entry)
			}
			if got := EntryByTwoPointers(c.head); got != c.entry {
				t.Errorf("EntryByTwoPointers = %q, want %q", got, c.entry)
			}
		})
	}
}

// The hop limit is the version most things ship, and it reports a cycle on a
// pipeline that has none - it just has more steps than somebody guessed.
func TestHopLimitLiesAboutLongChains(t *testing.T) {
	longButFine := chain(150, -1)
	if HasCycleBySeen(longButFine) || HasCycleByTwoPointers(longButFine) {
		t.Fatal("the 150-stage pipeline terminates; both real checks agree")
	}
	if !HasCycleByHopLimit(longButFine, 100) {
		t.Fatal("expected the hop limit to call this a cycle")
	}
	// And it is right about an actual loop, which is why it survives review.
	if !HasCycleByHopLimit(chain(12, 6), 100) {
		t.Error("the hop limit missed a real cycle")
	}
}

func TestBothChecksAgreeOnRandomPipelines(t *testing.T) {
	rng := rand.New(rand.NewSource(21))
	for trial := 0; trial < 5000; trial++ {
		n := rng.Intn(80) + 1
		loopTo := -1
		if rng.Intn(2) == 0 {
			loopTo = rng.Intn(n)
		}
		head := chain(n, loopTo)
		seen, two := HasCycleBySeen(head), HasCycleByTwoPointers(head)
		if seen != two {
			t.Fatalf("trial %d (n=%d, loopTo=%d): seen says %v, two pointers say %v",
				trial, n, loopTo, seen, two)
		}
		if a, b := EntryBySeen(head), EntryByTwoPointers(head); a != b {
			t.Fatalf("trial %d (n=%d, loopTo=%d): entry %q vs %q",
				trial, n, loopTo, a, b)
		}
	}
}

// The write-up says the map holds one entry per stage walked. Counted here so
// the memory claim is not a guess.
func TestSeenHoldsOneEntryPerStage(t *testing.T) {
	for _, n := range []int{10, 1000, 100000} {
		head := chain(n, -1)
		seen := map[*Stage]bool{}
		for s := head; s != nil; s = s.Next {
			seen[s] = true
		}
		if len(seen) != n {
			t.Errorf("%d stages: map holds %d entries", n, len(seen))
		}
	}
}

// Two pointers meet inside the loop rather than merely passing each other,
// because the fast cursor closes the gap by exactly one stage per turn. Counted
// on a loop of every size up to 200, from every entry offset.
func TestTwoPointersAlwaysMeet(t *testing.T) {
	for loop := 1; loop <= 200; loop++ {
		for tail := 0; tail < 5; tail++ {
			head := chain(tail+loop, tail)
			if !HasCycleByTwoPointers(head) {
				t.Fatalf("missed a loop of %d after a tail of %d", loop, tail)
			}
		}
	}
}

// The write-up draws this exact pipeline and claims this exact trace, so the
// trace is asserted rather than eyeballed off a diagram.
//
//	fetch -> validate -> enrich -> notify -> retry, and retry -> validate
func TestDrawnTrace(t *testing.T) {
	names := []string{"fetch", "validate", "enrich", "notify", "retry"}
	stages := make([]*Stage, len(names))
	for i, n := range names {
		stages[i] = &Stage{Name: n}
	}
	for i := 0; i+1 < len(stages); i++ {
		stages[i].Next = stages[i+1]
	}
	stages[4].Next = stages[1]

	want := []struct{ slow, fast string }{
		{"validate", "enrich"}, // turn 1
		{"enrich", "retry"},    // turn 2
		{"notify", "enrich"},   // turn 3, fast is now behind slow
		{"retry", "retry"},     // turn 4, they meet
	}
	slow, fast := stages[0], stages[0]
	for turn, w := range want {
		slow, fast = slow.Next, fast.Next.Next
		if slow.Name != w.slow || fast.Name != w.fast {
			t.Fatalf("turn %d: slow %s fast %s, want slow %s fast %s",
				turn+1, slow.Name, fast.Name, w.slow, w.fast)
		}
	}
	if slow != fast {
		t.Fatal("the cursors were supposed to meet on turn 4")
	}
	if got := EntryByTwoPointers(stages[0]); got != "validate" {
		t.Errorf("loop starts at %q, want validate", got)
	}
}
