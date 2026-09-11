package sumofdistances

import (
	"reflect"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode.
func TestDistance(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int64
	}{
		{name: "example 1", nums: []int{1, 3, 1, 1, 2}, want: []int64{5, 0, 3, 4, 0}},
		{name: "example 2", nums: []int{0, 5, 3}, want: []int64{0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distance(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("distance() = %v, want %v", got, tt.want)
			}
		})
	}
}
