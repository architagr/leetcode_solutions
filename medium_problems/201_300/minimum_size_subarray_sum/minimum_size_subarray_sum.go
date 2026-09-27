package minimumsizesubarraysum

import "math"

func minSubArrayeLen(target int, nums []int) int {
	prefixSum := computePrefixSum(nums)
	if prefixSum[len(prefixSum)-1] < target {
		return 0
	}
	// left is exclusive: the window is (left, right] and its sum is
	// prefix[right] - prefix[left]. All values are positive, so growing the
	// window always raises the sum and shrinking always lowers it.
	left, right, result := -1, 0, math.MaxInt
	getValue := func(index int) int {
		if index < 0 {
			return 0
		}
		return prefixSum[index]
	}
	for right < len(nums) {
		leftValue, rightValue := getValue(left), getValue(right)
		// Big enough: record it and try a shorter window. Too small: grow.
		// Each step moves one edge forward, so at most 2n iterations.
		if rightValue-leftValue >= target {
			result = minValue(result, right-left)
			left++
		} else {
			right++
		}
	}
	return result
}

func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func computePrefixSum(nums []int) []int {
	arr := make([]int, len(nums))
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		arr[i] = sum
	}
	return arr
}
