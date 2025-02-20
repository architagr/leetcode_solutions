package minstack

type dataNode struct {
	val, min int
}
type MinStack struct {
	dataArr []dataNode
	len     int
}

func Constructor() MinStack {
	return MinStack{
		dataArr: make([]dataNode, 0),
		len:     0,
	}
}

func (this *MinStack) Push(val int) {
	defer func() {
		this.len++
	}()
	if len(this.dataArr) == 0 {
		this.dataArr = append(this.dataArr, dataNode{
			val: val,
			min: val,
		})
		return
	}
	currMin := this.GetMin()
	if currMin > val {
		currMin = val
	}
	this.dataArr = append(this.dataArr, dataNode{
		val: val,
		min: currMin,
	})
}

func (this *MinStack) Pop() {
	defer func() {
		this.len--
	}()
	this.dataArr = this.dataArr[:this.len-1]
}

func (this *MinStack) Top() int {
	return this.dataArr[this.len-1].val
}

func (this *MinStack) GetMin() int {
	return this.dataArr[this.len-1].min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
