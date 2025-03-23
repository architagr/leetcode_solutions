package cheapestflightswithinkstops

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	n        int
	flights  [][]int
	src      int
	dst      int
	k        int
	expected int
}

func TestFindCheapestPrice(t *testing.T) {
	testcases := []testcase{
		{
			n: 4,
			flights: [][]int{
				{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200},
			},
			src:      0,
			dst:      3,
			k:        1,
			expected: 700,
		},
		{
			n: 3,
			flights: [][]int{
				{0, 1, 100}, {1, 2, 100}, {0, 2, 500},
			},
			src:      0,
			dst:      2,
			k:        1,
			expected: 200,
		},
		{
			n: 3,
			flights: [][]int{
				{0, 1, 100}, {1, 2, 100}, {0, 2, 500},
			},
			src:      0,
			dst:      2,
			k:        0,
			expected: 500,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, findCheapestPrice(tc.n, tc.flights, tc.src, tc.dst, tc.k))
		})
	}
}
