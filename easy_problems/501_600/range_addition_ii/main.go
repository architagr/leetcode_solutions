package rangeadditionii

func maxCount(m int, n int, ops [][]int) int {
	minM, minN := m, n
	for _, op := range ops {
		if op[0] < minM {
			minM = op[0]
		}
		if op[1] < minN {
			minN = op[1]
		}
	}
	return minN * minM
}
