package valid_perfect_square

func IsPerfectSquare(num int) bool {
	x := sqrt(num)
	return x*x == num
}

func sqrt(x int) int {
	if x == 1 {
		return 1
	}
	left := 1
	right := x

	for left < right {
		mid := left + (right-left)/2 // 8, 4
		if check(mid, x) {
			left = mid + 1 // 5
		} else {
			right = mid // 8
		}
	}
	return left - 1
}

func check(num, target int) bool {
	if num*num <= target {
		return true
	}
	return false
}
