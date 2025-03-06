package findmissingandrepeatedvalues

func findMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)
	m := make(map[int]bool, n*n)
	for i := 1; i <= n*n; i++ {
		m[i] = true
	}
	result := make([]int, 2)
	for _, g := range grid {
		for _, v := range g {
			if _, ok := m[v]; !ok {
				result[0] = v
			}
			delete(m, v)
		}
	}
	for k, _ := range m {
		result[1] = k
		break
	}
	return result

}
