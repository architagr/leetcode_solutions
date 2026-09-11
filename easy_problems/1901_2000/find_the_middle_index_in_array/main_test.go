package find_the_middle_index_in_array

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestFindMiddleIndex(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "example 1", nums: []int{2, 3, -1, 8, 4}, want: 3},
		{name: "example 2", nums: []int{1, -1, 4}, want: 2},
		{name: "example 3", nums: []int{2, 5}, want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMiddleIndex(tt.nums); got != tt.want {
				t.Errorf("FindMiddleIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
