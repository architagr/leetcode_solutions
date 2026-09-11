package divideplayersintoteamsofequalskill

func dividePlayers(skill []int) int64 {
	if len(skill) == 2 {
		return int64(skill[0]) * int64(skill[1])
	}
	m := make(map[int64]int64)
	sum := int64(0)
	for _, val := range skill {
		sum += int64(val)
		m[int64(val)]++
	}
	chem := int64(0)
	sum = (sum * int64(2)) / int64(len(skill))
	for val, count := range m {
		v := sum - val
		if x, ok := m[v]; ok && x == count {
			if v == val {
				count /= 2
			}
			chem += count * (v * val)
			delete(m, val)
			delete(m, v)
		} else {
			return -1
		}
	}
	return chem
}
