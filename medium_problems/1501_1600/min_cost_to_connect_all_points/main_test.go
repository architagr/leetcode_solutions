package mincosttoconnectallpairs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	points   [][]int
	expected int
}

func TestMinCostConnectPoints(t *testing.T) {
	testcases := []testcase{
		{
			points: [][]int{
				{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0},
			},
			expected: 20,
		},
		{
			points: [][]int{
				{3, 12}, {-2, 5}, {-4, 1},
			},
			expected: 18,
		},
		{
			points: [][]int{
				{-1000000, -1000000}, {1000000, 1000000},
			},
			expected: 4000000,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, minCostConnectPoints(tc.points))
		})
	}
}
