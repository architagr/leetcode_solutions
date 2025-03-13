package zeroarraytransformationii

func minZeroArray(nums []int, queries [][]int) int {
	n, sum, k := len(nums), 0, 0
	differenceArray := make([]int, n+1)

	// Iterate through nums
	for index := 0; index < n; index++ {
		// Iterate through queries while current index of nums cannot equal zero
		for sum+differenceArray[index] < nums[index] {
			k++

			// Zero array isn't formed after all queries are processed
			if k > len(queries) {
				return -1
			}
			left, right, val := queries[k-1][0], queries[k-1][1], queries[k-1][2]

			// Process start and end of range
			if right >= index {
				differenceArray[maxValue(left, index)] += val
				differenceArray[right+1] -= val
			}
		}
		// Update prefix sum at current index
		sum += differenceArray[index]
	}
	return k
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
