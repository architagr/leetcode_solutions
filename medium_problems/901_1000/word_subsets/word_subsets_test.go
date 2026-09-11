package word_subsets

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputA, inputB, expected []string
}

var testCases []TestCase = []TestCase{
	{
		inputA:   []string{"amazon", "apple", "facebook", "google", "leetcode"},
		inputB:   []string{"e", "o"},
		expected: []string{"facebook", "google", "leetcode"},
	},
}

func TestUniversalWordSubset(t *testing.T) {

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := UniversalWordSubset(testCase.inputA, testCase.inputB)
			if !reflect.DeepEqual(got, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
