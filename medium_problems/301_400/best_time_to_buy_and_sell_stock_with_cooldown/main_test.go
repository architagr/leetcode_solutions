package besttimetobuyandsellstockwithcooldown

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	prices   []int
	expected int
}

func TestMaxProfit(t *testing.T) {
	testcases := []testcase{
		{
			prices:   []int{1, 2, 3, 0, 2},
			expected: 3,
		},
		{
			prices:   []int{1},
			expected: 0,
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, maxProfit(tc.prices))
		})
	}
}
