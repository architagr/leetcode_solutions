package applyoperationstoanarray

import (
	"reflect"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode.
func TestApplyOperations(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{name: "example 1", nums: []int{1, 2, 2, 1, 1, 0}, want: []int{1, 4, 2, 0, 0, 0}},
		{name: "example 2", nums: []int{0, 1}, want: []int{1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := applyOperations(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("applyOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
