package numberofbeautifulpairs

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestCountBeautifulPairs(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "example 1", nums: []int{2, 5, 1, 4}, want: 5},
		{name: "example 2", nums: []int{11, 21, 12}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBeautifulPairs(tt.nums); got != tt.want {
				t.Errorf("countBeautifulPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
