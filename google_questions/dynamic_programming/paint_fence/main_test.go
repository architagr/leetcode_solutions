package paintfence

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	n, k int

	expected int
}

func TestNumWays(t *testing.T) {
	testcases := []testcase{
		{
			n:        3,
			k:        2,
			expected: 6,
		},
		{
			n:        1,
			k:        1,
			expected: 1,
		},
		{
			n:        1,
			k:        4,
			expected: 4,
		},
		{
			n:        2,
			k:        4,
			expected: 16,
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, numWays(tc.n, tc.k))
		})
	}
}
