package longestsubarrayof1safterdeletingoneelement

func longestSubarray(nums []int) int {
	arr := make([]int, 0, len(nums))
	sum := 0
	for _, val := range nums {
		if val == 0 {
			if sum > 0 {
				arr = append(arr, sum)
				sum = 0
			}
			arr = append(arr, val)
		} else {
			sum++
		}
	}
	if sum > 0 {
		arr = append(arr, sum)
	}
	if len(arr) == 1 {
		return len(nums) - 1
	}
	max := 0
	for i, val := range arr {
		if val == 0 {
			left := 0
			if i > 0 {
				left = arr[i-1]
			}
			right := 0
			if i < len(arr)-1 {
				right = arr[i+1]
			}
			if left+right > max {
				max = left + right
			}
		}
	}
	return max
}
