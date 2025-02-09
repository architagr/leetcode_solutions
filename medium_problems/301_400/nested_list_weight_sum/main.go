package nestedlistweightsum

func depthSum(nestedList []*NestedInteger) int {
	return backtracking(1, nestedList)
}
func backtracking(currentDepth int, nestedList []*NestedInteger) int {
	if len(nestedList) == 0 {
		return 0
	}
	sum := 0
	for _, item := range nestedList {
		if item.IsInteger() {
			sum += currentDepth * item.GetInteger()
		} else {
			sum += backtracking(currentDepth+1, item.GetList())
		}
	}
	return sum
}
