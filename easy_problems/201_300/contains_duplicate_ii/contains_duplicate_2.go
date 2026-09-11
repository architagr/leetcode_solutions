package contains_duplicate_2

func containsDuplicate2(nums []int, k int) bool {
	mapA := make(map[int]int)
	n := len(nums)
	for i := 0; i < n; i++ {
		if index, ok := mapA[nums[i]]; ok {
			if i-index <= k {
				return true
			}
		}
		mapA[nums[i]] = i

	}
	return false
}
