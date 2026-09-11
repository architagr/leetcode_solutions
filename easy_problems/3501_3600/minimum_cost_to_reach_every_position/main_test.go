package minimumcosttoreacheveryposition

import (
	"reflect"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode.
func TestMinCosts(t *testing.T) {
	tests := []struct {
		name string
		cost []int
		want []int
	}{
		{name: "example 1", cost: []int{5, 3, 4, 1, 3, 2}, want: []int{5, 3, 3, 1, 1, 1}},
		{name: "example 2", cost: []int{1, 2, 4, 6, 7}, want: []int{1, 1, 1, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCosts(tt.cost); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minCosts() = %v, want %v", got, tt.want)
			}
		})
	}
}
