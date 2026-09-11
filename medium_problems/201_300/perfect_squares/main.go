package perfectsquares

func numSquares(A int) int {

	arr := make([]int, A+1)

	for i := 0; i <= A; i++ {
		arr[i] = -1
	}

	res := count(A, arr)

	return res

}
func count(A int, dp []int) int {
	if A <= 3 {
		return A
	}
	if dp[A] != -1 {
		return dp[A]
	}
	ans := A
	for i := 1; i*i <= A; i++ {
		temp := 1 + count(A-i*i, dp)
		if ans > temp {
			ans = temp
		}
	}
	dp[A] = ans

	return ans
}
