package minimumrecolorstogetkconsecutiveblackblocks

func minimumRecolors(blocks string, k int) int {
	ans := k
	k--
	whiteCount := 0

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
