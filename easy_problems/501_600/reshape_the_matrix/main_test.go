package reshape_the_matrix

import (
	"reflect"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode.
func TestMatrixReshape(t *testing.T) {
	tests := []struct {
		name string
		mat  [][]int
		r    int
		c    int
		want [][]int
	}{
		{name: "example 1", mat: [][]int{[]int{1, 2}, []int{3, 4}}, r: 1, c: 4, want: [][]int{[]int{1, 2, 3, 4}}},
		{name: "example 2", mat: [][]int{[]int{1, 2}, []int{3, 4}}, r: 2, c: 4, want: [][]int{[]int{1, 2}, []int{3, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatrixReshape(tt.mat, tt.r, tt.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatrixReshape() = %v, want %v", got, tt.want)
			}
		})
	}
}
