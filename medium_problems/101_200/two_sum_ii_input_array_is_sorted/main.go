package twosumiiinputarrayissorted

func twoSum(numbers []int, target int) []int {
	var dataMap map[int]int = make(map[int]int)

	for i, val := range numbers {
		var diff = target - val
		if data, ok := dataMap[diff]; ok {
			return []int{data + 1, i + 1}
		}
		dataMap[val] = i
	}
	return []int{0, 0}
}
