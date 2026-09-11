package range_sum_query_immutable

import "testing"

// The operation sequence from the problem statement.
func TestNumArraySumRange(t *testing.T) {
	na := Constructor([]int{-2, 0, 3, -5, 2, -1})

	ops := []struct {
		left, right, want int
	}{
		{0, 2, 1},
		{2, 5, -1},
		{0, 5, -3},
	}
	for _, op := range ops {
		if got := na.SumRange(op.left, op.right); got != op.want {
			t.Errorf("SumRange(%d, %d) = %v, want %v", op.left, op.right, got, op.want)
		}
	}
}
