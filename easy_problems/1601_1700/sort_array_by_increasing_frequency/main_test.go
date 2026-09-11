package sort_array_by_increasing_frequency

import (
	"reflect"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode.
func TestFrequencySort(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{name: "example 1", nums: []int{1, 1, 2, 2, 2, 3}, want: []int{3, 1, 1, 2, 2, 2}},
		{name: "example 2", nums: []int{2, 3, 1, 3, 2}, want: []int{1, 3, 3, 2, 2}},
		{name: "example 3", nums: []int{-1, 1, -6, 4, 5, -6, 1, 4, 1}, want: []int{5, -1, 4, 4, -6, -6, 1, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FrequencySort(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FrequencySort() = %v, want %v", got, tt.want)
			}
		})
	}
}
