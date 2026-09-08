package longestsubstringwithoutrepeatingcharacters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	s        string
	expected int
}

func TestLengthOfLongestSubstring(t *testing.T) {
	testcases := []testcase{
		{
			s:        "aabaab!bb",
			expected: 3,
		},
		{
			s:        "abcabcbb",
			expected: 3,
		},
		{
			s:        "bbbb",
			expected: 1,
		},
		{
			s:        "pwwkew",
			expected: 3,
		},
		{
			s:        "",
			expected: 0,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.s, func(t *testing.T) {
			assert.Equal(t, tc.expected, lengthOfLongestSubstring(tc.s))
		})
	}
}
