package closestequalelementqueries

import (
	"reflect"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode.
func TestSolveQueries(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		queries []int
		want    []int
	}{
		{name: "example 1", nums: []int{1, 3, 1, 4, 1, 3, 2}, queries: []int{0, 3, 5}, want: []int{2, -1, 3}},
		{name: "example 2", nums: []int{1, 2, 3, 4}, queries: []int{0, 1, 2, 3}, want: []int{-1, -1, -1, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveQueries(tt.nums, tt.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solveQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
