package contains_duplicate

func containsDuplicate(nums []int) bool {
	mapA := make(map[int]int)
	n := len(nums)
	for i := 0; i < n; i++ {
		if _, ok := mapA[nums[i]]; ok {
			return true
		} else {
			mapA[nums[i]] = i
		}
	}
	return false
}
