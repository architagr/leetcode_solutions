package movingaveragefromdatastream

type MovingAverage struct {
	arr   []int
	size  int
	total int
}

func Constructor(size int) MovingAverage {
	arr := make([]int, 0, size)
	return MovingAverage{
		arr:   arr,
		total: 0,
		size:  size,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	x := 0
	if len(this.arr) >= this.size {
		x = this.arr[0]
		this.arr = this.arr[1:]
	}
	this.arr = append(this.arr, val)
	this.total = this.total - x + val
	return float64(this.total) / float64(len(this.arr))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
