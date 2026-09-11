package permutation_sequence

import "fmt"

func GetPermutation(n int, k int) string {
	nums := buildNums(n)
	fact := buildfact(n)
	ans := ""
	for k > 1 && n > 0 {
		n--
		x := k / fact[n]
		if k%fact[n] == 0 {
			x--
		}
		ans += fmt.Sprintf("%d", nums[x])
		nums = append(nums[:x], nums[x+1:]...)

		k %= fact[n]
	}
	if len(nums) > 0 {
		if k > 0 {
			for i := 0; i < len(nums); i++ {
				ans += fmt.Sprintf("%d", nums[i])
			}
		} else {
			for i := len(nums) - 1; i >= 0; i-- {
				ans += fmt.Sprintf("%d", nums[i])
			}
		}
	}

	return ans
}
func buildfact(n int) map[int]int {
	fact := make(map[int]int)
	mul := 1
	for i := 1; i <= n; i++ {
		mul *= i
		fact[i] = mul
	}
	return fact
}
func buildNums(n int) []int {
	nums := make([]int, n)
	for i := 1; i <= n; i++ {
		nums[i-1] = i
	}
	return nums
}
