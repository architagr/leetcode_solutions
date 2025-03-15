package smalleststringwithswaps

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	s        string
	pairs    [][]int
	expected string
}

func TestSmallestStringWithSwaps(t *testing.T) {
	testcases := []testcase{
		{
			s: "dcab",
			pairs: [][]int{
				{0, 3}, {1, 2},
			},
			expected: "bacd",
		},
		{
			s: "dcab",
			pairs: [][]int{
				{0, 3}, {1, 2}, {0, 2},
			},
			expected: "abcd",
		},
		{
			s: "cba",
			pairs: [][]int{
				{0, 1}, {1, 2},
			},
			expected: "abc",
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, smallestStringWithSwaps(tc.s, tc.pairs))
		})
	}
}
