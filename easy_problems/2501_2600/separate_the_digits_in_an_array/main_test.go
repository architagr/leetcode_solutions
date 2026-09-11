package separatethedigitsinanarray

import (
	"reflect"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode.
func TestSeparateDigits(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{name: "example 1", nums: []int{13, 25, 83, 77}, want: []int{1, 3, 2, 5, 8, 3, 7, 7}},
		{name: "example 2", nums: []int{7, 1, 3, 9}, want: []int{7, 1, 3, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := separateDigits(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("separateDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
