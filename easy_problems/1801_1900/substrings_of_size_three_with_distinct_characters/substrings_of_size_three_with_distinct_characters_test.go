package substringsofsizethreewithdistinctcharacters

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    string
	expected int
}

func TestCountGoodSubstrings(t *testing.T) {

	testcases := []TestCase{
		{
			input:    "xyzzaz",
			expected: 1,
		},
		{
			input:    "aababcabc",
			expected: 4,
		},
	}

	for i, caseObj := range testcases {
		t.Run(fmt.Sprintf("running %d", i), func(tb *testing.T) {
			got := countGoodSubstrings(caseObj.input)
			if got != caseObj.expected {
				tb.Fatalf("tested %d expected %+v but got %+v", (i + 1), caseObj.expected, got)
			}
		})
	}

}
