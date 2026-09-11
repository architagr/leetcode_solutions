package reduce_array_size_to_half

import "sort"

func MinSetSize(arr []int) int {
	m := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		m[arr[i]]++
	}

	type Frequency struct {
		ele, count int
	}
	fCount := make([]Frequency, len(m))
	i := 0
	for ele, count := range m {
		fCount[i] = Frequency{
			ele:   ele,
			count: count,
		}
		i++
	}
	sort.Slice(fCount, func(i, j int) bool {
		return fCount[i].count > fCount[j].count
	})
	i = 0
	sum := 0
	for sum < len(arr)/2 {
		sum += fCount[i].count
		i++
	}
	return i
}
