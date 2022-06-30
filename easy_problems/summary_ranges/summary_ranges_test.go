package summary_ranges

import (
	"fmt"
	"reflect"
	"testing"
)

type testCaseData struct {
	nums           []int
	expectedLength []string
}

var testCases []testCaseData = []testCaseData{
	{nums: []int{1, 2, 3, 4, 6, 7, 8, 10, 12, 14, 15, 16}, expectedLength: []string{"1->4", "6->8", "10", "12", "14->16"}},
	{nums: []int{}, expectedLength: []string{}},
	{nums: []int{0}, expectedLength: []string{"0"}},
	{nums: []int{0, 1, 2, 4, 5, 7}, expectedLength: []string{"0->2", "4->5", "7"}},
	{nums: []int{0, 2, 3, 4, 6, 8, 9}, expectedLength: []string{"0", "2->4", "6", "8->9"}},
}

func TestSummaryRanges(t *testing.T) {

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test case no %d", (i+1)), func(t *testing.T) {
			got := SummaryRanges(testCase.nums)
			if !reflect.DeepEqual(got, testCase.expectedLength) {
				t.Errorf("error expected length %+v and got %+v\n", testCase.expectedLength, got)
			}
		})
	}
}
