package minimumdifficultyofajobschedule

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	jobDifficulty []int
	d             int
	expected      int
}

func TestMinDifficulty(t *testing.T) {
	testcases := []testcase{
		{
			jobDifficulty: []int{6, 5, 4, 3, 2, 1},
			d:             2,
			expected:      7,
		},
		{
			jobDifficulty: []int{9, 9, 9},
			d:             4,
			expected:      -1,
		},
		{
			jobDifficulty: []int{1, 1, 1},
			d:             3,
			expected:      3,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, minDifficulty(tc.jobDifficulty, tc.d))
		})
	}
}
