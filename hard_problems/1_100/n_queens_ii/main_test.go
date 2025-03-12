package nqueensii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	n        int
	expected int
}

func TestSolveNQueens(t *testing.T) {
	testcases := []testcase{
		{
			n:        1,
			expected: 1,
		},
		{
			n:        4,
			expected: 2,
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf("%d_queen", tc.n), func(t *testing.T) {
			assert.EqualValues(t, tc.expected, totalNQueens(tc.n))
		})
	}
}
