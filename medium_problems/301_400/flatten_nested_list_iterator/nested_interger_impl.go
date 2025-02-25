package flattennestedlistiterator

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
	n    int
	list []*NestedInteger
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (n *NestedInteger) IsInteger() bool {
	return len(n.list) == 0
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (n *NestedInteger) GetInteger() int {
	return n.n
}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {
	n.n = value
}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (n *NestedInteger) Add(elem NestedInteger) {
	if len(n.list) == 0 {
		n.list = make([]*NestedInteger, 0)
	}
	n.list = append(n.list, &elem)
}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n *NestedInteger) GetList() []*NestedInteger {
	return n.list
}
