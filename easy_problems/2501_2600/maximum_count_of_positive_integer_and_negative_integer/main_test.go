package maximumcountofpositiveintegerandnegativeinteger

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestMaximumCount(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "example 1", nums: []int{-2, -1, -1, 1, 2, 3}, want: 3},
		{name: "example 2", nums: []int{-3, -2, -1, 0, 0, 1, 2}, want: 3},
		{name: "example 3", nums: []int{5, 20, 66, 1314}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumCount(tt.nums); got != tt.want {
				t.Errorf("maximumCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
