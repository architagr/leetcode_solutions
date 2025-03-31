package coinchange2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	amount   int
	coins    []int
	expected int
}

func TestChange(t *testing.T) {
	testcases := []testcase{
		{
			amount:   5,
			coins:    []int{1, 2, 5},
			expected: 4,
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, change(tc.amount, tc.coins))
		})
	}
}
