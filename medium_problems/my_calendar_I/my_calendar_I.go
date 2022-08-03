package my_calendar_I

type MyCalendar struct {
	Head *Node
}

type Node struct {
	start int
	end   int
	left  *Node
	right *Node
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (this *MyCalendar) Book(start int, end int) bool {
	if this.Head == nil {
		this.Head = &Node{start: start, end: end}
		return true
	}
	return insert(this.Head, start, end)
}

func insert(head *Node, start int, end int) bool {
	if end <= head.start {
		if head.left == nil {
			head.left = &Node{start: start, end: end}
			return true
		}
		return insert(head.left, start, end)
	} else if start >= head.end {
		if head.right == nil {
			head.right = &Node{start: start, end: end}
			return true
		}
		return insert(head.right, start, end)
	}
	return false
}
