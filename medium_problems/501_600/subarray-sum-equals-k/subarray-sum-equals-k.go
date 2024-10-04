package subarraysumequalsk

func subarraySum(nums []int, k int) int {
	sum, count := 0, 0
	var preMap = make(map[int]int)
	preMap[0] = 1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		count += preMap[sum-k]
		preMap[sum]++

	}
	return count
}
