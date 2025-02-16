package wordladder

type stackNode struct {
	word string
	seq  int
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]bool, len(wordList))
	for _, word := range wordList {
		wordSet[word] = true
	}
	if _, ok := wordSet[endWord]; !ok {
		return 0
	}
	delete(wordSet, beginWord)
	stack := make([]stackNode, 0, len(wordSet))
	stack = append(stack, stackNode{
		word: beginWord,
		seq:  1,
	})
	return process(endWord, wordSet, &stack)
}

func process(endWord string, wordSet map[string]bool, stack *[]stackNode) int {
	pop := func() stackNode {
		c := (*stack)[0]
		(*stack) = (*stack)[1:]
		return c
	}
	push := func(str string, seq int) {
		delete(wordSet, str)
		(*stack) = append((*stack), stackNode{
			word: str,
			seq:  seq,
		})
	}

	for len(*stack) > 0 {
		beginWord := pop()
		newWord := make([]byte, len(beginWord.word))
		copy(newWord, []byte(beginWord.word))
		for i := 0; i < len(beginWord.word); i++ {
			for j := 0; j < 26; j++ {
				newWord[i] = byte('a' + j)
				if string(newWord) == endWord {
					return beginWord.seq + 1
				}
				if _, ok := wordSet[string(newWord)]; ok {
					push(string(newWord), beginWord.seq+1)
				}
			}
			newWord[i] = beginWord.word[i]
		}
	}
	return 0
}
