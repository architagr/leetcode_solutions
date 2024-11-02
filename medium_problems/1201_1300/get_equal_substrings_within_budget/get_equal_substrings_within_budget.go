package getequalsubstringswithinbudget

func equalSubstring(s, t string, maxCost int) int {
	maxLength := 0
	sum := 0
	l := 0
	for r := 0; r < len(s); r++ {
		sum += absDiff(s[r], t[r])
		for sum > maxCost {
			sum -= absDiff(s[l], t[l])
			l++
		}
		maxLength = maxVal(maxLength, r-l+1)
	}
	return maxLength
}
func absDiff(a, b byte) int {
	diff := int((a - 'a')) - int((b - 'a'))

	if diff < 0 {
		return diff * -1
	}
	return diff
}
func maxVal(a, b int) int {
	if a > b {
		return a
	}
	return b
}
