package evaluatereversepolishnotation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	input    []string
	expected int
}

func TestEvalRPN(t *testing.T) {
	testcases := []testcase{
		{
			input:    []string{"2", "1", "+", "3", "*"},
			expected: 9,
		},
		{
			input:    []string{"4", "13", "5", "/", "+"},
			expected: 6,
		},
		{
			input:    []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			expected: 22,
		},
		{
			input:    []string{"4", "3", "-"},
			expected: 1,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, evalRPN(tc.input))
		})
	}
}
