package max_consecutive_ones

import (
	"fmt"
	"testing"
)

type testCaseData struct {
	nums           []int
	expectedLength int
}

var testCases []testCaseData = []testCaseData{
	{nums: []int{1, 1, 0}, expectedLength: 2},
	{nums: []int{0, 0, 1, 1, 1, 0, 1, 1, 0}, expectedLength: 3},
	{nums: []int{0, 0, 1, 1, 1, 0, 1, 1, 1, 1}, expectedLength: 4},
	{nums: []int{1}, expectedLength: 1},
	{nums: []int{0}, expectedLength: 0},
}

func TestFindMaxConsecutiveOnes(t *testing.T) {

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test case no %d", (i+1)), func(t *testing.T) {
			got := FindMaxConsecutiveOnes(testCase.nums)
			if got != testCase.expectedLength {
				t.Errorf("error expected length %d and got %d\n", testCase.expectedLength, got)
			}
		})
	}
}
