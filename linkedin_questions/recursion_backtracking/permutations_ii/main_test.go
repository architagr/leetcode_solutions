package permutationsii

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
			nums: []int{1, 1, 2},
			expected: [][]int{
				{1, 1, 2},
				{1, 2, 1},
				{2, 1, 1},
			},
		},
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
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := permuteUnique(tc.nums)
			assert.Equal(t, tc.expected, got)
		})
	}
}
