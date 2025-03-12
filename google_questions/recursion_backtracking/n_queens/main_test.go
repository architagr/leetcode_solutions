package nqueens

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	n        int
	expected [][]string
}

func TestSolveNQueens(t *testing.T) {
	testcases := []testcase{
		{
			n: 1,
			expected: [][]string{
				{"Q"},
			},
		},
		{
			n: 4,
			expected: [][]string{
				{
					".Q..",
					"...Q",
					"Q...",
					"..Q.",
				},
				{
					"..Q.",
					"Q...",
					"...Q",
					".Q..",
				},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf("%d_queen", tc.n), func(t *testing.T) {
			assert.EqualValues(t, tc.expected, solveNQueens(tc.n))
		})
	}
}
