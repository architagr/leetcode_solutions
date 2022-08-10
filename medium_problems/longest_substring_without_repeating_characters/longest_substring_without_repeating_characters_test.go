package longest_substring_without_repeating_characters

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    string
	expected int
}

func TestLengthOfLongestSubstring(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input:    " ",
		expected: 1,
	})
	testCases = append(testCases, TestCase{
		input:    "abcabcbb",
		expected: 3,
	})
	testCases = append(testCases, TestCase{
		input:    "bbbbb",
		expected: 1,
	})
	testCases = append(testCases, TestCase{
		input:    "pwwkew",
		expected: 3,
	})
	testCases = append(testCases, TestCase{
		input:    "dvdf",
		expected: 3,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			got := LengthOfLongestSubstring(testCase.input)
			if got != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
