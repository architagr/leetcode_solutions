package permutationinstring

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputA, inputB string
	expected       bool
}

func TestFindMaxLength(t *testing.T) {
	testCases := []TestCase{
		{
			inputA:   "ab",
			inputB:   "eidbaooo",
			expected: true,
		},
		{
			inputA:   "ab",
			inputB:   "eidboaoo",
			expected: false,
		},
		{
			inputA:   "abc",
			inputB:   "eidbcaoo",
			expected: true,
		},
		{
			inputA:   "abbc",
			inputB:   "eibdacbcbboo",
			expected: false,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t1 *testing.T) {
			if got := checkInclusion(tc.inputA, tc.inputB); got != tc.expected {
				t1.Errorf("expected %t but got %t", tc.expected, got)
			}
		})
	}
}
