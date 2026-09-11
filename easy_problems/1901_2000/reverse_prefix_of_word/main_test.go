package reverseprefixofword

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	word     string
	ch       byte
	expected string
}

func TestReversePrefix(t *testing.T) {
	testcases := []testcase{
		{
			word:     "abcdefd",
			ch:       'd',
			expected: "dcbaefd",
		},
		{
			word:     "xyxzxe",
			ch:       'z',
			expected: "zxyxxe",
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, reversePrefix(tc.word, tc.ch))
		})
	}
}
