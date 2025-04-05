package uniquepaths

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	m, n     int
	expected int
}

func TestUniquePaths(t *testing.T) {
	testcaes := []testcase{
		{
			m:        3,
			n:        7,
			expected: 28,
		},
		{
			m:        1,
			n:        2,
			expected: 1,
		},
		{
			m:        3,
			n:        2,
			expected: 3,
		},
		{
			m:        23,
			n:        12,
			expected: 193536720,
		},
	}

	for i, tc := range testcaes {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, uniquePaths(tc.m, tc.n))
		})

	}
}
