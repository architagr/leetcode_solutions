package unique_number_of_occurrences

func UniqueOccurrences(arr []int) bool {
	eleCount := make(map[int]int)
	ferquencyCount := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		eleCount[arr[i]]++
	}
	for _, count := range eleCount {
		if _, ok := ferquencyCount[count]; ok {
			return false
		} else {
			ferquencyCount[count]++
		}
	}
	return true
}
