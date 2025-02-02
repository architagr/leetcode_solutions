package canpalceflower

func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			i++
			continue
		}
		if i+1 < len(flowerbed) {
			if flowerbed[i+1] != 1 {
				count++
				i++
			}
		} else {
			count++
		}
	}
	return count >= n
}
