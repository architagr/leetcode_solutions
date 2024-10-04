package subarraysumsdivisiblebyk

func subarraysDivByK(nums []int, k int) int {
	sum, count := 0, 0
	mapSet := make(map[int]int)
	mapSet[0] = 1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		m := (sum%k + k) % k
		count += mapSet[m]
		mapSet[m]++
	}
	return count
}
