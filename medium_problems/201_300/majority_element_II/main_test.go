package majority_element_II

import (
	"reflect"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode.
func TestMajorityElementII(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{name: "example 1", nums: []int{3, 2, 3}, want: []int{3}},
		{name: "example 2", nums: []int{1}, want: []int{1}},
		{name: "example 3", nums: []int{1, 2}, want: []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MajorityElementII(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MajorityElementII() = %v, want %v", got, tt.want)
			}
		})
	}
}
