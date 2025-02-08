package permutationsii

var fact = map[int]int{
	1: 1,
	2: 2,
	3: 6,
	4: 24,
	5: 120,
	6: 720,
	7: 5040,
	8: 40320,
}

func permuteUnique(nums []int) [][]int {
	result := make([][]int, 0, fact[len(nums)])
	seenNum := make(map[int]map[int]bool, len(nums))
	for i := range nums {
		seenNum[i] = make(map[int]bool)
	}
	f(nums, make([]int, len(nums)), 0, &result, seenNum)
	return result
}

func f(nums, current []int, index int, result *[][]int, seenNum map[int]map[int]bool) {

	if len(nums) == 0 {
		x := append([]int{}, current...)
		*result = append(*result, x)
		return
	}
	for i := 0; i < len(nums); i++ {
		if seenNum[index][nums[i]] {
			continue
		}

		seenNum[index][nums[i]] = true
		nums[0], nums[i] = nums[i], nums[0]
		current[index] = nums[0]
		f(nums[1:], current, index+1, result, seenNum)
		nums[0], nums[i] = nums[i], nums[0]

	}
	seenNum[index] = make(map[int]bool)
}
