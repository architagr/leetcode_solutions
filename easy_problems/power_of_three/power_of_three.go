package power_of_three

func isPowerOfThree(n int) bool {
	// 1162261467 = 3^19
	return n > 0 && 1162261467%n == 0
}
