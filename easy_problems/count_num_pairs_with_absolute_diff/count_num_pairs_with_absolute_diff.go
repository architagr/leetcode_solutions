package count_num_pairs_with_absolute_diff

func CountDiff(nums []int, k int) int {
	sum := 0
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if v, ok := m[nums[i]]; ok {
			sum += v
		}
		m[nums[i]+k]++
		m[nums[i]-k]++
	}
	return sum
}
