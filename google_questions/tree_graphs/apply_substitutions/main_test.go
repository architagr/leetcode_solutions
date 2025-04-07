package applysubstitutions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	replacements [][]string
	text         string
	expected     string
}

func TestApplySubstitutions(t *testing.T) {
	testcases := []testcase{
		{
			replacements: [][]string{
				{"A", "abc"}, {"B", "def"},
			},
			text:     "%A%_%B%",
			expected: "abc_def",
		},
		{
			replacements: [][]string{
				{"A", "bce"}, {"B", "ace"}, {"C", "abc%B%"},
			},
			text:     "%A%_%B%_%C%",
			expected: "bce_ace_abcace",
		},
		{
			replacements: [][]string{
				{"A", "bce"}, {"B", "ace"}, {"C", "abc%D%"}, {"D", "adc%B%"},
			},
			text:     "%A%_%B%_%C%_%D%",
			expected: "bce_ace_abcadcace_adcace",
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, applySubstitutions(tc.replacements, tc.text))
		})
	}
}
