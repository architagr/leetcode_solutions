package sumofdistancesintree

import (
	"reflect"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode.
func TestSumOfDistancesInTree(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		edges [][]int
		want  []int
	}{
		{name: "example 1", n: 6, edges: [][]int{[]int{0, 1}, []int{0, 2}, []int{2, 3}, []int{2, 4}, []int{2, 5}}, want: []int{8, 12, 6, 10, 10, 10}},
		{name: "example 2", n: 1, edges: [][]int{}, want: []int{0}},
		{name: "example 3", n: 2, edges: [][]int{[]int{1, 0}}, want: []int{1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfDistancesInTree(tt.n, tt.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumOfDistancesInTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
