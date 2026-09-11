package plates_between_candles

func PlatesBetweenCandles(s string, queries [][]int) []int {
	prefixSum := make([]int, len(s))
	leftArr := make([]int, len(s))
	rightArr := make([]int, len(s))
	c, index := 0, -1
	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			c++
		} else {
			index = i
		}

		prefixSum[i] = c
		leftArr[i] = index
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '|' {
			index = i
		}

		rightArr[i] = index
	}
	ans := make([]int, len(queries))
	for i, q := range queries {
		x, y := rightArr[q[0]], leftArr[q[1]]
		if x >= 0 && y >= 0 && x < y {
			ans[i] = prefixSum[y] - prefixSum[x]
		}
	}
	return ans
}
