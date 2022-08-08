package number_of_arithmetic_triplets

func ArithmeticTriplets(nums []int, diff int) int {

	count := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		l, r := i+1, n-1
		for l < r {
			if nums[l]-nums[i] == diff && nums[r]-nums[l] == diff {
				count++
				l++
			} else if nums[l]-nums[i] != diff {
				l++
			} else if nums[r]-nums[l] != diff {
				r--
			}
		}
	}
	return count
}
