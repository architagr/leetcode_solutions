package find_and_replace_pattern

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input    []string
	pattern  string
	expected []string
}

func TestFindAndReplacePattern(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"},
		pattern:  "abb",
		expected: []string{"mee", "aqq"},
	})
	testCases = append(testCases, TestCase{
		input:    []string{"a", "b", "c"},
		pattern:  "a",
		expected: []string{"a", "b", "c"},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := FindAndReplacePattern(testCase.input, testCase.pattern)

			if !reflect.DeepEqual(got, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
