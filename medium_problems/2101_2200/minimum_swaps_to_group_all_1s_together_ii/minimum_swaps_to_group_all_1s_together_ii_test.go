package minimumswapstogroupall1stogetherii

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    []int
	expected int
}

func TestMinSwaps(t *testing.T) {
	testCases := []TestCase{
		{
			input:    []int{0, 1, 0, 1, 1, 0, 0},
			expected: 1,
		},
		{
			input:    []int{0, 1, 1, 1, 0, 0, 1, 1, 0},
			expected: 2,
		},
		{
			input:    []int{1, 1, 0, 0, 1},
			expected: 0,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := minSwaps(testCase.input)

			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
