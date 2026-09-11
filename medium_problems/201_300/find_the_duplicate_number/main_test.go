package find_the_duplicate_number

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestFindDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "example 1", nums: []int{1, 3, 4, 2, 2}, want: 2},
		{name: "example 2", nums: []int{3, 1, 3, 4, 2}, want: 3},
		{name: "example 3", nums: []int{3, 3, 3, 3, 3}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicate(tt.nums); got != tt.want {
				t.Errorf("findDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
