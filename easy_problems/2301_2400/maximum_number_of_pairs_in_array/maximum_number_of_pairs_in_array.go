package maximum_number_of_pairs_in_array

func NumberOfPairs(nums []int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	pair, left := 0, 0
	for _, val := range m {
		pair += val / 2
		left += val % 2
	}
	return []int{pair, left}
}
