package grumpybookstoreowner

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	n := len(customers)
	// allGoodCount: prefix sum of every customer, as if never grumpy.
	// grumpyCount: prefix sum of customers actually satisfied (those in
	// non-grumpy minutes), despite the name.
	allGoodCount, grumpyCount := make([]int, n), make([]int, n)
	allGoodCount[0] = customers[0]
	if grumpy[0] == 0 {
		grumpyCount[0] = customers[0]
	}
	for i := 1; i < n; i++ {
		allGoodCount[i] = allGoodCount[i-1] + customers[i]
		currCust := customers[i]
		if grumpy[i] == 1 {
			currCust = 0
		}
		grumpyCount[i] = grumpyCount[i-1] + currCust
	}
	if minutes >= n {
		return allGoodCount[n-1]
	}
	max := allGoodCount[minutes-1] + grumpyCount[n-1] - grumpyCount[minutes-1]
	for i := minutes; i < n; i++ {
		// Window (i-minutes, i]: satisfied-as-usual before it, everyone inside
		// it, satisfied-as-usual after it.
		currCount := grumpyCount[i-minutes] + allGoodCount[i] - allGoodCount[i-minutes] + grumpyCount[n-1] - grumpyCount[i]
		if max < currCount {
			max = currCount
		}

	}
	return max
}
