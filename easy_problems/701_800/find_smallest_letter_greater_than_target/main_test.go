package findsmallestlettergreaterthantarget

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	letters  []byte
	target   byte
	expected byte
}

func TestNextGreatestLetter(t *testing.T) {
	testcases := []testcase{
		{
			letters:  []byte{'c', 'f', 'j'},
			target:   'd',
			expected: 'f',
		},
		{
			letters:  []byte{'c', 'f', 'j'},
			target:   'j',
			expected: 'c',
		},
		{
			letters:  []byte{'c', 'f', 'j'},
			target:   'a',
			expected: 'c',
		},
		{
			letters:  []byte{'c', 'f', 'j'},
			target:   'c',
			expected: 'f',
		},
		{
			letters:  []byte{'c', 'f', 'j'},
			target:   'z',
			expected: 'c',
		},

		{
			letters:  []byte{'a', 'a', 'b', 'b', 'c', 'c', 'e', 'e'},
			target:   'd',
			expected: 'e',
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := nextGreatestLetter(tc.letters, tc.target)
			assert.Equal(t, string(tc.expected), string(got))
		})
	}
}
