package singlenumberii

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "example 1", nums: []int{2, 2, 3, 2}, want: 3},
		{name: "example 2", nums: []int{0, 1, 0, 1, 0, 1, 99}, want: 99},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
