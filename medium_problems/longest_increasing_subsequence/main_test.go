package longest_increasing_subsequence

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "example 1", nums: []int{10, 9, 2, 5, 3, 7, 101, 18}, want: 4},
		{name: "example 2", nums: []int{0, 1, 0, 3, 2, 3}, want: 4},
		{name: "example 3", nums: []int{7, 7, 7, 7, 7, 7, 7}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLIS(tt.nums); got != tt.want {
				t.Errorf("LengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
