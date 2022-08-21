package sumofalloddlengthsubarrays

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    []int
	expected int
}

func TestSumOddLengthSubarrays(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    []int{1, 4, 2, 5, 3},
		expected: 58,
	})
	testCases = append(testCases, TestCase{
		input:    []int{1, 2},
		expected: 3,
	})

	testCases = append(testCases, TestCase{
		input:    []int{10, 11, 12},
		expected: 66,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := SumOddLengthSubarrays(testCase.input)

			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
