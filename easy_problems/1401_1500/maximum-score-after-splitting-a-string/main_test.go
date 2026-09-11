package maximumscoreaftersplittingastring

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestMaxScore(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{name: "example 1", s: "011101", want: 5},
		{name: "example 2", s: "00111", want: 5},
		{name: "example 3", s: "1111", want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScore(tt.s); got != tt.want {
				t.Errorf("maxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
