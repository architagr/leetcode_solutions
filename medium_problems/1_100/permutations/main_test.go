package permutations

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	nums     []int
	expected [][]int
}

func TestPermute(t *testing.T) {
	testcases := []testcase{
		{
			nums: []int{1, 2, 3},
			expected: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 2, 1},
				{3, 1, 2},
			},
		},
		{
			nums: []int{0, 1},
			expected: [][]int{
				{0, 1},
				{1, 0},
			},
		},
		{
			nums: []int{1},
			expected: [][]int{
				{1},
			},
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := permute(tc.nums)
			assert.Equal(t, tc.expected, got)
		})
	}
}
