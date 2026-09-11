package top_k_frequent_elements

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

// normalize sorts the result, since this problem accepts any order.
func normalize(v []int) []int {
	out := append([]int{}, v...)
	sort.Slice(out, func(a, b int) bool { return fmt.Sprint(out[a]) < fmt.Sprint(out[b]) })
	return out
}

// Cases are the worked examples from the problem statement on LeetCode.
func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{name: "example 1", nums: []int{1, 1, 1, 2, 2, 3}, k: 2, want: []int{1, 2}},
		{name: "example 2", nums: []int{1}, k: 1, want: []int{1}},
		{name: "example 3", nums: []int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2}, k: 2, want: []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TopKFrequent(tt.nums, tt.k); !reflect.DeepEqual(normalize(got), normalize(tt.want)) {
				t.Errorf("TopKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
