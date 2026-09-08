package sortcolors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	nums     []int
	expected []int
}

func TestSortColors(t *testing.T) {
	testcases := []testcase{
		{
			nums:     []int{1, 0, 0},
			expected: []int{0, 0, 1},
		},
		{
			nums:     []int{1, 0},
			expected: []int{0, 1},
		},
		{
			nums:     []int{1, 1},
			expected: []int{1, 1},
		},
		{
			nums:     []int{2, 2},
			expected: []int{2, 2},
		},
		{
			nums:     []int{2, 1},
			expected: []int{1, 2},
		},
		{
			nums:     []int{0, 0},
			expected: []int{0, 0},
		},
		{
			nums:     []int{2, 0, 2, 1, 1, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			nums:     []int{2, 0, 1},
			expected: []int{0, 1, 2},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			sortColors(tc.nums)
			assert.Equal(t, tc.expected, tc.nums)
		})
	}
}
