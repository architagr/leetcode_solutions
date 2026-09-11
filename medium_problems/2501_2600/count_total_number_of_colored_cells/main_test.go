package counttotalnumberofcoloredcells

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestColoredCells(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int64
	}{
		{name: "example 1", n: 1, want: 1},
		{name: "example 2", n: 2, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coloredCells(tt.n); got != tt.want {
				t.Errorf("coloredCells() = %v, want %v", got, tt.want)
			}
		})
	}
}
