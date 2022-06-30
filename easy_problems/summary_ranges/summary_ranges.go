package summary_ranges

import "fmt"

func SummaryRanges(nums []int) []string {
	min, n := 0, len(nums)
	str := make([]string, 0)
	if n == 0 {
		return str
	}
	min = nums[0]
	for i := 1; i < n; i++ {
		if nums[i-1]+1 != nums[i] {
			str = getStr(str, nums[i-1], min)
			min = nums[i]
		}
	}
	str = getStr(str, nums[n-1], min)
	return str
}
func getStr(str []string, old, min int) []string {
	if old == min {
		str = append(str, fmt.Sprintf("%d", min))
	} else {
		str = append(str, fmt.Sprintf("%d->%d", min, old))
	}
	return str
}
