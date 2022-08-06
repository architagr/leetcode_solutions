package count_number_of_bad_pairs

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    []int
	expected int64
}

func TestCountBadPairs(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    []int{1, 2, 3, 4, 5},
		expected: 0,
	})

	testCases = append(testCases, TestCase{
		input:    []int{56, 30, 2, 73, 28, 81, 75, 75, 18, 63, 54, 10, 69, 85, 33, 89, 12, 78, 57, 60, 54, 65, 74, 63, 53, 77},
		expected: 322,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := CountBadPairs(testCase.input)
			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
