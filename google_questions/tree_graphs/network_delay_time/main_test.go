package networkdelaytime

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	times          [][]int
	n, k, expected int
}

func TestNetworkDelayTime(t *testing.T) {
	testcases := []testcase{
		{
			times: [][]int{
				{2, 4, 10}, {5, 2, 38}, {3, 4, 33},
				{4, 2, 76}, {3, 2, 64}, {1, 5, 54},
				{1, 4, 98}, {2, 3, 61}, {2, 1, 0},
				{3, 5, 77}, {5, 1, 34}, {3, 1, 79},
				{5, 3, 2}, {1, 2, 59}, {4, 3, 46},
				{5, 4, 44}, {2, 5, 89}, {4, 5, 21},
				{1, 3, 86}, {4, 1, 95},
			},
			n:        5,
			k:        1,
			expected: 69,
		},
		{
			times: [][]int{
				{2, 1, 1}, {2, 3, 1}, {3, 4, 1},
			},
			n:        4,
			k:        2,
			expected: 2,
		},
		{
			times: [][]int{
				{1, 2, 1},
			},
			n:        2,
			k:        1,
			expected: 1,
		},
		{
			times: [][]int{
				{1, 2, 1},
			},
			n:        2,
			k:        2,
			expected: -1,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, networkDelayTime(tc.times, tc.n, tc.k))
		})
	}
}
