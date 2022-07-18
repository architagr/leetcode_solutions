package design_an_ordered_stream

type OrderedStream struct {
	current int
	data    []string
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		current: 0,
		data:    make([]string, n),
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.data[idKey-1] = value
	if this.data[this.current] == "" {
		return []string{}
	} else {
		for i, v := range this.data[this.current:] {
			if v == "" {
				temp := this.current
				this.current += i
				return this.data[temp:this.current]
			}
		}
	}
	return this.data[this.current:]
}
