package findpeakelement

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestFindPeakElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "example 1", nums: []int{1, 2, 3, 1}, want: 2},
		{name: "example 2", nums: []int{1, 2, 1, 3, 5, 6, 4}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPeakElement(tt.nums); got != tt.want {
				t.Errorf("findPeakElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
