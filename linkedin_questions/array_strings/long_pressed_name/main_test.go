package longpressedname

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	name, typed string
	expected    bool
}

func TestIsLongPressedName(t *testing.T) {
	testcases := []testcase{
		{
			name:     "alex",
			typed:    "aaleexa",
			expected: false,
		},
		{
			name:     "alex",
			typed:    "aaleex",
			expected: true,
		},
		{
			name:     "saeed",
			typed:    "ssaaedd",
			expected: false,
		},
		{
			name:     "alexd",
			typed:    "ale",
			expected: false,
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, isLongPressedName(tc.name, tc.typed))
		})
	}
}
