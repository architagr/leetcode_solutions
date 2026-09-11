package ugly_number

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    int
	expected bool
}

var testCases []TestCase = []TestCase{
	{
		input:    6,
		expected: true,
	},
	{
		input:    1,
		expected: true,
	},
	{
		input:    -2147483648,
		expected: false,
	},
}

func TestIsUgly(t *testing.T) {

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := IsUgly(testCase.input)
			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
