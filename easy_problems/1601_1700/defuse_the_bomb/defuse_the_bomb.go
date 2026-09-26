package defusethebomb

func decrypt(code []int, k int) []int {
	n := len(code)
	// A new slice because every number is replaced simultaneously: sums must
	// read the original values. make zero-fills, which covers k == 0.
	ans := make([]int, n)
	if k > 0 {
		sum := 0
		for i := 0; i < k; i++ {
			sum += code[i%n]
		}
		// sum holds code[i..i+k-1]; dropping code[i] and adding code[i+k]
		// leaves the k numbers after i. The % n wraps past the end.
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
		// The mirror image, walking right to left. +n before % because Go's %
		// keeps the sign of the left operand.
		for i := n - 1; i >= 0; i-- {
			sum = sum - code[i] + code[((i-k)+n)%n]
			ans[i] = sum
		}
	}
	return ans
}
