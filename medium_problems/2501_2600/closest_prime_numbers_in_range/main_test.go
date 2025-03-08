package closestprimenumbersinrange

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	left, right int
	expected    []int
}

func TestClosestPrimes(t *testing.T) {
	testcases := []testcase{
		{
			left:     10,
			right:    19,
			expected: []int{11, 13},
		},
		{
			left:     4,
			right:    6,
			expected: []int{-1, -1},
		},
		{
			left:     1,
			right:    1000000,
			expected: []int{2, 3},
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := closestPrimes(tc.left, tc.right)
			assert.Equal(t, tc.expected, got)
		})
	}
}
