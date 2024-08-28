package contiguousarray

// n
func findMaxLength(nums []int) int {
	diff := 0
	max := 0
	mapData := make(map[int]int)
	mapData[0] = -1
	for i, x := range nums {
		if x == 0 {
			diff++
		} else {
			diff--
		}
		if index, ok := mapData[diff]; ok {
			if (i - index) > max {
				max = i - index
			}
		} else {
			mapData[diff] = i
		}

	}
	return max
}

// n^2
func findMaxLength1(nums []int) int {
	n := len(nums)
	oneCount, zeroCount := make([]int, n), make([]int, n)
	sumZero, sumOne := 0, 0
	max := 0
	for i, x := range nums {
		if x == 0 {
			sumZero++
		} else {
			sumOne++
		}
		oneCount[i] = sumOne
		zeroCount[i] = sumZero
		if oneCount[i] == zeroCount[i] {
			max = i + 1
		}
	}
	if oneCount[n-1] == zeroCount[n-1] {
		return n
	}
	for start := 0; start < n; start++ {
		for end := start + 1; end < n; end++ {
			zero := zeroCount[end] - zeroCount[start]
			one := oneCount[end] - oneCount[start]
			length := end - start
			if one == zero && length > max {
				max = length
			}
		}
	}
	return max
}
