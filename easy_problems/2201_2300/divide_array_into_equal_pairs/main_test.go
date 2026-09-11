package divide_array_into_equal_pairs

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestDivideArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{name: "example 1", nums: []int{3, 2, 3, 2, 2, 2}, want: true},
		{name: "example 2", nums: []int{1, 2, 3, 4}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DivideArray(tt.nums); got != tt.want {
				t.Errorf("DivideArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
