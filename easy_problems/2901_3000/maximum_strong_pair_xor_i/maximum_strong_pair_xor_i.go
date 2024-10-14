package maximumstrongpairxori

func maximumStrongPairXor(nums []int) int {

	ans := 0
	arr := make([]int, 101)

	for _, val := range nums {
		arr[val] = 1
	}
	for i := 100; i > 0; i-- {
		if arr[i] == 1 {
			for j := (i + 1) / 2; j <= i; j++ {
				x := i ^ j
				if arr[j] == 1 && x > ans {
					ans = x
				}
			}
		}
	}
	return ans
}
