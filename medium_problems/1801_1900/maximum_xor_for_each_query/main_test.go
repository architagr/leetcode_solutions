package maximum_xor_for_each_query

import (
	"reflect"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode.
func TestGetMaximumXor(t *testing.T) {
	tests := []struct {
		name       string
		nums       []int
		maximumBit int
		want       []int
	}{
		{name: "example 1", nums: []int{0, 1, 1, 3}, maximumBit: 2, want: []int{0, 3, 2, 3}},
		{name: "example 2", nums: []int{2, 3, 4, 7}, maximumBit: 3, want: []int{5, 2, 6, 5}},
		{name: "example 3", nums: []int{0, 1, 2, 2, 5, 7}, maximumBit: 3, want: []int{4, 3, 6, 4, 6, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMaximumXor(tt.nums, tt.maximumBit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMaximumXor() = %v, want %v", got, tt.want)
			}
		})
	}
}
