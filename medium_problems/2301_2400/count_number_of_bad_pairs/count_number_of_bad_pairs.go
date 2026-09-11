package count_number_of_bad_pairs

func CountBadPairs(nums []int) int64 {

	n := len(nums)
	x := make([]int, n)
	mapUnique := make(map[int]bool)
	mapRepeating := make(map[int]int)
	for i := 0; i < n; i++ {
		_, ok := mapUnique[nums[i]-i]
		x[i] = nums[i] - i
		if ok {
			mapRepeating[nums[i]-i]++
		} else {
			mapUnique[nums[i]-i] = true
		}
	}
	countGoodPairs := int64(0)
	for _, val := range mapRepeating {
		countGoodPairs += count(int64(val))
	}
	totalPairs := count(int64(n - 1))
	return totalPairs - countGoodPairs
}

func count(n int64) int64 {
	return n * (n + 1) / 2
}
