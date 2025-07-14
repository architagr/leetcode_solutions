package nextgreaterelementi

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	next := make(map[int]int)

	// set the last item with next = -1
	next[nums2[len(nums2)-1]] = -1

	// backward loop for get the next greater
	// element for each item in nums2
	for i := len(nums2) - 2; i >= 0; i-- {
		elNext := nums2[i+1]
		for elNext != -1 && elNext < nums2[i] {
			elNext = next[elNext]
		}
		next[nums2[i]] = elNext
	}

	// Fill the nums1 array with the next greater elements
	for i, num := range nums1 {
		nums1[i] = next[num]
	}

	return nums1
}
