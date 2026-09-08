package lettercombinationsofaphonenumber

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	input    string
	expected []string
}

func TestLetterCombinations(t *testing.T) {
	testcases := []testcase{
		{
			input:    "23",
			expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "2",
			expected: []string{"a", "b", "c"},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := letterCombinations(tc.input)
			assert.Equal(t, tc.expected, got)
		})
	}
}
