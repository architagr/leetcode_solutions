package minimumtimetotypewordusingspecialtypewriter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	word     string
	expected int
}

func TestMinTimeToType(t *testing.T) {
	testcases := []testcase{
		{
			word:     "abc",
			expected: 5,
		},
		{
			word:     "bza",
			expected: 7,
		},
		{
			word:     "zjpc",
			expected: 34,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("%d-%s", i, tc.word), func(t *testing.T) {
			assert.Equal(t, tc.expected, minTimeToType(tc.word))
		})
	}
}
