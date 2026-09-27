package countsubstringsthatsatisfykconstrainti

func countKConstraintSubstrings(s string, k int) int {
	count := 0
	count0 := 0
	count1 := 0
	// left is exclusive: the window is s[left+1..right], length right-left.
	left := -1
	for right := 0; right < len(s); right++ {
		if s[right] == '1' {
			count1++
		} else {
			count0++
		}
		// Valid if EITHER count is within k, so it only breaks when both aren't.
		for left <= right && count0 > k && count1 > k {
			left++
			if s[left] == '1' {
				count1--
			} else {
				count0--
			}
		}

		// Any piece of a valid window is valid, so every substring ending at
		// right inside this window counts: right-left of them.
		count += right - left
	}

	return count
}
