package findthescoreofallprefixesofanarray

import (
	"reflect"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode.
func TestFindPrefixScore(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int64
	}{
		{name: "example 1", nums: []int{2, 3, 7, 5, 10}, want: []int64{4, 10, 24, 36, 56}},
		{name: "example 2", nums: []int{1, 1, 2, 4, 8, 16}, want: []int64{2, 4, 8, 16, 32, 64}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPrefixScore(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findPrefixScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
