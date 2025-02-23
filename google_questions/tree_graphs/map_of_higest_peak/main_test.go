package mapofhigestpeak

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	isWater  [][]int
	expected [][]int
}

func TestHigestPeak(t *testing.T) {
	testcases := []testcase{
		{
			isWater: [][]int{
				{0, 0, 1},
				{1, 0, 0},
				{0, 0, 0},
			},
			expected: [][]int{
				{1, 1, 0},
				{0, 1, 1},
				{1, 2, 2},
			},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, highestPeak(tc.isWater))
		})
	}
}
