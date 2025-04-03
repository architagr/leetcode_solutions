package decodeways

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	s        string
	expected int
}

func TestNumDecodings(t *testing.T) {
	testcases := []testcase{
		{
			s:        "12",
			expected: 2,
		},
		{
			s:        "226",
			expected: 3,
		},
		{
			s:        "06",
			expected: 0,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, numDecodings(tc.s))
		})
	}
}
