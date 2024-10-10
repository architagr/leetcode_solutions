package defusethebomb

func decrypt(code []int, k int) []int {
	n := len(code)
	ans := make([]int, n)
	if k > 0 {
		sum := 0
		for i := 0; i < k; i++ {
			sum += code[i%n]
		}
		for i := 0; i < n; i++ {
			sum = sum - code[i] + code[(i+k)%n]
			ans[i] = sum
		}
	} else if k < 0 {
		sum := 0
		k *= -1
		for i := 0; i < k; i++ {
			sum += code[(n-1-i)%n]
		}
		for i := n - 1; i >= 0; i-- {
			sum = sum - code[i] + code[((i-k)+n)%n]
			ans[i] = sum
		}
	}
	return ans
}
