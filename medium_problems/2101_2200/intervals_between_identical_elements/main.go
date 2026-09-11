package intervalsbetweenidenticalelements

func getDistances(arr []int) []int64 {
	countMap := make(map[int]int64)
	sumMap := make(map[int]int64)
	n := int64(len(arr))
	result := make([]int64, n)
	var i int64 = 0
	for ; i < n; i++ {
		num := arr[i]
		result[i] += (countMap[num] * i) - sumMap[num]
		countMap[num]++
		sumMap[num] += i
	}
	i = n - 1
	countMap = make(map[int]int64)
	sumMap = make(map[int]int64)
	for ; i >= 0; i-- {
		num := arr[i]
		result[i] += sumMap[num] - (countMap[num] * i)
		countMap[num]++
		sumMap[num] += i
	}
	return result
}
