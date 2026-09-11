package thetwosneakynumbersofdigitville

func getSneakyNumbers(nums []int) []int {
	m := map[int]bool{}
	arr := []int{}
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			arr = append(arr, nums[i])
		}

		if len(arr) == 2 {
			return arr
		}

		m[nums[i]] = true

	}
	return arr
}
