package uniqueemailaddresses

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	emails   []string
	expected int
}

func TestNumUniqueEmails(t *testing.T) {
	testcases := []testcase{
		{
			emails:   []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"},
			expected: 2,
		},
		{
			emails:   []string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"},
			expected: 3,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, numUniqueEmails(tc.emails))
		})
	}
}
