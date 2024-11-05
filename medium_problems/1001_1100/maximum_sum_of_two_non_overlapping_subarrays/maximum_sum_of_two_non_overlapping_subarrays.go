package maximumsumoftwononoverlappingsubarrays

func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	ans, sum := 0, 0
	firstLen--
	for i := 0; i < firstLen; i++ {
		sum += nums[i]
	}
	for i := firstLen; i < len(nums); i++ {
		a, b := 0, 0
		sum += nums[i]
		if i-firstLen >= secondLen {
			a = getMaxSubArraySum(nums, 0, i-firstLen, secondLen)
		}
		if len(nums)-1-i >= secondLen {
			b = getMaxSubArraySum(nums, i+1, len(nums), secondLen)
		}
		if m := findMax(a, b); m > 0 {
			ans = findMax(ans, sum+m)
		}
		sum -= nums[i-firstLen]
	}
	return ans
}

func getMaxSubArraySum(nums []int, start, end, k int) int {
	ans, sum := 0, 0
	k--
	for i := 0; i < k; i++ {
		sum += nums[i+start]
	}
	for i := k; i+start < end; i++ {
		sum += nums[i+start]
		ans = findMax(ans, sum)
		sum -= nums[i+start-k]
	}
	return ans
}

func findMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
