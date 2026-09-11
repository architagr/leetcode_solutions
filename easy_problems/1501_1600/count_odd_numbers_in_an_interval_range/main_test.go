package countoddnumbersinanintervalrange

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestCountOdds(t *testing.T) {
	tests := []struct {
		name string
		low  int
		high int
		want int
	}{
		{name: "example 1", low: 3, high: 7, want: 3},
		{name: "example 2", low: 8, high: 10, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOdds(tt.low, tt.high); got != tt.want {
				t.Errorf("countOdds() = %v, want %v", got, tt.want)
			}
		})
	}
}
