package sort_characters_by_frequency

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    string
	expected []string
}

func TestReverseList(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    "tree",
		expected: []string{"eert", "eetr"},
	})
	testCases = append(testCases, TestCase{
		input:    "Aabb",
		expected: []string{"bbAa", "bbaA"},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := frequencySort(testCase.input)

			if !verify(testCase.expected, got) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
func verify(expectedList []string, got string) bool {

	for _, val := range expectedList {
		if val == got {
			return true
		}
	}
	return false
}
