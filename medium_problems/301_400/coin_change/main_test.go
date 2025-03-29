package coinchange

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	coins    []int
	amount   int
	expected int
}

func TestCoinChange(t *testing.T) {
	testcases := []testcase{
		{
			coins:    []int{1, 2, 5},
			amount:   11,
			expected: 3,
		},
		{
			coins:    []int{2},
			amount:   3,
			expected: -1,
		},
		{
			coins:    []int{1},
			amount:   0,
			expected: 0,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, coinChange(tc.coins, tc.amount))
		})
	}
}
