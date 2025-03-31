package paintfence

func numWays(n int, k int) int {

	dpTable := make([]int, n+1)
	return dp(n, k, &dpTable)
}
func dp(i, k int, dpTable *[]int) int {
	if i == 1 {
		return k
	}
	if i == 2 {
		return k * k
	}
	if (*dpTable)[i] == 0 {
		(*dpTable)[i] = (k - 1) * (dp(i-1, k, dpTable) + dp(i-2, k, dpTable))
	}
	return (*dpTable)[i]
}
