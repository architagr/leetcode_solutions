package reverse_pairs

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestReversePairs(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "example 1", nums: []int{1, 3, 2, 3, 1}, want: 2},
		{name: "example 2", nums: []int{2, 4, 3, 5, 1}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReversePairs(tt.nums); got != tt.want {
				t.Errorf("ReversePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
