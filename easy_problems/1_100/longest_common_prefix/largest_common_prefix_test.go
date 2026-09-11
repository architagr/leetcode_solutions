package largestCommonPrefix

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {

	var testCases = []struct {
		input    []string
		expected string
	}{
		{input: []string{"flower", "flow", "flight"}, expected: "fl"},
		{input: []string{"dog", "racecar", "car"}, expected: ""},
		{input: []string{"racecar", "dog", "car"}, expected: ""},
		{input: []string{"racecar"}, expected: "racecar"},
		{input: []string{"ab", "a"}, expected: "a"},
		{input: []string{"aa", "ab"}, expected: "a"},
		{input: []string{"b", "a"}, expected: ""},
		{input: []string{"abcd", "abc"}, expected: "abc"},
		{input: []string{"ab", "abc", "abcd"}, expected: "ab"},
		{input: []string{"abc", "abc", "abc"}, expected: "abc"},
		{input: []string{"aac", "acab", "aa", "abba", "aa"}, expected: "a"},
	}

	for _, testcase := range testCases {
		t.Run(fmt.Sprintf("%+v as input", testcase.input), func(test *testing.T) {
			if got := longestCommonPrefix(testcase.input); got != testcase.expected {
				test.Fatalf("for input %+v expected is %s got %s", testcase.input, testcase.expected, got)
			}
		})
	}
}
