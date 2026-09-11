package number_of_arithmetic_triplets

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    []int
	space    int
	expected int
}

func TestArithmeticTriplets(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    []int{0, 1, 4, 6, 7, 10},
		space:    3,
		expected: 2,
	})
	testCases = append(testCases, TestCase{
		input:    []int{4, 5, 6, 7, 8, 9},
		space:    2,
		expected: 2,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := ArithmeticTriplets(testCase.input, testCase.space)
			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
