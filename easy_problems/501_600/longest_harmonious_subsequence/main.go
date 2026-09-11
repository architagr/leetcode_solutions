package longestharmonioussubsequence

func findLHS(nums []int) int {
	m := make(map[int]int)
	for _, val := range nums {
		m[val]++
	}
	max := 0
	for key, val := range m {
		if f, ok := m[key-1]; ok {
			c := f + val
			if c > max {
				max = c
			}
		}
	}
	return max
}
