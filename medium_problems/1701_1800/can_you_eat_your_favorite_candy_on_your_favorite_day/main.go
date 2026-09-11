package canyoueatyourfavoritecandyonyourfavoriteday

func canEat(candiesCount []int, queries [][]int) []bool {
	result := make([]bool, len(queries))
	preSum := prefixSum(candiesCount)

	for i, query := range queries {
		favType := query[0]
		favDay := query[1]
		cap := query[2]
		result[i] = favDay < preSum[favType+1] && preSum[favType] < (favDay+1)*cap
	}
	return result
}

func prefixSum(c []int) []int {
	result := make([]int, len(c)+1)
	result[0] = 0
	for i := 0; i < len(c); i++ {
		result[i+1] = result[i] + c[i]
	}
	return result
}
