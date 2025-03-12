package fruitintobaskets

func totalFruit(fruits []int) int {
	fruitBaskets := make(map[int]int)
	ans := 0
	start := 0
	for i, fruit := range fruits {
		fruitBaskets[fruit]++
		for len(fruitBaskets) > 2 {
			fruitBaskets[fruits[start]]--
			if fruitBaskets[fruits[start]] == 0 {
				delete(fruitBaskets, fruits[start])
			}
			start++
		}
		ans = maxValue(ans, i-start+1)
	}
	return ans
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
