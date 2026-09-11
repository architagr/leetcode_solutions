package maximumscorefromperformingmultiplicationoperations

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestMaximumScore(t *testing.T) {
	tests := []struct {
		name        string
		nums        []int
		multipliers []int
		want        int
	}{
		{name: "example 1", nums: []int{1, 2, 3}, multipliers: []int{3, 2, 1}, want: 14},
		{name: "example 2", nums: []int{-5, -3, -3, -2, 7, 1}, multipliers: []int{-10, -5, 3, 4, 6}, want: 102},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumScore(tt.nums, tt.multipliers); got != tt.want {
				t.Errorf("maximumScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
