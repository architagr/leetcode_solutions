package check_if_there_valid_partition_for_the_array

func ValidPartition(nums []int) bool {
	n := len(nums)
	x := make([]int, n)
	for i := 1; i < n; i++ {
		x[i] = nums[i] - nums[i-1]
		if x[i] > 1 || x[i] < 0 {
			return false
		}
	}
	zeroCount, oneCount := 0, 0
	for i := 1; i < n; i++ {

		for i < n && x[i] == 0 {
			zeroCount++
			i++
		}
		for i < n && x[i] == 1 {
			oneCount++
			i++
		}
		if oneCount != 0 {
			if oneCount%2 != 0 {
				return false
			} else if zeroCount > 0 && (zeroCount%4 != 0 && zeroCount%3 != 0 && zeroCount%2 != 0) {
				return false
			}
		} else {
			if zeroCount > 0 && (zeroCount%4 != 0 && zeroCount%3 != 0 && zeroCount%2 != 0) {
				return false
			}
		}
		zeroCount, oneCount = 0, 0
	}

	return true
}
