package countdistinctnumbersonboard

func distinctIntegers(n int) int {
	if n < 3 {
		return 1
	}
	return n - 1
}
