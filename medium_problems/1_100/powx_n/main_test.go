package powxn

import (
	"math"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode.
func TestMyPow(t *testing.T) {
	tests := []struct {
		name string
		x    float64
		n    int
		want float64
	}{
		{name: "example 1", x: 2.0, n: 10, want: 1024.0},
		{name: "example 2", x: 2.1, n: 3, want: 9.261},
		{name: "example 3", x: 2.0, n: -2, want: 0.25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.x, tt.n); math.Abs(got-tt.want) > 1e-5 {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
