package validword

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	word     string
	expected bool
}

func TestIsValid(t *testing.T) {
	testcases := []testcase{
		{
			word:     "234Adas",
			expected: true,
		},
		{
			word:     "a3",
			expected: false,
		},
		{
			word:     "a3$e",
			expected: false,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, isValid(tc.word))
		})
	}
}
