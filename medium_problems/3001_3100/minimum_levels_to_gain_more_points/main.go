package minimumlevelstogainmorepoints

func minimumLevels(possible []int) int {
	n := len(possible)
	sumArr := make([]int, n)
	x := 0
	for i := 0; i < n; i++ {
		if possible[i] == 1 {
			x += 1
		} else {
			x -= 1
		}
		sumArr[i] = x
	}
	for i := 0; i < n-1; i++ {
		if sumArr[i] > x-sumArr[i] {
			return i + 1
		}
	}
	return -1
}
