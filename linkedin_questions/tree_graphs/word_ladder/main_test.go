package wordladder

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcase struct {
	beginWord string
	endWord   string
	wordList  []string
	expected  int
}

func TestLadderLength(t *testing.T) {
	testcases := []testcase{
		{
			beginWord: "hot",
			endWord:   "dog",
			wordList:  []string{"hot", "dog", "dot"},
			expected:  3,
		},
		{
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			expected:  5,
		},
		{
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			expected:  0,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := ladderLength(tc.beginWord, tc.endWord, tc.wordList)
			assert.Equal(t, tc.expected, got)
		})
	}
}
