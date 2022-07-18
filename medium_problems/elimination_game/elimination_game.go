package elimination_game

func LastRemaning(n int) int {
	return leftToRight(n)
}

func leftToRight(n int) int {
	if n <= 2 {
		return n
	}
	return 2 * rightToLeft(n/2)
}

func rightToLeft(n int) int {
	if n <= 2 {
		return 1
	}
	if n%2 == 1 {
		return 2 * leftToRight(n/2)
	}
	return 2*leftToRight(n/2) - 1
}
