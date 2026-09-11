package make_array_zero_by_subtracting_equal_amounts

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestMinimumOperations(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "example 1", nums: []int{1, 5, 0, 3, 5}, want: 3},
		{name: "example 2", nums: []int{0}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinimumOperations(tt.nums); got != tt.want {
				t.Errorf("MinimumOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
