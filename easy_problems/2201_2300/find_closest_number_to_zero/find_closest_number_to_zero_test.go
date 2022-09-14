package find_closest_number_to_zero

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    []int
	expected int
}

func TestFindClosestNumber(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    []int{1, 4, 2, 5, 3},
		expected: 1,
	})
	testCases = append(testCases, TestCase{
		input:    []int{1, 4, -1, 5, 3},
		expected: 1,
	})

	testCases = append(testCases, TestCase{
		input:    []int{2, 1, 4, -1, 5, 3},
		expected: 1,
	})

	testCases = append(testCases, TestCase{
		input:    []int{2, -1, 4, 1, 5, 3},
		expected: 1,
	})

	testCases = append(testCases, TestCase{
		input:    []int{2, -1, 4, -1, 5, 3},
		expected: -1,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := FindClosestNumber(testCase.input)

			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
