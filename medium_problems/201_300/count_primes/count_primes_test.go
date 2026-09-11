package count_primes

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    int
	expected int
}

func TestCountBadPairs(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    10,
		expected: 4,
	})
	testCases = append(testCases, TestCase{
		input:    2,
		expected: 0,
	})
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := CountPrimes(testCase.input)
			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
