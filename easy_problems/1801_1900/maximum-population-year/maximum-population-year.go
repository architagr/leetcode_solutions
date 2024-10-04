package maximumpopulationyear

func maximumPopulation(logs [][]int) int {
	minAge, maxPx := 0, 0
	line := make([]int, 2052)

	for _, px := range logs {
		line[px[0]]++
		line[px[1]]--
	}
	countPpl := 0
	for year, count := range line {
		countPpl += count
		if countPpl > 0 && maxPx < countPpl {
			maxPx = countPpl
			minAge = year

		}
	}
	return minAge
}
