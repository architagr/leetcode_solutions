package distinctprimefactorsofproductofarray

func distinctPrimeFactors(nums []int) int {
	m := make(map[int]bool)
	arr := make([]int, 1001)
	for _, val := range nums {
		arr[val] = 1
	}
	for i := 2; i < 1000; i++ {
		for j := 1000; j >= 2; j-- {
			if arr[j] != 1 {
				continue
			}
			if j%i == 0 {
				arr[j] = 0
				arr[j/i] = 1
				m[i] = true
			}
		}
	}
	return len(m)
}
