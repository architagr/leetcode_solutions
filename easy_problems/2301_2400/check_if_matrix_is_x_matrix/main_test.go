package x_matrix

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestCheckXMatrix(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want bool
	}{
		{name: "example 1", grid: [][]int{[]int{2, 0, 0, 1}, []int{0, 3, 1, 0}, []int{0, 5, 2, 0}, []int{4, 0, 0, 2}}, want: true},
		{name: "example 2", grid: [][]int{[]int{5, 7, 0}, []int{0, 3, 1}, []int{0, 5, 0}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkXMatrix(tt.grid); got != tt.want {
				t.Errorf("checkXMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
