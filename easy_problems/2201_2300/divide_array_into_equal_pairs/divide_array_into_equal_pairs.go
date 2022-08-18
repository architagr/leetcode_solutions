package divide_array_into_equal_pairs

func DivideArray(nums []int) bool {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	for _, count := range m {
		if count&1 != 0 {
			return false
		}
	}
	return true
}
