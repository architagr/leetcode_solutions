package majority_element

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "example 1", nums: []int{3, 2, 3}, want: 3},
		{name: "example 2", nums: []int{2, 2, 1, 1, 1, 2, 2}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MajorityElement(tt.nums); got != tt.want {
				t.Errorf("MajorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
