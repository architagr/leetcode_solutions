package wordpattern

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	pattern, s string
	expected   bool
}

func TestWordPattern(t *testing.T) {
	testcases := []testcase{
		{
			pattern:  "abba",
			s:        "dog cat cat dog",
			expected: true,
		},
		{
			pattern:  "abba",
			s:        "dog cat cat fish",
			expected: false,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, wordPattern(tc.pattern, tc.s))
		})
	}
}
