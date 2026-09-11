package longestharmonioussubsequence

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestFindLHS(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "example 1", nums: []int{1, 3, 2, 2, 5, 2, 3, 7}, want: 5},
		{name: "example 2", nums: []int{1, 2, 3, 4}, want: 2},
		{name: "example 3", nums: []int{1, 1, 1, 1}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLHS(tt.nums); got != tt.want {
				t.Errorf("findLHS() = %v, want %v", got, tt.want)
			}
		})
	}
}
