package numberofsubarraysofsizekandaveragegreaterthanorequaltothreshold

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputNum                []int
	inputk, inputThreshhold int
	expected                int
}

func TestNumOfSubarrays(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputNum:        []int{2, 2, 2, 2, 5, 5, 5, 8},
		inputk:          3,
		inputThreshhold: 4,
		expected:        3,
	})
	testCases = append(testCases, TestCase{
		inputNum:        []int{11, 13, 17, 23, 29, 31, 7, 5, 2, 3},
		inputk:          3,
		inputThreshhold: 5,
		expected:        6,
	})
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := numOfSubarrays(testCase.inputNum, testCase.inputk, testCase.inputThreshhold)
			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
