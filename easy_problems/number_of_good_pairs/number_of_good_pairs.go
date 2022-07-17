package number_of_good_pairs

func CountGoodPairs(nums []int) int {
	m := make(map[int]int)
	count := 0
	for i := 0; i < len(nums); i++ {
		if v, ok := m[nums[i]]; ok {
			count += v
		}
		m[nums[i]]++
	}
	return count
}
