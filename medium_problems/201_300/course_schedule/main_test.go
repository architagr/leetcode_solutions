package courseschedule

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	numCourses    int
	prerequisites [][]int
	expected      bool
}

func TestFindOrder(t *testing.T) {
	testcases := []testcase{
		{
			numCourses: 2,
			prerequisites: [][]int{
				{1, 0},
			},
			expected: true,
		},
		{
			numCourses: 4,
			prerequisites: [][]int{
				{1, 0}, {2, 0}, {3, 1}, {3, 2},
			},
			expected: true,
		},
		{
			numCourses:    1,
			prerequisites: [][]int{},
			expected:      true,
		},
		{
			numCourses: 3,
			prerequisites: [][]int{
				{1, 0}, {1, 2}, {0, 1},
			},
			expected: false,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, canFinish(tc.numCourses, tc.prerequisites))
		})
	}
}
