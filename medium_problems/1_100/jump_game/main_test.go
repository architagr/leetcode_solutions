package jumpgame

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	nums     []int
	expected bool
}

func TestCanJump(t *testing.T) {
	testcases := []testcase{
		{
			nums:     []int{2, 3, 1, 1, 4},
			expected: true,
		},
		{
			nums:     []int{3, 2, 1, 0, 4},
			expected: false,
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := canJump(tc.nums)
			assert.Equal(t, tc.expected, got)
		})
	}
}
