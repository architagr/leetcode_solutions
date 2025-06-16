package maximumdifferencebetweenincreasingelements

import (
	"fmt"
	"testing"
)

func TestMaximumDifference(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{7, 1, 5, 4}, 4},
		{[]int{9, 4, 3, 2}, -1},
		{[]int{1, 2, 3, 4}, 3},
		{[]int{1, 2, 3, 0}, 2},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if got := maximumDifference(tt.nums); got != tt.want {
				t.Errorf("maximumDifference(%v) = %v; want %v", tt.nums, got, tt.want)
			}
		})
	}
}
