package numberofwaystosplitarray

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestWaysToSplitArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "example 1", nums: []int{10, 4, -8, 7}, want: 2},
		{name: "example 2", nums: []int{2, 3, 1, 0}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := waysToSplitArray(tt.nums); got != tt.want {
				t.Errorf("waysToSplitArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
