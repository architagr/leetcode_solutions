package insertdeletegetrandomo1duplicatesallowed

import "testing"

// The operation sequence from the problem statement. This collection
// allows duplicates, so Insert reports whether the value was absent
// beforehand rather than whether anything was stored.
func TestRandomizedCollection(t *testing.T) {
	rc := Constructor()

	if got := rc.Insert(1); !got {
		t.Errorf("Insert(1) = %v, want true, the collection was empty", got)
	}
	if got := rc.Insert(1); got {
		t.Errorf("second Insert(1) = %v, want false, 1 was already present", got)
	}
	if got := rc.Insert(2); !got {
		t.Errorf("Insert(2) = %v, want true", got)
	}
	// Two 1s and one 2, so GetRandom must return 1 about twice as often.
	counts := map[int]int{}
	for i := 0; i < 300; i++ {
		counts[rc.GetRandom()]++
	}
	if counts[1]+counts[2] != 300 {
		t.Errorf("GetRandom returned values outside the collection: %v", counts)
	}
	if counts[1] == 0 || counts[2] == 0 {
		t.Errorf("GetRandom never returned one of the values: %v", counts)
	}
	if got := rc.Remove(1); !got {
		t.Errorf("Remove(1) = %v, want true", got)
	}
	// One 1 and one 2 remain.
	for i := 0; i < 50; i++ {
		if v := rc.GetRandom(); v != 1 && v != 2 {
			t.Fatalf("GetRandom() = %v, want 1 or 2", v)
		}
	}
}
