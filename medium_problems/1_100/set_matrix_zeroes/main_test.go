package set_matrix_zeroes

import (
	"reflect"
	"testing"
)

// SetZeroes works in place, so the case compares the argument after the call.
// Cases are the worked examples from the problem statement on LeetCode.
func TestSetZeroes(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		{name: "example 1", matrix: [][]int{[]int{1, 1, 1}, []int{1, 0, 1}, []int{1, 1, 1}}, want: [][]int{[]int{1, 0, 1}, []int{0, 0, 0}, []int{1, 0, 1}}},
		{name: "example 2", matrix: [][]int{[]int{0, 1, 2, 0}, []int{3, 4, 5, 2}, []int{1, 3, 1, 5}}, want: [][]int{[]int{0, 0, 0, 0}, []int{0, 4, 5, 0}, []int{0, 3, 1, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arg := make([][]int, len(tt.matrix))
			for i := range tt.matrix {
				arg[i] = append([]int{}, tt.matrix[i]...)
			}
			SetZeroes(arg)
			got := arg
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetZeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}
