package trapping_rain_water

func Trap(height []int) int {
	arr := make([]int, len(height))
	ans := 0
	max := height[0]
	arr[0] = 0
	for i := 1; i < len(height); i++ {
		if max < height[i] {
			max = height[i]
		}
		arr[i] = max - height[i]
	}
	max = height[len(height)-1]
	for i := len(height) - 1; i >= 0; i-- {
		if max < height[i] {
			max = height[i]
		}
		ans += minValue(arr[i], (max - height[i]))
	}
	return ans
}

func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}
