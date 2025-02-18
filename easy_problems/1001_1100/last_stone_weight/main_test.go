package laststoneweight

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	stones   []int
	expected int
}

func TestLastStoneWeight(t *testing.T) {
	testcase := []testcase{
		{
			stones:   []int{2, 7, 4, 1, 8, 1},
			expected: 1,
		},
		{
			stones:   []int{1},
			expected: 1,
		},
		{
			stones:   []int{1, 1},
			expected: 0,
		},
	}

	for i, tc := range testcase {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, lastStoneWeight(tc.stones))
		})
	}
}
