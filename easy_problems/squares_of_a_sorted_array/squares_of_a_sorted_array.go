package squares_of_a_sorted_array

func SortedSquares(nums []int) []int {
	pivotIndex := 0
	arr := make([]int, len(nums))
	for ; pivotIndex < len(nums) && nums[pivotIndex] < 0; pivotIndex++ {
		continue
	}
	i := pivotIndex
	pivotIndex -= 1
	for k := 0; k < len(nums); k++ {
		if i == len(nums) || (pivotIndex >= 0 && (nums[pivotIndex]*-1) < nums[i]) {
			arr[k] = nums[pivotIndex] * nums[pivotIndex]
			pivotIndex--
		} else {
			arr[k] = nums[i] * nums[i]
			i++
		}
	}
	return arr
}
