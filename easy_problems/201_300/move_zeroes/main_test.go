package movezeroes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	input    []int
	expected []int
}

func TestMoveZeroes(t *testing.T) {
	testcases := []testcase{
		{
			input:    []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			input:    []int{0},
			expected: []int{0},
		},
		{
			input:    []int{1},
			expected: []int{1},
		},
		{
			input:    []int{1, 0, 2},
			expected: []int{1, 2, 0},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			moveZeroes(tc.input)
			assert.Equal(t, tc.expected, tc.input)
		})

	}
}
