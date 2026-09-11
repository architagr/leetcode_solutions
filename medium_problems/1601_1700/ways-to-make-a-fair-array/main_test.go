package waystomakeafairarray

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestWaysToMakeFair(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "example 1", nums: []int{2, 1, 6, 4}, want: 1},
		{name: "example 2", nums: []int{1, 1, 1}, want: 3},
		{name: "example 3", nums: []int{1, 2, 3}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := waysToMakeFair(tt.nums); got != tt.want {
				t.Errorf("waysToMakeFair() = %v, want %v", got, tt.want)
			}
		})
	}
}
