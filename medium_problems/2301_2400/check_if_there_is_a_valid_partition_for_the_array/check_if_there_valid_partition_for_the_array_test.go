package check_if_there_valid_partition_for_the_array

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    []int
	expected bool
}

func TestValidPartition(t *testing.T) {
	testCases := make([]TestCase, 0)
	// testCases = append(testCases, TestCase{
	// 	input:    []int{4, 4, 4, 5, 6},
	// 	expected: true,
	// })
	// testCases = append(testCases, TestCase{
	// 	input:    []int{1, 1, 1, 2},
	// 	expected: false,
	// })

	testCases = append(testCases, TestCase{
		input:    []int{4, 4, 4, 5, 6, 7, 7},
		expected: true,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := ValidPartition(testCase.input)
			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
