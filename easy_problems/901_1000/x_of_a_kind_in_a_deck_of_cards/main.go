package xofakindinadeckofcards

func hasGroupsSizeX(deck []int) bool {
	cnt := make(map[int]int, 0)
	for _, v := range deck {
		cnt[v]++
	}
	res := cnt[deck[0]]
	for _, i := range cnt {
		res = gcd(res, i)
	}
	return res >= 2
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
