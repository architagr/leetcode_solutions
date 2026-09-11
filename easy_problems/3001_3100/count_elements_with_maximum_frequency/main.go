package countelementswithmaximumfrequency

func maxFrequencyElements(nums []int) int {
	frequencyMap := make(map[int]int, 100)

	maxFrequency := 0
	count := 0
	for _, n := range nums {
		frequencyMap[n]++
		currentFrequency := frequencyMap[n]
		if currentFrequency > maxFrequency {
			maxFrequency = currentFrequency
			count = 1
		} else if currentFrequency == maxFrequency {
			count++
		}
	}

	return count * maxFrequency
}
