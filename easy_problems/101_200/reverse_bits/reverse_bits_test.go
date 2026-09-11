package reverse_bits

import (
	"fmt"
	"testing"
)

type testCaseData struct {
	nums, expected uint32
}

var testCases []testCaseData = []testCaseData{
	{nums: uint32(43261596), expected: uint32(964176192)},
}

func TestReverseBits(t *testing.T) {

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test case no %d", (i+1)), func(t *testing.T) {
			got := ReverseBits(testCase.nums)
			if got != testCase.expected {
				t.Errorf("error expected length %d and got %d\n", testCase.expected, got)
			}
		})
	}
}
