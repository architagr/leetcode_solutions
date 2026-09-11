package prime_number_of_set_bits_in_binary_representation

import (
	"fmt"
	tpkg "testing"
)

type TestCase struct {
	left, right int
	expected    int
}

var testCases []TestCase = []TestCase{
	{
		left:     6,
		right:    10,
		expected: 4,
	},
}

func TestCountPrimeSetBits(t *tpkg.T) {

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *tpkg.T) {
			got := CountPrimeSetBits(testCase.left, testCase.right)
			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
