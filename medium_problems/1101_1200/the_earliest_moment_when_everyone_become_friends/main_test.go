package theearliestmomentwheneveryonebecomefriends

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	n        int
	logs     [][]int
	expected int
}

func TestCountComponents(t *testing.T) {
	testcases := []testcase{
		{
			n: 6,
			logs: [][]int{
				{20190101, 0, 1},
				{20190104, 3, 4},
				{20190107, 2, 3},
				{20190211, 1, 5},
				{20190224, 2, 4},
				{20190301, 0, 3},
				{20190312, 1, 2},
				{20190322, 4, 5},
			},
			expected: 20190301,
		},
		{
			n: 4,
			logs: [][]int{
				{0, 2, 0}, {1, 0, 1}, {3, 0, 3}, {4, 1, 2}, {7, 3, 1},
			},
			expected: 3,
		},
		{
			n: 4,
			logs: [][]int{
				{7, 3, 1}, {2, 3, 0}, {3, 2, 1}, {6, 0, 1}, {0, 2, 0}, {4, 3, 2},
			},
			expected: 3,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, earliestAcq(tc.logs, tc.n))
		})
	}
}
