package container_with_most_water

func MaxArea(height []int) int {
	i, j := 0, len(height)-1
	ans := (j - i) * (minValue(height[i], height[j]))
	if height[i] > height[j] {
		j--
	} else {
		i++
	}
	for i < j {
		ans = maxValue(ans, ((j - i) * (minValue(height[i], height[j]))))
		if height[i] > height[j] {
			j--
		} else {
			i++
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
