package find_all_numbers_disappeared_in_an_array

import (
	"reflect"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode.
func TestFindDisappearedNumbers(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{name: "example 1", nums: []int{4, 3, 2, 7, 8, 2, 3, 1}, want: []int{5, 6}},
		{name: "example 2", nums: []int{1, 1}, want: []int{2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindDisappearedNumbers(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindDisappearedNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
