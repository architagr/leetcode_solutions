package degreeofanarray

func findShortestSubArray(nums []int) int {
	countHash := make(map[int]int)
	sindexHash := make(map[int]int)
	eindexHash := make(map[int]int)
	max := 0
	for i, val := range nums {
		countHash[val]++
		eindexHash[val] = i
		if _, ok := sindexHash[val]; !ok {
			sindexHash[val] = i
		}
		if countHash[val] > max {
			max = countHash[val]
		}
	}
	minLength := len(nums)
	for num, freq := range countHash {
		if freq == max {
			// Calculate the length of the subarray for this number
			length := eindexHash[num] - sindexHash[num] + 1

			// Update minLength with the minimum subarray length found so far
			if length < minLength {
				minLength = length
			}
		}
	}
	return minLength
}
