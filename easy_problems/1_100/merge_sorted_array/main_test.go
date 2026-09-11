package mergesortedarray

import (
	"reflect"
	"testing"
)

// merge works in place, so the case compares the argument after the call.
// Cases are the worked examples from the problem statement on LeetCode.
func TestMerge(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{name: "example 1", nums1: []int{1, 2, 3, 0, 0, 0}, m: 3, nums2: []int{2, 5, 6}, n: 3, want: []int{1, 2, 2, 3, 5, 6}},
		{name: "example 2", nums1: []int{1}, m: 1, nums2: []int{}, n: 0, want: []int{1}},
		{name: "example 3", nums1: []int{0}, m: 0, nums2: []int{1}, n: 1, want: []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arg := append([]int{}, tt.nums1...)
			merge(arg, tt.m, tt.nums2, tt.n)
			got := arg
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
