package twosumiiidatastructuredesign

import "testing"

// The operation sequence from the problem statement: add 1, 3 and 5,
// then look for a pair summing to 4 (1 + 3) and to 7 (no such pair).
func TestTwoSum(t *testing.T) {
	ts := Constructor()
	ts.Add(1)
	ts.Add(3)
	ts.Add(5)

	if got := ts.Find(4); !got {
		t.Errorf("Find(4) = %v, want true (1 + 3)", got)
	}
	if got := ts.Find(7); got {
		t.Errorf("Find(7) = %v, want false", got)
	}
}
