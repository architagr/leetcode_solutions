package maximalsquare

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	matrix   [][]byte
	expected int
}

func TestMaximalSquare(t *testing.T) {
	testcases := []testcase{
		{
			matrix: [][]byte{
				{'1', '0', '1', '1', '0', '1'},
				{'1', '1', '1', '1', '1', '1'},
				{'0', '1', '1', '0', '1', '1'},
				{'1', '1', '1', '0', '1', '0'},
				{'0', '1', '1', '1', '1', '1'},
				{'1', '1', '0', '1', '1', '1'},
			},
			expected: 4,
		},
		{
			matrix: [][]byte{
				{'0'},
			},
			expected: 0,
		},
		{
			matrix: [][]byte{
				{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'},
			},
			expected: 4,
		},
		{
			matrix: [][]byte{
				{'1', '0', '1', '1', '1'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'},
			},
			expected: 9,
		},
		{
			matrix: [][]byte{
				{'0', '1'}, {'1', '0'},
			},
			expected: 1,
		},
		{
			matrix: [][]byte{
				{'1', '1'}, {'1', '0'},
			},
			expected: 1,
		},
		{
			matrix: [][]byte{
				{'1', '1'}, {'1', '1'},
			},
			expected: 4,
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, maximalSquare(tc.matrix))
		})
	}
}
