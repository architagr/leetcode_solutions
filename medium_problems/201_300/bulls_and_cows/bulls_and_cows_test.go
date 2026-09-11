package bulls_and_cows

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputS, inputG, expected string
}

func TestGetHint(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputS:   "1807",
		inputG:   "7810",
		expected: "1A3B",
	})

	testCases = append(testCases, TestCase{
		inputS:   "1123",
		inputG:   "0111",
		expected: "1A1B",
	})

	testCases = append(testCases, TestCase{
		inputS:   "1122",
		inputG:   "1222",
		expected: "3A0B",
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := GetHint(testCase.inputS, testCase.inputG)

			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
