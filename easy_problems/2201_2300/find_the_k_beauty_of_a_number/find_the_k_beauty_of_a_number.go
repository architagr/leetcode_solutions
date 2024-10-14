package findthekbeautyofanumber

import "strconv"

func divisorSubstrings(num int, k int) int {
	k--
	d := 1
	x := 0
	ans := 0
	numsStr := strconv.Itoa(num)
	for i := 0; i < k; i++ {
		d *= 10

	}
	x, _ = strconv.Atoi(string(numsStr[:k]))

	for ; k < len(numsStr); k++ {
		a := int(numsStr[k] - '0')
		x = (x * 10) + a

		if x > 0 && num%x == 0 {
			ans++
		}
		x %= d
	}
	return ans
}
