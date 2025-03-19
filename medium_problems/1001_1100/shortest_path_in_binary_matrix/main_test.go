package shortestpathinbinarymatrix

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	grid     [][]int
	expected int
}

func TestShortestPathBinaryMatrix(t *testing.T) {
	testcases := []testcase{
		{
			grid: [][]int{
				{0, 1}, {1, 0},
			},
			expected: 2,
		},
		{
			grid: [][]int{
				{0, 0, 0}, {1, 1, 0}, {1, 1, 0},
			},
			expected: 4,
		},
		{
			grid: [][]int{
				{1, 0, 0}, {1, 1, 0}, {1, 1, 0},
			},
			expected: -1,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, shortestPathBinaryMatrix(tc.grid))
		})
	}
}
