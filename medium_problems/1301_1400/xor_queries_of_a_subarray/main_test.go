package xor_queries_of_a_subarray

import (
	"reflect"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode.
func TestXorQueries(t *testing.T) {
	tests := []struct {
		name    string
		arr     []int
		queries [][]int
		want    []int
	}{
		{name: "example 1", arr: []int{1, 3, 4, 8}, queries: [][]int{[]int{0, 1}, []int{1, 2}, []int{0, 3}, []int{3, 3}}, want: []int{2, 7, 14, 8}},
		{name: "example 2", arr: []int{4, 8, 2, 10}, queries: [][]int{[]int{2, 3}, []int{1, 3}, []int{0, 0}, []int{0, 3}}, want: []int{8, 0, 4, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := XorQueries(tt.arr, tt.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XorQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
