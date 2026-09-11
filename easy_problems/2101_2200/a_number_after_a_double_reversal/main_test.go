package number_after_a_double_reversal

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestIsSameAfterReversals(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want bool
	}{
		{name: "example 1", num: 526, want: true},
		{name: "example 2", num: 1800, want: false},
		{name: "example 3", num: 0, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSameAfterReversals(tt.num); got != tt.want {
				t.Errorf("IsSameAfterReversals() = %v, want %v", got, tt.want)
			}
		})
	}
}
