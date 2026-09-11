package insertdeletegetrandomo1

import "testing"

// The operation sequence from the problem statement. GetRandom is not
// deterministic, so it is checked against the set's current contents
// rather than one expected value.
func TestRandomizedSet(t *testing.T) {
	rs := Constructor()

	if got := rs.Insert(1); !got {
		t.Errorf("Insert(1) = %v, want true", got)
	}
	if got := rs.Remove(2); got {
		t.Errorf("Remove(2) = %v, want false, 2 is not present", got)
	}
	if got := rs.Insert(2); !got {
		t.Errorf("Insert(2) = %v, want true", got)
	}
	if got := rs.GetRandom(); got != 1 && got != 2 {
		t.Errorf("GetRandom() = %v, want 1 or 2", got)
	}
	if got := rs.Remove(1); !got {
		t.Errorf("Remove(1) = %v, want true", got)
	}
	if got := rs.Insert(2); got {
		t.Errorf("Insert(2) = %v, want false, 2 is already present", got)
	}
	for i := 0; i < 10; i++ {
		if got := rs.GetRandom(); got != 2 {
			t.Fatalf("GetRandom() = %v, want 2, the only remaining element", got)
		}
	}
}
