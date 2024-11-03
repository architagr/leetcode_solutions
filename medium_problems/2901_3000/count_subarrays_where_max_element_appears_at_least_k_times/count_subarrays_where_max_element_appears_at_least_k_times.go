package countsubarrayswheremaxelementappearsatleastktimes

func countSubarrays(nums []int, k int) int64 {
	max, maxCount := 0, 0
	for _, val := range nums {
		if val > max {
			maxCount = 1
			max = val
		} else if val == max {
			maxCount++
		}
	}
	if maxCount < k {
		return int64(0)
	}

	l, r := -1, 0
	ans := int64(0)
	maxCount = 0
	for r < len(nums) {
		if nums[r] == max {
			maxCount++
		}

		for maxCount >= k {
			ans += int64(len(nums)) - int64(r)
			l++
			if nums[l] == max {
				maxCount--
			}
		}
		r++
	}
	return ans
}
