package minimumarraychangestomakedifferencesequal

func minChanges(nums []int, k int) int {
	n := len(nums)
	diffCount := make(map[int]int)
	oneChangeCount := make([]int, k+1)
	for i := 0; i < n/2; i++ {
		a, b := nums[i], nums[n-i-1]
		d := abs(a - b)
		diffCount[d]++
		minVal, maxVal := min(a, b), max(a, b)
		maxDiff := max(k-minVal, maxVal)
		oneChangeCount[maxDiff]++
	}
	for i := k - 1; i >= 0; i-- {
		oneChangeCount[i] += oneChangeCount[i+1]
	}

	ans := n + 1
	for diff, count := range diffCount {
		one := oneChangeCount[diff] - count
		two := ((n / 2) - oneChangeCount[diff]) * 2
		ans = min(ans, (one + two))
	}
	return ans
}
func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
