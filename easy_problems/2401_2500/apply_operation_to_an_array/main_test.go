package applyoperationtoanarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	input    []int
	expected []int
}

func TestApplyOperation(t *testing.T) {
	testcases := []testcase{
		{
			input:    []int{1, 2, 2, 1, 1, 0},
			expected: []int{1, 4, 2, 0, 0, 0},
		},
		{
			input:    []int{0, 1},
			expected: []int{1, 0},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, applyOperations(tc.input))
		})
	}
}
