package longestcommonsubsequesnce

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	text1, text2 string
	expected     int
}

func TestLongestCommonSubsequence(t *testing.T) {
	testcases := []testcase{
		{
			text1:    "abcde",
			text2:    "ace",
			expected: 3,
		},
		{
			text1:    "abc",
			text2:    "abc",
			expected: 3,
		},
		{
			text1:    "abc",
			text2:    "def",
			expected: 0,
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, longestCommonSubsequence(tc.text1, tc.text2))
		})
	}
}
