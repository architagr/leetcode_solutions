package countsubstringsthatsatisfykconstrainti

func countKConstraintSubstrings(s string, k int) int {
	count := 0
	count0 := 0
	count1 := 0
	left := -1
	for right := 0; right < len(s); right++ {
		if s[right] == '1' {
			count1++
		} else {
			count0++
		}
		for left <= right && count0 > k && count1 > k {
			left++
			if s[left] == '1' {
				count1--
			} else {
				count0--
			}
		}

		count += right - left
	}

	return count
}
