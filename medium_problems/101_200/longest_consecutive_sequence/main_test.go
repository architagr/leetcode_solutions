package longest_consecutive_sequence

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "example 1", nums: []int{100, 4, 200, 1, 3, 2}, want: 4},
		{name: "example 2", nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, want: 9},
		{name: "example 3", nums: []int{1, 0, 1, 2}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestConsecutive(tt.nums); got != tt.want {
				t.Errorf("LongestConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}
