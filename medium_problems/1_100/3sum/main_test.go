package sum3

import (
	"reflect"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode.
func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{name: "example 1", nums: []int{-1, 0, 1, 2, -1, -4}, want: [][]int{[]int{-1, -1, 2}, []int{-1, 0, 1}}},
		{name: "example 2", nums: []int{0, 1, 1}, want: [][]int{}},
		{name: "example 3", nums: []int{0, 0, 0}, want: [][]int{[]int{0, 0, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ThreeSum(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ThreeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
