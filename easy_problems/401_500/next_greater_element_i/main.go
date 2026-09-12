package nextgreaterelementi

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// next[v] is the first value to the right of v in nums2 that beats it.
	next := make(map[int]int, len(nums2))

	// The stack holds values still waiting for something bigger, and it stays
	// decreasing from bottom to top: anything smaller than the incoming value
	// has just found its answer and leaves.
	stack := make([]int, 0, len(nums2))

	for _, v := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < v {
			next[stack[len(stack)-1]] = v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, v)
	}

	// Whatever is still waiting never found anything bigger.
	for _, v := range stack {
		next[v] = -1
	}

	result := make([]int, len(nums1))
	for i, num := range nums1 {
		result[i] = next[num]
	}
	return result
}
