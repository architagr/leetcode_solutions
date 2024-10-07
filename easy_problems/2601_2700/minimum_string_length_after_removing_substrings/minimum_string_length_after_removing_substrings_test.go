package minimumstringlengthafterremovingsubstrings

import (
	"fmt"
	"testing"
)

type TestCase struct {
	s        string
	expected int
}

func TestMinLength(t *testing.T) {
	testCases := []TestCase{
		{
			s:        "ABFCACDB",
			expected: 2,
		},
		{
			s:        "ACBBD",
			expected: 5,
		},
		{
			s:        "A",
			expected: 1,
		},
		{
			s:        "CCCCDDDD",
			expected: 0,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t1 *testing.T) {
			if got := minLength(tc.s); got != tc.expected {
				t1.Errorf("expected %d but got %d", tc.expected, got)
			}
		})
	}
}
