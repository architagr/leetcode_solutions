package minimum_value_to_get_positive_step_by_step_sum

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestMinStartValue(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "example 1", nums: []int{-3, 2, -3, 4, 2}, want: 5},
		{name: "example 2", nums: []int{1, 2}, want: 1},
		{name: "example 3", nums: []int{1, -2, -3}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinStartValue(tt.nums); got != tt.want {
				t.Errorf("MinStartValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
