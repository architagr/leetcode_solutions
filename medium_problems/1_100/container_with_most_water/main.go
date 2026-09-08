package containerwithmostwater

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	ans := 0

	for left <= right {
		area := (right - left) * minValue(height[left], height[right])
		ans = maxValue(ans, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}

func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
