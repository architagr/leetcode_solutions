package sumofuniqueelements

func sumOfUnique(nums []int) int {
	sum := 0
	h := make(map[int]int)
	for _, val := range nums {
		h[val]++
		if h[val] == 2 {
			sum -= val
		}
		if h[val] == 1 {
			sum += val
		}
	}
	return sum
}
