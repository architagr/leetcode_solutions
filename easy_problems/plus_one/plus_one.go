package plus_one

func plusOne(digits []int) []int {

	n := len(digits)
	x := 1
	flag := true
	for flag {
		if digits[n-1]+x >= 10 {
			x = (digits[n-1] + x) / 10
			digits[n-1] = (digits[n-1] + x) % 10
			n--
			if x == 0 {
				flag = false
			} else if n == 0 {
				digits = append([]int{x}, digits...)
				flag = false
			}
		} else {
			flag = false
			digits[n-1] = digits[n-1] + x
		}
	}
	return digits
}
