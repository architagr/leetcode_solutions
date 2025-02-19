package mergestringsalternatively

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	word1, word2 string
	expected     string
}

func TestMergeAlternately(t *testing.T) {
	testcases := []testcase{
		{
			word1:    "ab",
			word2:    "pqrs",
			expected: "apbqrs",
		},
		{
			word1:    "abc",
			word2:    "pqr",
			expected: "apbqcr",
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, mergeAlternately(tc.word1, tc.word2))
		})
	}
}
