package findandreplaceinstring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	s        string
	indices  []int
	sources  []string
	targets  []string
	expected string
}

func TestFindReplaceString(t *testing.T) {
	testcases := []testcase{
		{
			s:        "abcde",
			indices:  []int{2, 2, 3},
			sources:  []string{"cde", "cdef", "dk"},
			targets:  []string{"fe", "f", "xyz"},
			expected: "abfe",
		},
		{
			s:        "abcd",
			indices:  []int{0, 2},
			sources:  []string{"a", "cd"},
			targets:  []string{"eee", "ffff"},
			expected: "eeebffff",
		},
		{
			s:        "abcd",
			indices:  []int{0, 2},
			sources:  []string{"ab", "ec"},
			targets:  []string{"eee", "ffff"},
			expected: "eeecd",
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, findReplaceString(tc.s, tc.indices, tc.sources, tc.targets))
		})
	}
}
