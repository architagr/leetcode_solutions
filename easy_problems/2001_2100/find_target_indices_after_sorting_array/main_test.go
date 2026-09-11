package target_Indices_after_sorting_array

import (
	"reflect"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode.
func TestTargetIndices(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{name: "example 1", nums: []int{1, 2, 5, 2, 3}, target: 2, want: []int{1, 2}},
		{name: "example 2", nums: []int{1, 2, 5, 2, 3}, target: 3, want: []int{3}},
		{name: "example 3", nums: []int{1, 2, 5, 2, 3}, target: 5, want: []int{4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TargetIndices(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TargetIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}
