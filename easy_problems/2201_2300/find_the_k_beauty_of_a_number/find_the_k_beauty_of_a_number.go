package findthekbeautyofanumber

import "strconv"

func divisorSubstrings(num int, k int) int {
	k--
	d := 1
	x := 0
	ans := 0
	numsStr := strconv.Itoa(num)
	// d = 10^(k-1): x % d keeps the last k-1 digits, which is how the window
	// drops its leading digit.
	for i := 0; i < k; i++ {
		d *= 10

	}
	// Prime with the first k-1 digits. For k == 1 this parses "" and the
	// ignored error leaves x at 0, which happens to be correct.
	x, _ = strconv.Atoi(string(numsStr[:k]))

	for ; k < len(numsStr); k++ {
		a := int(numsStr[k] - '0')
		// Append a digit on the right; the window is now k digits. Leading
		// zeros fall out naturally: "04" is 0*10 + 4.
		x = (x * 10) + a

		// 0 is not a divisor, and num % 0 would panic; && stops first.
		if x > 0 && num%x == 0 {
			ans++
		}
		x %= d
	}
	return ans
}
