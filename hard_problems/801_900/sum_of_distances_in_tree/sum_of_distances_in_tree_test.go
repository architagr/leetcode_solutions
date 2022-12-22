package sum_of_distances_in_tree

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	target   int
	stations [][]int
	expected []int
}

var testCases []TestCase = []TestCase{
	{
		target:   6,
		stations: [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}},
		expected: []int{8, 12, 6, 10, 10, 10},
	},
}

func TestMinRefuelStops(t *testing.T) {
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := sumOfDistancesInTree(testCase.target, testCase.stations)
			if !reflect.DeepEqual(got, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
