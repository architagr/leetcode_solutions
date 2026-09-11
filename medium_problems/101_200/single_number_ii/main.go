package singlenumberii

func singleNumber(nums []int) int {
	one, two := 0, 0

	for i := 0; i < len(nums); i++ {
		x := nums[i]

		one = (one ^ x) & (^two)
		two = (two ^ x) & (^one)
	}

	return one
}
