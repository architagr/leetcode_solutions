package countnumberofdistinctintegersafterreverseoperations

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestCountDistinctIntegers(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "example 1", nums: []int{1, 13, 10, 12, 31}, want: 6},
		{name: "example 2", nums: []int{2, 2, 2}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countDistinctIntegers(tt.nums); got != tt.want {
				t.Errorf("countDistinctIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}
