package minimumpathsum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	grid     [][]int
	expected int
}

func TestMinPathSum(t *testing.T) {
	testcaes := []testcase{
		{
			grid:     [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
			expected: 7,
		},
		{
			grid:     [][]int{{1, 2, 3}, {4, 5, 6}},
			expected: 12,
		},
	}

	for i, tc := range testcaes {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, minPathSum(tc.grid))
		})

	}
}
