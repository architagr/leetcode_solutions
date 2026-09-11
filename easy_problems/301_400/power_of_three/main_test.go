package power_of_three

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestIsPowerOfThree(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{name: "example 1", n: 27, want: true},
		{name: "example 2", n: 0, want: false},
		{name: "example 3", n: -1, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfThree(tt.n); got != tt.want {
				t.Errorf("isPowerOfThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
