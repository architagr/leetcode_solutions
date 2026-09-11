package plates_between_candles

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputA   string
	queries  [][]int
	expected []int
}

func TestPartitionList(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputA:   "**|**|***|",
		queries:  [][]int{{2, 5}, {5, 9}},
		expected: []int{2, 3},
	})
	testCases = append(testCases, TestCase{
		inputA:   "***|**|*****|**||**|*",
		queries:  [][]int{{1, 17}, {4, 5}, {14, 17}, {5, 11}, {15, 16}},
		expected: []int{9, 0, 0, 0, 0},
	})
	testCases = append(testCases, TestCase{
		inputA:   "***",
		queries:  [][]int{{2, 2}},
		expected: []int{0},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := PlatesBetweenCandles(testCase.inputA, testCase.queries)
			if !reflect.DeepEqual(got, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
