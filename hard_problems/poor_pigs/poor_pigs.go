package poor_pigs

import "math"

func PoorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	t := (minutesToTest / minutesToDie) + 1
	pig := 0
	for int(math.Pow(float64(t), float64(pig))) < buckets {
		pig++
	}
	return pig
}
