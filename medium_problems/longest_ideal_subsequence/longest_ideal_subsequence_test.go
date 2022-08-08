package longest_ideal_subsequence

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    string
	k        int
	expected int
}

func TestValidPartition(t *testing.T) {
	testCases := make([]TestCase, 0)

	// testCases = append(testCases, TestCase{
	// 	input:    "acfgbd",
	// 	k:        2,
	// 	expected: 4,
	// })
	// testCases = append(testCases, TestCase{
	// 	input:    "abcd",
	// 	k:        3,
	// 	expected: 4,
	// })
	// testCases = append(testCases, TestCase{
	// 	input:    "pvjcci",
	// 	k:        4,
	// 	expected: 2,
	// })
	testCases = append(testCases, TestCase{
		input:    "eduktdb",
		k:        15,
		expected: 5,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := LongestIdealString(testCase.input, testCase.k)
			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
