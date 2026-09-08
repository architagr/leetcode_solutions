package uniquepathsii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	obstacleGrid [][]int
	expected     int
}

func TestUniquePaths(t *testing.T) {
	testcaes := []testcase{
		{
			obstacleGrid: [][]int{{1, 0}},
			expected:     0,
		},
		{
			obstacleGrid: [][]int{{1}, {0}},
			expected:     0,
		},
		{
			obstacleGrid: [][]int{{0, 1, 0}},
			expected:     0,
		},
		{
			obstacleGrid: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			expected:     2,
		},
		{
			obstacleGrid: [][]int{{0, 1}, {0, 0}},
			expected:     1,
		},
		{
			obstacleGrid: [][]int{{1}},
			expected:     0,
		},
	}

	for i, tc := range testcaes {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, uniquePathsWithObstacles(tc.obstacleGrid))
		})

	}
}
