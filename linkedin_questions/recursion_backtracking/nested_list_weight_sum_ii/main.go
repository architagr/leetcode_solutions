package nestedlistweightsumii

func depthSumInverse(nestedList []*NestedInteger) int {
	maxDepth := maxDepth(1, nestedList)

	return backtracking(1, maxDepth, nestedList)
}

func backtracking(currentDepth, maxDepth int, nestedList []*NestedInteger) int {
	if len(nestedList) == 0 {
		return 0
	}
	sum := 0
	for _, item := range nestedList {
		if item.IsInteger() {
			sum += (maxDepth - currentDepth + 1) * item.GetInteger()
		} else {
			sum += backtracking(currentDepth+1, maxDepth, item.GetList())
		}
	}
	return sum
}

func maxDepth(depth int, nestedList []*NestedInteger) int {
	if len(nestedList) == 0 {
		return depth
	}
	max := depth
	for _, l := range nestedList {
		if !l.IsInteger() {
			max = maxVal(max, maxDepth(depth+1, l.GetList()))
		}
	}
	return max
}

func maxVal(a, b int) int {
	if a > b {
		return a
	}
	return b
}
