package minimum_average_difference

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestMinimumAverageDifference(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "example 1", nums: []int{2, 5, 3, 9, 5, 3}, want: 3},
		{name: "example 2", nums: []int{0}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumAverageDifference(tt.nums); got != tt.want {
				t.Errorf("minimumAverageDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
