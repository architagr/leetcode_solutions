package coursescheduleii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	numCourses    int
	prerequisites [][]int
	expected      []int
}

func TestFindOrder(t *testing.T) {
	testcases := []testcase{
		{
			numCourses: 2,
			prerequisites: [][]int{
				{1, 0},
			},
			expected: []int{0, 1},
		},
		{
			numCourses: 4,
			prerequisites: [][]int{
				{1, 0}, {2, 0}, {3, 1}, {3, 2},
			},
			expected: []int{0, 1, 2, 3},
		},
		{
			numCourses:    1,
			prerequisites: [][]int{},
			expected:      []int{0},
		},
		{
			numCourses: 3,
			prerequisites: [][]int{
				{1, 0}, {1, 2}, {0, 1},
			},
			expected: []int{},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, findOrder(tc.numCourses, tc.prerequisites))
		})
	}
}
