package jumpgamevi

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestMaxResult(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{name: "example 1", nums: []int{1, -1, -2, 4, -7, 3}, k: 2, want: 7},
		{name: "example 2", nums: []int{10, -5, -2, 4, 0, 3}, k: 3, want: 17},
		{name: "example 3", nums: []int{1, -5, -20, 4, -1, 3, -6, -3}, k: 2, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxResult(tt.nums, tt.k); got != tt.want {
				t.Errorf("maxResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
