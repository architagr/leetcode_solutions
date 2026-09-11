package minimumamountoftimetocollectgarbage

func garbageCollection(garbage []string, travel []int) int {
	pCount, pLastIndex := 0, 0
	mCount, mLastIndex := 0, 0
	gCount, gLastIndex := 0, 0
	count := 0
	prefixSum := make([]int, len(travel))

	for i := 0; i < len(travel); i++ {
		count += travel[i]
		prefixSum[i] = count
	}
	for index, g := range garbage {
		for i := 0; i < len(g); i++ {
			if g[i] == 'M' {
				mCount++
				mLastIndex = index
			} else if g[i] == 'G' {
				gCount++
				gLastIndex = index
			} else if g[i] == 'P' {
				pCount++
				pLastIndex = index
			}
		}
	}

	count = mCount + gCount + pCount
	if mLastIndex > 0 {
		count += prefixSum[mLastIndex-1]
	}

	if gLastIndex > 0 {
		count += prefixSum[gLastIndex-1]
	}

	if pLastIndex > 0 {
		count += prefixSum[pLastIndex-1]
	}

	return count

}
