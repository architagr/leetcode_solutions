package perfectsquares

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestNumSquares(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{name: "example 1", n: 12, want: 3},
		{name: "example 2", n: 13, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSquares(tt.n); got != tt.want {
				t.Errorf("numSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
