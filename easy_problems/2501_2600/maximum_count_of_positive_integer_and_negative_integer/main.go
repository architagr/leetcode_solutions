package maximumcountofpositiveintegerandnegativeinteger

func maximumCount(nums []int) int {
	negitiveCount, positiveCount := 0, 0
	for _, num := range nums {
		if num < 0 {
			negitiveCount++
		}
		if num > 0 {
			positiveCount++
		}
	}

	if negitiveCount > positiveCount {
		return negitiveCount
	}
	return positiveCount
}
