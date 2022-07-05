package longest_consecutive_sequence

func LongestConsecutive(nums []int) int {

	mapN := make(map[int]int)
	n, max := len(nums), 0
	if n <= 1 {
		return n
	}
	for i := 0; i < n; i++ {
		mapN[nums[i]] = 1
	}

	for key, val := range mapN {

		x := key + 1
		v, ok := mapN[x]
		for ok {
			//carry the count of current to next value in sequence

			if v == 0 {
				// if count is zero that means it was visited earlier
				// and we should to move to the next value in the sequence
				x++
				v, ok = mapN[x]
			} else {
				// carry the count of current to next value in sequence and then set the count of current
				// number to zero
				mapN[x] = val + v
				mapN[key] = 0
				if max < val+v {
					max = val + v
				}
				ok = false
			}
		}
	}
	if max == 0 {
		return 1
	}
	return max
}
