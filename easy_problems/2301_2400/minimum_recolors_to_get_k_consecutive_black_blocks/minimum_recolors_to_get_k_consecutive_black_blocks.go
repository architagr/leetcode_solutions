package minimumrecolorstogetkconsecutiveblackblocks

func minimumRecolors(blocks string, k int) int {
	// Repainting a window costs exactly its number of whites, so the answer
	// is the fewest whites in any window of length k. k itself is the worst
	// case: a window of all whites.
	ans := k
	// k becomes an offset: blocks[i-k] is the oldest block in the window.
	k--
	whiteCount := 0

	// Prime one block short, so the loop is add, check, remove every time.
	for i := 0; i < k; i++ {
		if blocks[i] == 'W' {
			whiteCount++
		}
	}
	for i := k; i < len(blocks); i++ {
		if blocks[i] == 'W' {
			whiteCount++
		}
		if whiteCount < ans {
			ans = whiteCount
		}
		if blocks[i-k] == 'W' {
			whiteCount--
		}
	}

	return ans
}
