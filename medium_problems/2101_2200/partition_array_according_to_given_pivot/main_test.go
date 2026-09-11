package partition_array_according_given_pivot

import (
	"reflect"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode.
func TestPivotArray(t *testing.T) {
	tests := []struct {
		name  string
		nums  []int
		pivot int
		want  []int
	}{
		{name: "example 1", nums: []int{9, 12, 5, 10, 14, 3, 10}, pivot: 10, want: []int{9, 5, 3, 10, 10, 12, 14}},
		{name: "example 2", nums: []int{-3, 4, 3, 2}, pivot: 2, want: []int{-3, 2, 4, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotArray(tt.nums, tt.pivot); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pivotArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
