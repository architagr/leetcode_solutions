package brickwall

func leastBricks(wall [][]int) int {
	m := make(map[int64]int)
	for i := 0; i < len(wall); i++ {
		sum := int64(0)
		for j := 0; j < len(wall[i])-1; j++ {
			sum += int64(wall[i][j])
			m[sum]++
		}
	}
	min := len(wall)
	for _, c := range m {
		cut := len(wall) - c
		if cut < min {
			min = cut
		}
	}
	return min
}
