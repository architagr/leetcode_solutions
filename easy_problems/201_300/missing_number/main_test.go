package missing_number

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestMissingNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "example 1", nums: []int{3, 0, 1}, want: 2},
		{name: "example 2", nums: []int{0, 1}, want: 2},
		{name: "example 3", nums: []int{9, 6, 4, 2, 3, 5, 7, 0, 1}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MissingNumber(tt.nums); got != tt.want {
				t.Errorf("MissingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
