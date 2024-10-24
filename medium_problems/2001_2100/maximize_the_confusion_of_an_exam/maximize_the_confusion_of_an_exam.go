package maximizetheconfusionofanexam

func maxConsecutiveAnswers(answerKey string, k int) int {

	return findMax(countData(answerKey, k, 'T'), countData(answerKey, k, 'F'))

}

func countData(answerKey string, k int, c byte) int {
	f := 0
	start := 0
	if answerKey[start] == c {
		f++
	}
	ans := 0
	end := 1
	for ; end < len(answerKey); end++ {
		if answerKey[end] == c {
			f++
		}
		if f > k {
			ans = findMax(ans, end-start)
			for f > k {
				if answerKey[start] == c {
					f--
				}
				start++
			}
		}
	}
	if f <= k {
		ans = findMax(ans, end-start)
	}
	return ans
}
func findMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
