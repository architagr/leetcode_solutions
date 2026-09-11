package shufflethearray

import (
	"reflect"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode.
func TestShuffle(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		n    int
		want []int
	}{
		{name: "example 1", nums: []int{2, 5, 1, 3, 4, 7}, n: 3, want: []int{2, 3, 5, 4, 1, 7}},
		{name: "example 2", nums: []int{1, 2, 3, 4, 4, 3, 2, 1}, n: 4, want: []int{1, 4, 2, 3, 3, 2, 4, 1}},
		{name: "example 3", nums: []int{1, 1, 2, 2}, n: 2, want: []int{1, 2, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shuffle(tt.nums, tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shuffle() = %v, want %v", got, tt.want)
			}
		})
	}
}
