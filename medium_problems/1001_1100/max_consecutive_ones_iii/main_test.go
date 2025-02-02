package maxconsecutiveonesiii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	nums     []int
	k        int
	expected int
}

func TestLongestOnes(t *testing.T) {
	testcases := []testCase{
		{
			nums:     []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
			k:        2,
			expected: 6,
		},
		{
			nums:     []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
			k:        3,
			expected: 10,
		},
		{
			nums:     []int{0, 0, 0, 1},
			k:        4,
			expected: 4,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, longestOnes(tc.nums, tc.k))
		})
	}
}
