package maximizetheconfusionofanexam

func maxConsecutiveAnswers(answerKey string, k int) int {

	// Longest all-F run = longest window with at most k Ts (change them), and
	// vice versa. Solve both and keep the better one.
	return findMax(countData(answerKey, k, 'T'), countData(answerKey, k, 'F'))

}

// countData returns the longest window of answerKey with at most k copies of
// c: the longest run that changing those copies could make uniform.
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
			// The window only grows between violations, so [start, end) is
			// its peak; record before shrinking past one copy of c.
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
