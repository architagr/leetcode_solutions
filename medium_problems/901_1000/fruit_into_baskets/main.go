package fruitintobaskets

func totalFruit(fruits []int) int {
	// Fruit type -> count in the window. Keys are deleted at zero, so
	// len(fruitBaskets) is the number of types, i.e. baskets in use.
	fruitBaskets := make(map[int]int)
	ans := 0
	start := 0
	for i, fruit := range fruits {
		fruitBaskets[fruit]++
		// Longest window with at most two distinct types: shrink until it fits.
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
