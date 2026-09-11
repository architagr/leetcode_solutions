package sumofuniqueelements

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestSumOfUnique(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "example 1", nums: []int{1, 2, 3, 2}, want: 4},
		{name: "example 2", nums: []int{1, 1, 1, 1, 1}, want: 0},
		{name: "example 3", nums: []int{1, 2, 3, 4, 5}, want: 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfUnique(tt.nums); got != tt.want {
				t.Errorf("sumOfUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
