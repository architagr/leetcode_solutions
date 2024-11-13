package maximumpointsyoucanobtainfromcards

func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	max := 0
	curr := 0
	for start := n - k; start < n; start++ {
		curr += cardPoints[start]
	}
	max = maxVal(max, curr)
	start := n - k
	for i := 0; i < k; i++ {
		curr -= cardPoints[(start+i)%n]
		curr += cardPoints[i]
		max = maxVal(max, curr)
	}
	return max
}

func maxVal(a, b int) int {
	if a > b {
		return a
	}
	return b
}
