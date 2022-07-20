package count_vowel_substrings

import (
	"fmt"
	"testing"
)

type TestCase struct {
	word   string
	output int
}

func TestCountVowelSubstrings(t *testing.T) {
	testCases := []TestCase{
		{
			word:   "aeiouu",
			output: 2,
		},
		{
			word:   "unicornarihan",
			output: 0,
		},
		{
			word:   "cuaieuouac",
			output: 7,
		},
		{
			word:   "poazaeuioauoiioaouuouaui",
			output: 31,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("run itration number %d", (i+1)), func(test *testing.T) {
			output := CountVowelSubstrings(testCase.word)
			if output != testCase.output {
				test.Errorf("expected %+v, got %+v", testCase.output, output)
			}
		})

	}
}
