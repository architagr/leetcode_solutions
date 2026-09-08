package permutations

var fact = map[int]int{
	1: 1,
	2: 2,
	3: 6,
	4: 24,
	5: 120,
	6: 720,
}

func permute(nums []int) [][]int {
	result := make([][]int, 0, fact[len(nums)])
	f(nums, make([]int, len(nums)), 0, &result)
	return result
}

func f(nums, current []int, index int, result *[][]int) {
	if len(nums) == 0 {
		x := append([]int{}, current...)
		*result = append(*result, x)
		return
	}
	for i := 0; i < len(nums); i++ {
		nums[0], nums[i] = nums[i], nums[0]
		current[index] = nums[0]
		f(nums[1:], current, index+1, result)
		nums[0], nums[i] = nums[i], nums[0]
	}
}
