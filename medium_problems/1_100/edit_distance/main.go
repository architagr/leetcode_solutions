package editdistance

func minDistance(word1 string, word2 string) int {
	memo := make([][]int, len(word1)+1)
	for i := range memo {
		memo[i] = make([]int, len(word2)+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	min := func(a, b, c int) int {
		if a < b {
			if a < c {
				return a
			}
			return c
		}
		if b < c {
			return b
		}
		return c
	}
	var minDistanceRecur func(word1, word2 string, word1Index, word2Index int) int
	minDistanceRecur = func(word1, word2 string, word1Index, word2Index int) int {
		if word1Index == 0 {
			return word2Index
		}
		if word2Index == 0 {
			return word1Index
		}
		if memo[word1Index][word2Index] != -1 {
			return memo[word1Index][word2Index]
		}
		var minEditDistance int
		if word1[word1Index-1] == word2[word2Index-1] {
			minEditDistance = minDistanceRecur(
				word1,
				word2,
				word1Index-1,
				word2Index-1,
			)
		} else {
			insertOperation := minDistanceRecur(word1, word2, word1Index, word2Index-1)
			deleteOperation := minDistanceRecur(word1, word2, word1Index-1, word2Index)
			replaceOperation := minDistanceRecur(word1, word2, word1Index-1, word2Index-1)
			minEditDistance = min(insertOperation, deleteOperation, replaceOperation) + 1
		}
		memo[word1Index][word2Index] = minEditDistance
		return minEditDistance
	}
	return minDistanceRecur(word1, word2, len(word1), len(word2))
}
