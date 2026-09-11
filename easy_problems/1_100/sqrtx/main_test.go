package sqrtx

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestMySqrt(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{name: "example 1", x: 4, want: 2},
		{name: "example 2", x: 8, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
