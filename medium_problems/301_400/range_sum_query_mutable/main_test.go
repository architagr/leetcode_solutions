package range_sum_query_mutable

import "testing"

// The operation sequence from the problem statement: a sum, an update
// that changes the middle element, then the same range again.
func TestNumArrayUpdateAndSumRange(t *testing.T) {
	na := Constructor([]int{1, 3, 5})

	if got := na.SumRange(0, 2); got != 9 {
		t.Fatalf("SumRange(0, 2) = %v, want 9", got)
	}
	na.Update(1, 2)
	if got := na.SumRange(0, 2); got != 8 {
		t.Errorf("SumRange(0, 2) after Update(1, 2) = %v, want 8", got)
	}
}
