package contains_duplicate_2

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestContainsDuplicate2(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want bool
	}{
		{name: "example 1", nums: []int{1, 2, 3, 1}, k: 3, want: true},
		{name: "example 2", nums: []int{1, 0, 1, 1}, k: 1, want: true},
		{name: "example 3", nums: []int{1, 2, 3, 1, 2, 3}, k: 2, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsDuplicate2(tt.nums, tt.k); got != tt.want {
				t.Errorf("containsDuplicate2() = %v, want %v", got, tt.want)
			}
		})
	}
}
