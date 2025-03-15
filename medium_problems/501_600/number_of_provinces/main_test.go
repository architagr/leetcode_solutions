package numberofprovinces

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	isConnected [][]int
	expected    int
}

func TestFindCircleNum(t *testing.T) {
	testcases := []testcase{
		{
			isConnected: [][]int{
				{1, 1, 0}, {1, 1, 0}, {0, 0, 1},
			},
			expected: 2,
		},
		{
			isConnected: [][]int{
				{1, 0, 0}, {0, 1, 0}, {0, 0, 1},
			},
			expected: 3,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, findCircleNum(tc.isConnected))
		})
	}
}
