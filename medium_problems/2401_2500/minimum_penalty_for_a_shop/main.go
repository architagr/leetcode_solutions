package minimumpenaltyforashop

func bestClosingTime(customers string) int {
	count := len(customers)
	nPenalty := make([]int, count+1)
	y, n := 0, 0

	for i := count - 1; i >= 0; i-- {
		if customers[i] == 'Y' {
			y++
		}

		nPenalty[i] += y
	}
	min := count + 1
	ans := count + 1

	for i := 0; i < count; i++ {
		if customers[i] == 'N' {
			n++
		}
		nPenalty[i+1] += n
		if min > nPenalty[i] {
			min = nPenalty[i]
			ans = i
		}
	}
	if min > nPenalty[count] {
		ans = count
	}
	return ans
}
