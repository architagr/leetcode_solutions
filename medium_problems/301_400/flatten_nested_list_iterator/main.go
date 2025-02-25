package flattennestedlistiterator

type NestedIterator struct {
	originalList []*NestedInteger
	list         []int
	index        int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	obj := &NestedIterator{
		originalList: nestedList,
		index:        0,
	}
	obj.flattenList(nestedList)
	return obj
}

func (this *NestedIterator) flattenList(list []*NestedInteger) {
	for _, item := range list {
		if item.IsInteger() {
			this.list = append(this.list, item.GetInteger())
		} else {
			this.flattenList(item.GetList())
		}
	}
}

func (this *NestedIterator) Next() int {
	defer func() {
		this.index++
	}()
	return this.list[this.index]
}

func (this *NestedIterator) HasNext() bool {
	return this.index < len(this.list)
}
