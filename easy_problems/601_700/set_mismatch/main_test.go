package set_mismatch

import (
	"reflect"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode.
func TestFindErrorNums(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{name: "example 1", nums: []int{1, 2, 2, 4}, want: []int{2, 3}},
		{name: "example 2", nums: []int{1, 1}, want: []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindErrorNums(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindErrorNums() = %v, want %v", got, tt.want)
			}
		})
	}
}
