package equal_Row_and_column_pairs

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    [][]int
	expected int
}

func TestEqualPairs(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    [][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}},
		expected: 1,
	})
	testCases = append(testCases, TestCase{
		input:    [][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}},
		expected: 3,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := EqualPairs(testCase.input)

			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
