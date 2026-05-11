package besttimetobuyandsellstockusingstatergy

import (
	"fmt"
	"testing"
)

type testcases struct {
	price, statergy []int
	k               int
	expected        int64
}

func TestBestTimeToBuyAndSellStockUsingStateStrategy(t *testing.T) {
	testCases := []testcases{
		{
			price:    []int{4, 2, 8},
			statergy: []int{-1, 0, 1},
			k:        2,
			expected: 10,
		},
		{
			price:    []int{5, 4, 3},
			statergy: []int{1, 1, 0},
			k:        2,
			expected: 9,
		},
		{
			price:    []int{5, 14, 16, 9},
			statergy: []int{-1, 0, 0, -1},
			k:        2,
			expected: 5,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("testcase_%d", i), func(t *testing.T) {
			actual := maxProfit(tc.price, tc.statergy, tc.k)
			if actual != tc.expected {
				t.Errorf("expected %v but got %v for testcase %v", tc.expected, actual, i)
			}
		})
	}
}
