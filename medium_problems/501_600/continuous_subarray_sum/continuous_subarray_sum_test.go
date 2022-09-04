package continuous_subarray_sum

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    []int
	k        int
	expected bool
}

func TestCheckSubarraySum(t *testing.T) {
	testCases := make([]TestCase, 0)
	// testCases = append(testCases, TestCase{
	// 	input:    []int{23, 2, 4, 6, 7},
	// 	k:        6,
	// 	expected: true,
	// })

	// testCases = append(testCases, TestCase{
	// 	input:    []int{23, 2, 6, 4, 7},
	// 	k:        13,
	// 	expected: false,
	// })

	// testCases = append(testCases, TestCase{
	// 	input:    []int{23, 2, 4, 6, 6},
	// 	k:        7,
	// 	expected: true,
	// })

	// testCases = append(testCases, TestCase{
	// 	input:    []int{5, 0, 0, 0},
	// 	k:        7,
	// 	expected: true,
	// })

	// testCases = append(testCases, TestCase{
	// 	input:    []int{0},
	// 	k:        7,
	// 	expected: false,
	// })

	// testCases = append(testCases, TestCase{
	// 	input:    []int{0, 1, 0, 3, 0, 4, 0, 4, 0},
	// 	k:        5,
	// 	expected: false,
	// })

	testCases = append(testCases, TestCase{
		input:    []int{1, 3, 0, 6},
		k:        6,
		expected: true,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := CheckSubarraySum(testCase.input, testCase.k)

			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
