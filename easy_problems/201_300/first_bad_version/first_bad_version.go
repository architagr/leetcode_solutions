package first_bad_version

func FirstBadVersion(n int) int {
	max, min := n, 1
	mid := 0
	for max-min > 1 {
		mid = (max + min) / 2
		ok := isBadVersion(mid)
		if ok {
			max = mid
		} else {
			min = mid
		}
	}
	if isBadVersion(min) {
		return min
	}
	return max
}

func isBadVersion(a int) bool {
	return false
}
