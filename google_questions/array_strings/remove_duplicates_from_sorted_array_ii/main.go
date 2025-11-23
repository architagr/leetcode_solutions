package removeduplicatesfromsortedarrayii

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 1     // Pointer to iterate through the array
	j := 1     // Pointer to track position for valid elements
	count := 1 // Count of occurrences of the current element

	for i < len(nums) {
		if nums[i] == nums[i-1] {
			count++
			if count > 2 {
				i++
				continue
			}
		} else {
			count = 1
		}
		nums[j] = nums[i]
		j++
		i++
	}
	return j
}
