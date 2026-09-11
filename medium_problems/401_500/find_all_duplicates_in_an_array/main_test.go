package find_all_duplicates_in_an_array

import (
	"reflect"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode.
func TestFindDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{name: "example 1", nums: []int{4, 3, 2, 7, 8, 2, 3, 1}, want: []int{2, 3}},
		{name: "example 2", nums: []int{1, 1, 2}, want: []int{1}},
		{name: "example 3", nums: []int{1}, want: []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindDuplicates(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
