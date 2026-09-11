package p01matrix

import (
	"reflect"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode.
func TestUpdateMatrix(t *testing.T) {
	tests := []struct {
		name string
		mat  [][]int
		want [][]int
	}{
		{name: "example 1", mat: [][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{0, 0, 0}}, want: [][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{0, 0, 0}}},
		{name: "example 2", mat: [][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{1, 1, 1}}, want: [][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{1, 2, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := updateMatrix(tt.mat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
