package maximalsquare

func maximalSquare(matrix [][]byte) int {
	ans := 0
	dpTable := make([][]int, len(matrix)+1)
	for i := 0; i < len(dpTable); i++ {
		dpTable[i] = make([]int, len(matrix[0])+1)
	}

	for i := 1; i < len(dpTable); i++ {
		for j := 1; j < len(dpTable[i]); j++ {
			if matrix[i-1][j-1] == '1' {
				dpTable[i][j] = minValue(minValue(dpTable[i][j-1], dpTable[i-1][j]), dpTable[i-1][j-1]) + 1
				ans = maxValue(ans, dpTable[i][j])
			}
		}
	}

	return ans * ans
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}
