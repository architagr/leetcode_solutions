package integer_to_roman

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    int
	expected string
}

func TestIntToRoman(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    3,
		expected: "III",
	})

	testCases = append(testCases, TestCase{
		input:    58,
		expected: "LVIII",
	})

	testCases = append(testCases, TestCase{
		input:    19,
		expected: "XIX",
	})

	testCases = append(testCases, TestCase{
		input:    1994,
		expected: "MCMXCIV",
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := IntToRoman(testCase.input)

			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
