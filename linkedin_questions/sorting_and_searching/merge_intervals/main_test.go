package mergeintervals

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	nums   [][]int
	output [][]int
}

func TestMerge(t *testing.T) {
	testCases := []TestCase{
		{
			nums:   [][]int{{2, 3}, {5, 5}, {2, 2}, {3, 4}, {3, 4}},
			output: [][]int{{2, 4}, {5, 5}},
		},
		{
			nums:   [][]int{{1, 4}, {0, 2}, {3, 5}},
			output: [][]int{{0, 5}},
		},
		{
			nums:   [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}},
			output: [][]int{{1, 10}},
		},
		{
			nums:   [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			output: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			nums:   [][]int{{1, 4}, {4, 5}},
			output: [][]int{{1, 5}},
		},
		{
			nums:   [][]int{{1, 4}, {2, 3}},
			output: [][]int{{1, 4}},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("run itration number %d", (i+1)), func(test *testing.T) {
			output := merge(testCase.nums)
			if !reflect.DeepEqual(output, testCase.output) {
				test.Errorf("expected %+v, got %+v", testCase.output, output)
			}
		})

	}
}
