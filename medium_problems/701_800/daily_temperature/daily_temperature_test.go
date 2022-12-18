package daily_temperature

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input    []int
	expected []int
}

var testCases []TestCase = []TestCase{
	// {[]int{1, 6, 10, 8, 7, 3, 2}, 18},
	// {[]int{1, 0, 3}, 5},
	// {[]int{1, 2, 3}, 6},
	// {[]int{1, 2, 2}, 4},
	// {[]int{1, 3, 2}, 4},
	{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
	{[]int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
	{[]int{30, 60, 90}, []int{1, 1, 0}},
}

func TestDailyTemperatures(t *testing.T) {

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			if got := dailyTemperatures(testCase.input); !reflect.DeepEqual(got, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
