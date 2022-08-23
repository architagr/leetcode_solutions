package find_the_highest_altitude0

func LargestAltitude(gain []int) int {
	max, sum := 0, 0

	for i := 0; i < len(gain); i++ {
		sum += gain[i]
		if sum > max {
			max = sum
		}
	}
	return max
}
