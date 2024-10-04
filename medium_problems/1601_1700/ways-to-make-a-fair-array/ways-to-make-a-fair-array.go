package waystomakeafairarray

func waysToMakeFair(nums []int) int {
	even, odd := evenPrefix(nums), oddPrefix(nums)
	count := 0
	for i := 0; i < len(nums); i++ {
		if oddSum(even, odd, i) == evenSum(even, odd, i) {
			count++
		}
	}
	return count
}

func oddSum(even, odd []int, i int) int {
	result, n := 0, len(even)
	if i > 0 {
		result = odd[i-1]
	}
	result += even[n-1]
	result -= even[i]
	return result
}

func evenSum(even, odd []int, i int) int {
	result, n := 0, len(even)
	if i > 0 {
		result = even[i-1]
	}
	result += odd[n-1]
	result -= odd[i]
	return result
}

func evenPrefix(nums []int) []int {
	n, sum := len(nums), 0
	result := make([]int, n)
	for i := 0; i < n; i = i + 2 {
		sum += nums[i]
		result[i] = sum
		if i+1 <= n-1 {
			result[i+1] = sum
		}
	}
	return result
}

func oddPrefix(nums []int) []int {
	n, sum := len(nums), 0
	result := make([]int, n)
	for i := 1; i < n; i = i + 2 {
		sum += nums[i]
		result[i] = sum
		if i+1 <= n-1 {
			result[i+1] = sum
		}
	}
	return result
}
