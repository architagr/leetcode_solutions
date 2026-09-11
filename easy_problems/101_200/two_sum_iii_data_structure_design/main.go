package twosumiiidatastructuredesign

type TwoSum struct {
	data         map[int]int
	oldValidated map[int]bool
}

func Constructor() TwoSum {
	return TwoSum{
		data:         make(map[int]int),
		oldValidated: make(map[int]bool),
	}
}

func (this *TwoSum) Add(number int) {
	this.data[number]++
}

func (this *TwoSum) Find(value int) bool {
	_, ok := this.oldValidated[value]
	if ok {
		return true
	}
	for num, _ := range this.data {
		count, ok := this.data[value-num]

		if ok {
			if value-num == num && count < 2 {
				continue
			}
			this.oldValidated[value] = true
			return true
		}
	}
	return false
}

/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */
