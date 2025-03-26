package deleteandearn

func deleteAndEarn(nums []int) int {
	frequencyMap := make([]int, 10001)
	maxNumber := 0
	for _, num := range nums {
		frequencyMap[num] += num
		maxNumber = maxValue(maxNumber, num)
	}
	a, b := 0, frequencyMap[1]

	for i := 2; i <= maxNumber; i++ {
		b, a = maxValue(b, a+frequencyMap[i]), b

	}
	return b
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
