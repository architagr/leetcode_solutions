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

// isBadVersion stands in for the API LeetCode supplies. It is a var so a
// test can point it at a known first bad version.
var isBadVersion = func(a int) bool {
	return false
}
