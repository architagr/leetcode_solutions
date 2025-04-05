package singleelementinasortedarray

func singleNonDuplicate(nums []int) int {
	lo := 0
	hi := len(nums) - 1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if mid%2 == 1 {
			mid--
		}
		if nums[mid] == nums[mid+1] {
			lo = mid + 2
		} else {
			hi = mid
		}
	}
	return nums[lo]
}
