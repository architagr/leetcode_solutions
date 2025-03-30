package besttimetobuyandsellstockiv

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	k        int
	prices   []int
	expected int
}

func TestMaxProfit(t *testing.T) {
	testcases := []testcase{
		{
			k:        2,
			prices:   []int{2, 4, 1},
			expected: 2,
		},
		{
			k:        2,
			prices:   []int{3, 2, 6, 5, 0, 3},
			expected: 7,
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, maxProfit(tc.k, tc.prices))
		})
	}
}
