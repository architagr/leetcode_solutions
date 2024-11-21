package findthepowerofksizesubarraysi

func resultsArray(nums []int, k int) []int {
	n := len(nums)
	newN := n - k + 1
	result := make([]int, newN)
	for i := 0; i < newN; i++ {
		max := nums[i]
		flag := true
		for j := i + 1; j < i+k; j++ {
			if nums[j] != nums[j-1]+1 {
				flag = false
				break
			}
			max = nums[j]
		}
		if !flag {
			result[i] = -1
		} else {
			result[i] = max
		}
	}

	return result
}
