package decodexoredarray

func decode(encoded []int, first int) []int {
	ans := make([]int, len(encoded)+1)
	ans[0] = first
	for i := 0; i < len(encoded); i++ {
		ans[i+1] = ans[i] ^ encoded[i]
	}

	return ans
}
