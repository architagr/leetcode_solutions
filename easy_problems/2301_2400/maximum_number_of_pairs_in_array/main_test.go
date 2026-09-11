package maximum_number_of_pairs_in_array

import (
	"reflect"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode.
func TestNumberOfPairs(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{name: "example 1", nums: []int{1, 3, 2, 1, 3, 2, 2}, want: []int{3, 1}},
		{name: "example 2", nums: []int{1, 1}, want: []int{1, 0}},
		{name: "example 3", nums: []int{0}, want: []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumberOfPairs(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NumberOfPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
