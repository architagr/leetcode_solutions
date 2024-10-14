package findthekbeautyofanumber

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputNum, inputk int
	expected         int
}

func TestDivisorSubstrings(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputNum: 240,
		inputk:   2,
		expected: 2,
	})
	testCases = append(testCases, TestCase{
		inputNum: 430043,
		inputk:   2,
		expected: 2,
	})
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := divisorSubstrings(testCase.inputNum, testCase.inputk)
			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
