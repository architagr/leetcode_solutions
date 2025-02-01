package textjustification

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	input    []string
	maxWidth int
	expected []string
}

func TestFullJustify(t *testing.T) {
	testCases := []testCase{
		{
			input:    []string{"This", "is", "an", "example", "of", "text", "justification."},
			maxWidth: 16,
			expected: []string{"This    is    an", "example  of text", "justification.  "},
		},
		{
			input:    []string{"What", "must", "be", "acknowledgment", "shall", "be"},
			maxWidth: 16,
			expected: []string{"What   must   be", "acknowledgment  ", "shall be        "},
		},
		{
			input:    []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			maxWidth: 20,
			expected: []string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  ",
			},
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, fullJustify(tc.input, tc.maxWidth))
		})
	}
}

type testCaseSentence struct {
	inputWords      []string
	maxWidth        int
	isLeftJustified bool
	expected        string
}

func TestBuildJustifySentence(t *testing.T) {
	testcases := []testCaseSentence{
		{
			inputWords: []string{"This", "is", "an"},
			maxWidth:   16,
			expected:   "This    is    an",
		},
		{
			inputWords: []string{"example", "of", "text"},
			maxWidth:   16,
			expected:   "example  of text",
		},
		{
			inputWords:      []string{"justification."},
			maxWidth:        16,
			isLeftJustified: true,
			expected:        "justification.  ",
		},
		{
			inputWords: []string{"What", "must", "be"},
			maxWidth:   16,
			expected:   "What   must   be",
		},
		{
			inputWords: []string{"acknowledgment"},
			maxWidth:   16,
			expected:   "acknowledgment  ",
		},
		{
			inputWords:      []string{"shall", "be"},
			maxWidth:        16,
			isLeftJustified: true,
			expected:        "shall be        ",
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("%d-%s", i, tc.expected), func(t *testing.T) {
			assert.Equal(t, tc.expected, buildJustifySentence(tc.inputWords, tc.maxWidth, tc.isLeftJustified))
		})

	}
}
