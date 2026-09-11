package minimumroundstocompletealltasks

func minimumRounds(tasks []int) int {
	difficultyMap := make(map[int]int)

	for _, val := range tasks {
		difficultyMap[val]++
	}
	count := 0
	for _, c := range difficultyMap {
		if c == 1 {
			return -1
		}
		x := c % 3
		count += c / 3
		if x > 0 {
			count++
		}
	}
	return count
}
