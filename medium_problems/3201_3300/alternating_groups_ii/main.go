package alternatinggroupii

func numberOfAlternatingGroups(colors []int, k int) int {
	// No group is longer than k, so copying the first k-1 tiles to the end is
	// enough to make every circular group a plain window. (If colors had spare
	// capacity, append would write into the caller's array; input slices don't.)
	extendedColors := append(colors, colors[:k-1]...)
	n := len(extendedColors)
	result := 0
	// Initialize the bounds of the sliding window
	left, right := 0, 1

	for right < n {
		// Check if the current color is the same as the last one
		if extendedColors[right] == extendedColors[right-1] {
			// Pattern breaks, reset window from the current position
			left = right
			right++
			continue
		}

		// Expand the window to the right
		right++

		// Skip counting sequence if its length is less than k
		if right-left < k {
			continue
		}

		// Record a valid sequence and shrink window from the left to search for more
		result++
		left++
	}

	return result
}
