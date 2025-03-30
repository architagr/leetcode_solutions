package longestincreasingsubsequence

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	tails := make([]int, len(nums))
	size := 0

	for _, num := range nums {
		left := binarySearch(tails, num, size)
		tails[left] = num
		if left == size {
			size++
		}
	}

	return size
}

func binarySearch(sub []int, num, size int) int {
	left, right := 0, size
	for left < right {
		mid := left + (right-left)/2
		if sub[mid] < num {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
