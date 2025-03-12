package validwordsquare

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	words    []string
	expected bool
}

func TestValidWordSquare(t *testing.T) {
	testcases := []testcase{
		{
			words:    []string{"abcd", "bnrt", "crmy", "dtye"},
			expected: true,
		},
		{
			words:    []string{"abcd", "bnrt", "crm", "dt"},
			expected: true,
		},
		{
			words:    []string{"ball", "area", "read", "lady"},
			expected: false,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, validWordSquare(tc.words))
		})
	}
}
