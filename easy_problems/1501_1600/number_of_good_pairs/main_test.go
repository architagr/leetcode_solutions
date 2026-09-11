package number_of_good_pairs

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestCountGoodPairs(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "example 1", nums: []int{1, 2, 3, 1, 1, 3}, want: 4},
		{name: "example 2", nums: []int{1, 1, 1, 1}, want: 6},
		{name: "example 3", nums: []int{1, 2, 3}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountGoodPairs(tt.nums); got != tt.want {
				t.Errorf("CountGoodPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
