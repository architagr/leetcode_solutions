package contiguousarray

import (
	"fmt"
	"testing"
)

type TestCase struct {
	nums     []int
	expected int
}

func TestFindMaxLength(t *testing.T) {
	testCases := []TestCase{
		{
			nums:     []int{0, 1, 0},
			expected: 2,
		},
		{
			nums:     []int{1, 0, 1, 0, 0, 1, 1, 1},
			expected: 6,
		},
		{
			nums:     []int{0, 0, 1, 0, 0, 1, 1, 1},
			expected: 8,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t1 *testing.T) {
			if got := findMaxLength(tc.nums); got != tc.expected {
				t1.Errorf("expected %d but got %d", tc.expected, got)
			}
		})
	}
}
