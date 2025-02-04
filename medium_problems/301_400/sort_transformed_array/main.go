package sorttransformedarray

func transform(x, a, b, c int) int {
	// Return the transformed result for element 'x'
	return (a * x * x) + (b * x) + c
}

func sortTransformedArray(nums []int, a, b, c int) []int {
	answer := make([]int, len(nums))
	left, right := 0, len(nums)-1
	index := 0

	if a < 0 {
		// When 'downward parabola' we will put the edge element (smaller elements) first.
		for ; left <= right; index++ {
			leftTransformedVal := transform(nums[left], a, b, c)
			rightTransformedVal := transform(nums[right], a, b, c)
			if leftTransformedVal < rightTransformedVal {
				answer[index] = leftTransformedVal
				left++
			} else {
				answer[index] = rightTransformedVal
				right--
			}
		}
	} else {
		for ; left <= right; index++ {
			// When 'upward parabola' or a 'straight line'
			// we will put the edge element (bigger elements) first.
			leftTransformedVal := transform(nums[left], a, b, c)
			rightTransformedVal := transform(nums[right], a, b, c)
			if leftTransformedVal > rightTransformedVal {
				answer[index] = leftTransformedVal
				left++
			} else {
				answer[index] = rightTransformedVal
				right--
			}
		}
		// Reverse the decreasing 'answer' array.
		for i := 0; i < len(answer)/2; i++ {
			answer[i], answer[len(answer)-i-1] = answer[len(answer)-i-1], answer[i]
		}
	}
	return answer
}
