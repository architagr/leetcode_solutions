package findsmallestlettergreaterthantarget

func nextGreatestLetter(letters []byte, target byte) byte {
	l, r := 0, len(letters)-1
	m := 0
	ans := letters[0]
	for l <= r {
		m = l + (r-l)/2
		if letters[m] <= target {
			// go right
			l = m + 1
		} else if letters[m] > target && m > 0 && letters[m-1] < target {
			return letters[m]
		} else {
			// go left
			ans = letters[m]
			r = m - 1
		}
	}

	for m < len(letters) && letters[m] == target {
		m++
	}
	if m >= len(letters) {
		ans = letters[0]
	}
	return ans
}
