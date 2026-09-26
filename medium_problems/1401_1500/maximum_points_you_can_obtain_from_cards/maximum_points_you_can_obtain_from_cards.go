package maximumpointsyoucanobtainfromcards

func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	max := 0
	curr := 0
	// Only the count taken from each end matters, not the order, so every
	// valid hand is a window of k that wraps around the end of the array.
	// Start with all k from the back.
	for start := n - k; start < n; start++ {
		curr += cardPoints[start]
	}
	max = maxVal(max, curr)
	start := n - k
	// Slide the window forward: give back one card from the back, take one
	// from the front. k swaps cover the other k configurations. start+i never
	// reaches n, so the % n is only there to make the wrap explicit.
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
