package sqrtx

func mySqrt(x int) int {
	return sqrt(1, x, x)
}

func sqrt(start, end, target int) int {
	if target < 2 {
		return target
	}
	if end-start <= 1 {
		s := start * start
		s1 := end * end
		if s <= target && s1 > target {
			return start
		}
		return end
	}
	a := (start + end) / 2
	s := a * a
	if s > target {
		return sqrt(start, a-1, target)
	}
	return sqrt(a, end, target)
}
