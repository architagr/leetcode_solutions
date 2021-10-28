package two_sum

func twoSum(nums []int, target int) []int{
	var dataMap map[int]int = make(map[int]int)

	for i, val := range nums {
		var diff = target - val
		if data, ok:= dataMap[diff]; ok{
			return []int{data, i}
		}
		dataMap[val] = i
	}
	return []int{-1, -1}
}