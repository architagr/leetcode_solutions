package permutation_sequence

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputN, inputK int
	expected       string
}

var testCases []TestCase = []TestCase{
	{
		inputN:   3,
		inputK:   2,
		expected: "132",
	},
	{
		inputN:   3,
		inputK:   3,
		expected: "213",
	},
	{
		inputN:   4,
		inputK:   9,
		expected: "2314",
	},
	{
		inputN:   3,
		inputK:   1,
		expected: "123",
	},
	{
		inputN:   2,
		inputK:   2,
		expected: "21",
	},
	{
		inputN:   3,
		inputK:   6,
		expected: "321",
	},
}

func TestGetPermutation(t *testing.T) {

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := GetPermutation(testCase.inputN, testCase.inputK)
			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
