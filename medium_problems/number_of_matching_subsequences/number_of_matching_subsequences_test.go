package number_of_matching_subsequences

import (
	"fmt"
	"testing"
)

type TestCase struct {
	s      string
	words  []string
	output int
}

func TestNumMatchingSubseq(t *testing.T) {
	testCases := []TestCase{
		{
			s:      "abcde",
			words:  []string{"a", "bb", "acd", "ace"},
			output: 3,
		},
		{
			s:      "dsahjpjauf",
			words:  []string{"ahjpjau", "ja", "ahbwzgqnuk", "tnmlanowax"},
			output: 2,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("run itration number %d", (i+1)), func(test *testing.T) {
			output := NumMatchingSubseq(testCase.s, testCase.words)
			if output != testCase.output {
				test.Errorf("expected %+v, got %+v", testCase.output, output)
			}
		})

	}
}
