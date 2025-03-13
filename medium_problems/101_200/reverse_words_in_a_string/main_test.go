package reversewordsinastring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	s        string
	expected string
}

func TestReverseWords(t *testing.T) {
	testcases := []testcase{
		{
			s:        "the sky is blue",
			expected: "blue is sky the",
		},
		{
			s:        "  hello world  ",
			expected: "world hello",
		},
		{
			s:        "a good   example",
			expected: "example good a",
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, reverseWords(tc.s))
		})
	}
}
