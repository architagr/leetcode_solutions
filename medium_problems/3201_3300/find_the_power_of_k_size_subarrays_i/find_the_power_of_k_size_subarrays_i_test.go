package findthepowerofksizesubarraysi

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputNums []int
	first     int
	expected  []int
}

func TestResultsArray(t *testing.T) {
	testCases := []TestCase{
		{
			inputNums: []int{1, 2, 3, 4, 3, 2, 5},
			first:     3,
			expected:  []int{3, 4, -1, -1, -1},
		},
		{
			inputNums: []int{2, 2, 2, 2, 2},
			first:     4,
			expected:  []int{-1, -1},
		},
		{
			inputNums: []int{2},
			first:     1,
			expected:  []int{2},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := resultsArray(testCase.inputNums, testCase.first)

			if !reflect.DeepEqual(got, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
