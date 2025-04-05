package singleelementinasortedarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	nums     []int
	expected int
}

func TestSingleNonDuplicate(t *testing.T) {
	testcases := []testcase{
		{
			nums:     []int{1, 1, 2, 2, 3, 3, 4},
			expected: 4,
		},
		{
			nums:     []int{1, 2, 2, 3, 3},
			expected: 1,
		},
		{
			nums:     []int{1},
			expected: 1,
		},
		{
			nums:     []int{1, 1, 2, 3, 3, 4, 4, 8, 8},
			expected: 2,
		},
		{
			nums:     []int{3, 3, 7, 7, 10, 11, 11},
			expected: 10,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, singleNonDuplicate(tc.nums))
		})
	}
}
