package countdayswithoutmeetings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	days     int
	meetings [][]int
	expected int
}

func TestCountDays(t *testing.T) {
	testcases := []testcase{
		{
			days: 1000000000,
			meetings: [][]int{
				{1, 1000000000},
			},
			expected: 0,
		},
		{
			days: 8,
			meetings: [][]int{
				{3, 4}, {4, 8}, {2, 5}, {3, 8},
			},
			expected: 1,
		},
		{
			days: 10,
			meetings: [][]int{
				{5, 7}, {1, 3}, {9, 10},
			},
			expected: 2,
		},
		{
			days: 5,
			meetings: [][]int{
				{2, 4}, {1, 3},
			},
			expected: 1,
		},
		{
			days: 6,
			meetings: [][]int{
				{1, 6},
			},
			expected: 0,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, countDays(tc.days, tc.meetings))
		})
	}
}
