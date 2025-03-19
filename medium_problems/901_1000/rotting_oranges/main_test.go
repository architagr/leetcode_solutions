package rottingoranges

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	grid     [][]int
	expected int
}

func TestOrangesRotting(t *testing.T) {
	testcases := []testcase{
		{
			grid: [][]int{
				{2, 1, 1}, {1, 1, 0}, {0, 1, 1},
			},
			expected: 4,
		},
		{
			grid: [][]int{
				{2, 1, 1}, {0, 1, 1}, {1, 0, 1},
			},
			expected: -1,
		},
		{
			grid: [][]int{
				{2, 0},
			},
			expected: 0,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, orangesRotting(tc.grid))
		})
	}
}
