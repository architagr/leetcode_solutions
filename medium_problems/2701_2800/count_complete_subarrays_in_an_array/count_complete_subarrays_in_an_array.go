package countcompletesubarraysinanarray

func countCompleteSubarrays(nums []int) int {
	ele := make(map[int]bool)
	for _, i := range nums {
		ele[i] = true
	}

	k := len(ele)

	fre := make(map[int]int)
	start := 0
	res := 0
	for end := 0; end < len(nums); end++ {
		fre[nums[end]]++
		for start <= end && len(fre) == k {
			res += len(nums) - end

			val := fre[nums[start]]
			if val == 1 {
				delete(fre, nums[start])
			} else {
				fre[nums[start]]--
			}
			start++
		}
	}

	return res
}
