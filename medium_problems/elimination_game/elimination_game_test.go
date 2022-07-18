package elimination_game

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input, expected int
}

func TestLastRemaning(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    9,
		expected: 6,
	})
	testCases = append(testCases, TestCase{
		input:    6,
		expected: 4,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := LastRemaning(testCase.input)

			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
