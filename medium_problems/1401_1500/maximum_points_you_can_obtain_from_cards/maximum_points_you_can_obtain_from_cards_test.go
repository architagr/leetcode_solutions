package maximumpointsyoucanobtainfromcards

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputNums []int
	first     int
	expected  int
}

func TestMaxScore(t *testing.T) {
	testCases := []TestCase{
		{
			inputNums: []int{1, 2, 3, 4, 5, 6, 1},
			first:     3,
			expected:  12,
		},
		{
			inputNums: []int{2, 2, 2},
			first:     2,
			expected:  4,
		},
		{
			inputNums: []int{9, 7, 7, 9, 7, 7, 9},
			first:     7,
			expected:  55,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := maxScore(testCase.inputNums, testCase.first)

			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
