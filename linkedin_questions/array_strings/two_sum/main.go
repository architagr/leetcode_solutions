package twosum

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int, len(nums))
	for index, number := range nums {
		expectedExisting := target - number
		if existingIndex, ok := indexMap[expectedExisting]; ok {
			return []int{existingIndex, index}
		}
		indexMap[number] = index
	}
	return []int{}
}
