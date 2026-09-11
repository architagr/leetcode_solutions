package pointsthatintersectwithcars

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestNumberOfPoints(t *testing.T) {
	tests := []struct {
		name string
		nums [][]int
		want int
	}{
		{name: "example 1", nums: [][]int{[]int{3, 6}, []int{1, 5}, []int{4, 7}}, want: 7},
		{name: "example 2", nums: [][]int{[]int{1, 3}, []int{5, 8}}, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfPoints(tt.nums); got != tt.want {
				t.Errorf("numberOfPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
