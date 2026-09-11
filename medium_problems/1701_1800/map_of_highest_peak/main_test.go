package mapofhighestpeak

import (
	"reflect"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode.
func TestHighestPeak(t *testing.T) {
	tests := []struct {
		name    string
		isWater [][]int
		want    [][]int
	}{
		{name: "example 1", isWater: [][]int{[]int{0, 1}, []int{0, 0}}, want: [][]int{[]int{1, 0}, []int{2, 1}}},
		{name: "example 2", isWater: [][]int{[]int{0, 0, 1}, []int{1, 0, 0}, []int{0, 0, 0}}, want: [][]int{[]int{1, 1, 0}, []int{0, 1, 1}, []int{1, 2, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := highestPeak(tt.isWater); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("highestPeak() = %v, want %v", got, tt.want)
			}
		})
	}
}
