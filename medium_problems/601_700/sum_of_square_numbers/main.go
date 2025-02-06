package sumofsquarenumbers

func judgeSquareSum(c int) bool {
	sqrt := getSqrt(c)
	l, r := 0, sqrt

	for l <= r {
		s := (l * l) + (r * r)
		if s == c {
			return true
		} else if s > c {
			r--
		} else {
			l++
		}
	}
	return false
}
func getSqrt(num int) int {
	l, r := 1, num
	m := 1
	for l <= r {
		m = l + (r-l)/2
		s := m * m
		if s == num {
			return m
		} else if s > num {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return m
}
