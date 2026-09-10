package maximumlevelsumofabinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	queue := make([]*TreeNode, 0)
	push := func(node *TreeNode) {
		queue = append(queue, node)
	}
	pop := func() *TreeNode {
		x := queue[0]
		queue = queue[1:]
		return x
	}
	// Levels are 1-based here because the problem numbers them from 1.
	currentLevel := 1
	maxLevel := 1
	// Seeded from root.Val rather than 0: values can be negative, and a
	// tree whose every level sums negative would never beat a 0 start,
	// leaving maxLevel at its initial value without a comparison ever
	// having chosen it.
	maxSum := root.Val
	sum := 0
	push(root)
	// The only nil ever in the queue is this sentinel - children are
	// pushed only when non-nil - so node == nil means "level boundary".
	push(nil)
	for len(queue) > 0 {
		node := pop()
		if node == nil {
			// Strictly greater, and that is the entire tie-break: the
			// problem wants the SMALLEST level with the maximal sum, so a
			// later level that only matches must not replace this one.
			// With >= the function returns a maximal level, but the wrong
			// one whenever two levels tie.
			if sum > maxSum {
				maxLevel = currentLevel
				maxSum = sum
			}
			currentLevel++
			sum = 0
			// Guarded, or the last sentinel would be re-pushed onto an
			// otherwise empty queue and spin the loop forever.
			if len(queue) > 0 {
				push(nil)
			}
			continue
		}
		sum += node.Val
		if node.Left != nil {
			push(node.Left)
		}
		if node.Right != nil {
			push(node.Right)
		}
	}
	return maxLevel
}
