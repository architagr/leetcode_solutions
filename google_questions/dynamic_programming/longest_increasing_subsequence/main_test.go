package longestincreasingsubsequence

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	nums     []int
	expected int
}

func TestWordBreak(t *testing.T) {
	testcases := []testcase{
		{
			nums:     []int{10, 9, 2, 5, 3, 7, 101, 18},
			expected: 4,
		},
		{
			nums:     []int{0, 1, 0, 3, 2, 3},
			expected: 4,
		},
		{
			nums:     []int{7, 7, 7, 7, 7},
			expected: 1,
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, lengthOfLIS(tc.nums))
		})
	}
}
