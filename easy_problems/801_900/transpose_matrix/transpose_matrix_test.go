package transpose_matrix

import (
	"fmt"
	"reflect"
	"testing"
)

type testCaseData struct {
	nums, expected [][]int
}

var testCases []testCaseData = []testCaseData{
	{nums: [][]int{{1, 2, 3}, {4, 5, 6}}, expected: [][]int{{1, 4}, {2, 5}, {3, 6}}},
	{nums: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, expected: [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}},
}

func TestTranspose(t *testing.T) {

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test case no %d", (i+1)), func(t *testing.T) {
			got := Transpose(testCase.nums)
			if !reflect.DeepEqual(got, testCase.expected) {
				t.Errorf("error expected length %+v and got %+v\n", testCase.expected, got)
			}
		})
	}
}
