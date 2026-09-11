package plus_one

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {

	var testCases = []struct {
		input    []int
		expected []int
	}{
		{input: []int{1, 2, 3}, expected: []int{1, 2, 4}},
		{input: []int{1, 2, 9}, expected: []int{1, 3, 0}},
		{input: []int{9, 9, 9}, expected: []int{1, 0, 0, 0}},
		{input: []int{9}, expected: []int{1, 0}},
	}

	for i, testcase := range testCases {
		t.Run(fmt.Sprintf("%d as input", (i+1)), func(test *testing.T) {
			if got := plusOne(testcase.input); !reflect.DeepEqual(got, testcase.expected) {
				test.Fatalf("for input %d expected is %+v got %+v", (i + 1), testcase.expected, got)
			}
		})
	}
}
