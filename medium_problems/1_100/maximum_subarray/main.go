package maximumsubarray

func maxSubArray(nums []int) int {
	// Initialize our variables using the first element.
	currentSubarray := nums[0]
	maxSubarray := nums[0]
	// Start with the 2nd element since we already used the first one.
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		// If current_subarray is negative, throw it away. Otherwise, keep adding to it.
		if currentSubarray < 0 {
			currentSubarray = num
		} else {
			currentSubarray += num
		}
		if maxSubarray < currentSubarray {
			maxSubarray = currentSubarray
		}
	}
	return maxSubarray
}
