package isomorphicstrings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	s, t     string
	expected bool
}

func TestIsIsomorphic(t *testing.T) {
	testcases := []testcase{
		{
			s:        "egg",
			t:        "add",
			expected: true,
		},
		{
			s:        "foo",
			t:        "bar",
			expected: false,
		},
		{
			s:        "bbbaaaba",
			t:        "aaabbbba",
			expected: false,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, isIsomorphic(tc.s, tc.t))
		})
	}
}
