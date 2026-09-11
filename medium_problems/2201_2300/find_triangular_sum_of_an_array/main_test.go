package triangular_sum_of_an_array

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestTriangularSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "example 1", nums: []int{1, 2, 3, 4, 5}, want: 8},
		{name: "example 2", nums: []int{5}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := triangularSum(tt.nums); got != tt.want {
				t.Errorf("triangularSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
