package minimumfallingpathsum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	matrix   [][]int
	expected int
}

func TestMinFallingPathSum(t *testing.T) {
	testcaes := []testcase{
		{
			matrix:   [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}},
			expected: 13,
		},
		{
			matrix:   [][]int{{-19, 57}, {-40, -5}},
			expected: -59,
		},
	}

	for i, tc := range testcaes {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, minFallingPathSum(tc.matrix))
		})

	}
}
