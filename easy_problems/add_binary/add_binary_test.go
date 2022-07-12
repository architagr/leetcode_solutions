package add_binary

import (
	"fmt"
	"testing"
)

type testCase struct {
	inputA, inputB, expected string
}

func TestAddBinary(t *testing.T) {
	testCases := make([]testCase, 0)
	testCases = append(testCases, testCase{
		inputA:   "11",
		inputB:   "1",
		expected: "100",
	})
	testCases = append(testCases, testCase{
		inputA:   "1010",
		inputB:   "1011",
		expected: "10101",
	})
	for i, tCase := range testCases {
		t.Run(fmt.Sprintf("testing case %d", (i+1)), func(tb *testing.T) {
			got := AddBinary(tCase.inputA, tCase.inputB)
			if got != tCase.expected {
				tb.Errorf("Tested %d and expected %s but got %s", (i + 1), tCase.expected, got)
			}
		})
	}

}
