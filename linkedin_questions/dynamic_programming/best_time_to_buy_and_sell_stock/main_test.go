package besttimetobuyandsellstock

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
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			prices:   []int{7, 6, 5, 4, 3, 1},
			expected: 0,
		},
		{
			prices:   []int{1, 2, 3, 4, 5, 6, 7},
			expected: 6,
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, maxProfit(tc.prices))
		})
	}
}
