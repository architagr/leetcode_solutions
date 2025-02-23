package partitiontokequalsumsubsets

import "sort"

func backtrack(arr []int, index, count, currSum, k, targetSum, mask int, memo map[int]bool) bool {

	n := len(arr)

	// We made k - 1 subsets with target sum and last subset will also have target sum.
	if count == k-1 {
		return true
	}

	// No need to proceed further.
	if currSum > targetSum {
		return false
	}

	// If we have already computed the current combination.
	if v, ok := memo[mask]; ok {
		return v
	}

	// When curr sum reaches target then one subset is made.
	// Increment count and reset current sum.
	if currSum == targetSum {
		ans := backtrack(arr, 0, count+1, 0, k, targetSum, mask, memo)
		memo[mask] = ans
		return ans
	}

	// Try not picked elements to make some combinations.
	for j := index; j < n; j++ {
		if ((mask >> j) & 1) == 0 {
			// Include this element in current subset.
			mask = (mask | (1 << j))

			// If using current jth element in this subset leads to make all valid subsets.
			if backtrack(arr, j+1, count, currSum+arr[j], k, targetSum, mask, memo) {
				return true
			}

			// Backtrack step.
			mask = (mask ^ (1 << j))
		}
	}

	// We were not able to make a valid combination after picking each element from the array,
	// hence we can't make k subsets.
	memo[mask] = false
	return false
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func canPartitionKSubsets(nums []int, k int) bool {
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	if total%k != 0 {
		return false
	}

	// Sort in decreasing order.
	sort.Ints(nums)
	reverse(nums)

	targetSum := total / k
	mask := 0

	// Memoize the ans using taken element's string as key.
	memo := make(map[int]bool)

	return backtrack(nums, 0, 0, 0, k, targetSum, mask, memo)
}
