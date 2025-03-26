package houserobber

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	nums     []int
	expected int
}

func TestRob(t *testing.T) {
	testcases := []testcase{
		{
			nums:     []int{1, 2, 3, 1},
			expected: 4,
		},
		{
			nums:     []int{2, 7, 9, 3, 1},
			expected: 12,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, rob(tc.nums))
		})
	}
}
