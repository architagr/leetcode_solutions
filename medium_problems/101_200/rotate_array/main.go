package rotatearray

func rotate(nums []int, k int) {
	k %= len(nums)

	reverseArray(nums, 0, len(nums)-1)

	reverseArray(nums, 0, k-1)
	reverseArray(nums, k, len(nums)-1)

}

func reverseArray(nums []int, start, end int) {
	for start <= end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
