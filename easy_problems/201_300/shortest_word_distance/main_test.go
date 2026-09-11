package shortworddistance

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	wordsDict    []string
	word1, word2 string
	expected     int
}

func TestShortestDisctance(t *testing.T) {
	testcases := []testcase{
		{
			wordsDict: []string{"practice", "makes", "perfect", "coding", "makes"},
			word1:     "coding",
			word2:     "practice",
			expected:  3,
		},
		{
			wordsDict: []string{"practice", "makes", "perfect", "coding", "makes"},
			word1:     "makes",
			word2:     "coding",
			expected:  1,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			assert.Equal(t, tc.expected, shortestDistance(tc.wordsDict, tc.word1, tc.word2))
		})
	}
}
