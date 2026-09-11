package toeplitz_matrix

import "testing"

// Cases are the worked examples from the problem statement on LeetCode.
func TestIsToePlitzMatrix(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   bool
	}{
		{name: "example 1", matrix: [][]int{[]int{1, 2, 3, 4}, []int{5, 1, 2, 3}, []int{9, 5, 1, 2}}, want: true},
		{name: "example 2", matrix: [][]int{[]int{1, 2}, []int{2, 2}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsToePlitzMatrix(tt.matrix); got != tt.want {
				t.Errorf("IsToePlitzMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
