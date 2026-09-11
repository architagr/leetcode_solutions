package merge_similar_items

import (
	"reflect"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode.
func TestMergeSimilarItems(t *testing.T) {
	tests := []struct {
		name   string
		items1 [][]int
		items2 [][]int
		want   [][]int
	}{
		{name: "example 1", items1: [][]int{[]int{1, 1}, []int{4, 5}, []int{3, 8}}, items2: [][]int{[]int{3, 1}, []int{1, 5}}, want: [][]int{[]int{1, 6}, []int{3, 9}, []int{4, 5}}},
		{name: "example 2", items1: [][]int{[]int{1, 1}, []int{3, 2}, []int{2, 3}}, items2: [][]int{[]int{2, 1}, []int{3, 2}, []int{1, 3}}, want: [][]int{[]int{1, 4}, []int{2, 4}, []int{3, 4}}},
		{name: "example 3", items1: [][]int{[]int{1, 3}, []int{2, 2}}, items2: [][]int{[]int{7, 1}, []int{2, 2}, []int{1, 4}}, want: [][]int{[]int{1, 7}, []int{2, 4}, []int{7, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSimilarItems(tt.items1, tt.items2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSimilarItems() = %v, want %v", got, tt.want)
			}
		})
	}
}
