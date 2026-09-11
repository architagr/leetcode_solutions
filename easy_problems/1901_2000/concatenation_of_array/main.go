package concatenationofarray

func getConcatenation(nums []int) []int {

	l := len(nums)
	ans := make([]int, l*2)

	for i, n := range nums {
		ans[i], ans[l+i] = n, n
	}
	return ans
}
