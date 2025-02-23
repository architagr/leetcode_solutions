package shortestworddistanceii

import "math"

type WordDistance struct {
	wordIndex map[string][]int
}

func Constructor(wordsDict []string) WordDistance {
	wordIndex := make(map[string][]int)

	for i, word := range wordsDict {
		indexList, exist := wordIndex[word]
		if !exist {
			indexList = make([]int, 0)
		}
		indexList = append(indexList, i)
		wordIndex[word] = indexList
	}
	return WordDistance{
		wordIndex: wordIndex,
	}
}

func (this *WordDistance) Shortest(word1 string, word2 string) int {
	min := math.MaxInt
	indexList1 := this.wordIndex[word1]
	indexList2 := this.wordIndex[word2]
	l, r := 0, 0

	for l < len(indexList1) && r < len(indexList2) {
		distance := this.abs(indexList1[l] - indexList2[r])
		min = this.min(distance, min)

		if indexList1[l] < indexList2[r] {
			l++
		} else {
			r++
		}
	}
	return min
}
func (this *WordDistance) min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func (this *WordDistance) abs(a int) int {
	if a < 0 {
		a *= -1
	}
	return a
}

/**
 * Your WordDistance object will be instantiated and called as such:
 * obj := Constructor(wordsDict);
 * param_1 := obj.Shortest(word1,word2);
 */
