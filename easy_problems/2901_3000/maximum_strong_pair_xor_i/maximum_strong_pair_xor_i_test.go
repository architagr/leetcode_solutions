package maximumstrongpairxori

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputCode []int
	expected  int
}

func TestMaximumStrongPairXor(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputCode: []int{1, 2, 3, 4, 5},
		expected:  7,
	})

	testCases = append(testCases, TestCase{
		inputCode: []int{10, 100},
		expected:  0,
	})

	testCases = append(testCases, TestCase{
		inputCode: []int{5, 6, 25, 30},
		expected:  7,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := maximumStrongPairXor(testCase.inputCode)
			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
