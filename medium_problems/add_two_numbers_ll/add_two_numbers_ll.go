package add_two_numbers_ll

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	templ1 := l1
	c := 0
	for l1.Next != nil && l2.Next != nil {
		x := l1.Val + l2.Val + c
		c, x = getDigits(x)
		l1.Val = x

		l1 = l1.Next
		l2 = l2.Next
	}

	if l1.Next == nil && l2 != nil {

		x := l1.Val + l2.Val + c
		c, x = getDigits(x)
		l1.Val = x

		l2 = l2.Next
		for l2 != nil && l2.Next != nil {
			c = addCarryToL2AndAddNewNodeinL1(c, l1, l2)
			l1 = l1.Next
			l2 = l2.Next
		}
		if l2 != nil {
			c = addCarryToL2AndAddNewNodeinL1(c, l1, l2)
			l1 = l1.Next
		}
		addCarryAtLast(c, l1)
	} else if l2.Next == nil && l1 != nil {

		x := l1.Val + l2.Val + c
		c, x = getDigits(x)
		l1.Val = x

		l1 = l1.Next
		for l1 != nil && l1.Next != nil {
			c = addCarryToL1(c, l1)
			l1 = l1.Next
		}
		if l1 != nil {
			c = addCarryToL1(c, l1)
		}
		addCarryAtLast(c, l1)
	}

	return templ1
}
func addCarryAtLast(c int, l1 *ListNode) {
	if c > 0 {
		newNode := new(ListNode)
		newNode.Val = c
		l1.Next = newNode
	}
}
func addCarryToL1(c int, l1 *ListNode) int {
	x := l1.Val + c
	c, x = getDigits(x)
	l1.Val = x

	return c
}
func getDigits(num int) (c int, x int) {
	c = num / 10
	x = num % 10
	return
}

func addCarryToL2AndAddNewNodeinL1(c int, l1 *ListNode, l2 *ListNode) int {
	x := l2.Val + c
	c, x = getDigits(x)
	newNode := new(ListNode)
	newNode.Val = x

	l1.Next = newNode

	return c
}
