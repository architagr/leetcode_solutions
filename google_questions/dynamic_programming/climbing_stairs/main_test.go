package climbingstairs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	input, expected int
}

func TestClimbStarirs(t *testing.T) {
	testcases := []testcase{
		{
			input:    2,
			expected: 2,
		},
		{
			input:    3,
			expected: 3,
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprint(tc.input), func(t *testing.T) {
			assert.Equal(t, tc.expected, climbStairs(tc.input))
		})
	}
}
