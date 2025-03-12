package licensekeyformatting

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	s        string
	k        int
	expected string
}

func TestLicenseKeyFormatting(t *testing.T) {
	testcases := []testcase{
		{
			s:        "5F3Z-2e-9-w",
			k:        4,
			expected: "5F3Z-2E9W",
		},
		{
			s:        "2-5g-3-J",
			k:        2,
			expected: "2-5G-3J",
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, licenseKeyFormatting(tc.s, tc.k))
		})
	}
}
