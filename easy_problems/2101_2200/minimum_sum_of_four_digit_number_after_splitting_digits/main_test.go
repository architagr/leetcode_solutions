package minimum_sum_of_four_digit_number_after_splitting_digits

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestMinimumSum(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want int
	}{
		{name: "example 1", num: 2932, want: 52},
		{name: "example 2", num: 4009, want: 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinimumSum(tt.num); got != tt.want {
				t.Errorf("MinimumSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
