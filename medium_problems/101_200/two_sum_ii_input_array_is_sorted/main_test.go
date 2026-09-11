package twosumiiinputarrayissorted

import (
	"reflect"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode.
func TestTwoSum(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		target  int
		want    []int
	}{
		{name: "example 1", numbers: []int{2, 7, 11, 15}, target: 9, want: []int{1, 2}},
		{name: "example 2", numbers: []int{2, 3, 4}, target: 6, want: []int{1, 3}},
		{name: "example 3", numbers: []int{-1, 0}, target: -1, want: []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.numbers, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
