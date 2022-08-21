package minimum_value_to_get_positive_step_by_step_sum

func MinStartValue(arr []int) int {
	sum := arr[0]
	min := sum
	for i := 1; i < len(arr); i++ {
		sum += arr[i]
		min = minValue(min, sum)
	}
	if min < 0 {
		return (-1 * min) + 1
	}
	return 1
}
func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}
