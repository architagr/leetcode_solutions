package rotate_image

import (
	"reflect"
	"testing"
)

// Rotate works in place, so the case compares the argument after the call.
// Cases are the worked examples from the problem statement on LeetCode.
func TestRotate(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		{name: "example 1", matrix: [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}, want: [][]int{[]int{7, 4, 1}, []int{8, 5, 2}, []int{9, 6, 3}}},
		{name: "example 2", matrix: [][]int{[]int{5, 1, 9, 11}, []int{2, 4, 8, 10}, []int{13, 3, 6, 7}, []int{15, 14, 12, 16}}, want: [][]int{[]int{15, 13, 2, 5}, []int{14, 3, 4, 1}, []int{12, 6, 8, 9}, []int{16, 7, 10, 11}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arg := make([][]int, len(tt.matrix))
			for i := range tt.matrix {
				arg[i] = append([]int{}, tt.matrix[i]...)
			}
			Rotate(arg)
			got := arg
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rotate() = %v, want %v", got, tt.want)
			}
		})
	}
}
