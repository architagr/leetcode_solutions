package minimumnumberofoperationstomoveallballstoeachbox

func minOperations(boxes string) []int {
	n := len(boxes)
	left := make([]int, n)
	right := make([]int, n)

	count := 0
	if boxes[0] == '1' {
		count++
	}
	// move all balls from left to right
	for i := 1; i < n; i++ {
		left[i] = left[i-1] + count
		if boxes[i] == '1' {
			count++
		}
	}
	count = 0
	if boxes[n-1] == '1' {
		count++
	}
	// move all balls from right to left
	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1] + count
		if boxes[i] == '1' {
			count++
		}
	}
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = left[i] + right[i]
	}
	return result
}
