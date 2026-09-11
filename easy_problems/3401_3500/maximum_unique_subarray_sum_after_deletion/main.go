package maximumuniquesubarraysumafterdeletion

func maxSum(nums []int) int {
	m := make(map[int]bool, len(nums))
	sum := 0
	maxNeg := -200
	count := 0
	for _, n := range nums {
		if n < 0 {
			count++
			if n > maxNeg {
				maxNeg = n
			}
			continue
		}
		if _, found := m[n]; !found {
			sum += n
			m[n] = true
		}
	}
	if count == len(nums) {
		return maxNeg
	}
	return sum
}
