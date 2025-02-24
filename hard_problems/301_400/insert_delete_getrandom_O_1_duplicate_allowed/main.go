package insertdeletegetrandomo1duplicateallowed

import "math/rand"

type RandomizedCollection struct {
	idxMap map[int][]int
	arr    []int
	count  int
}

func Constructor() RandomizedCollection {
	idxMap := make(map[int][]int)
	return RandomizedCollection{idxMap, nil, 0}
}

func (this *RandomizedCollection) Insert(val int) bool {
	exists := len(this.idxMap[val]) > 0
	this.arr = append(this.arr, val)
	this.mapAdd(val, this.count)
	this.count++
	return !exists
}

func (this *RandomizedCollection) Remove(val int) bool {
	idxs := this.idxMap[val]
	if len(idxs) == 0 {
		return false
	}
	idx := idxs[0]
	this.mapRem(val, idx)
	if idx != this.count-1 {
		otherVal := this.arr[this.count-1]
		this.arr[idx] = otherVal
		this.mapRem(otherVal, this.count-1)
		this.mapAdd(otherVal, idx)
	}
	this.arr = this.arr[:this.count-1]
	this.count--
	return true
}

func (this *RandomizedCollection) GetRandom() int {
	idx := rand.Intn(this.count)
	return this.arr[idx]
}

func (this *RandomizedCollection) mapAdd(num, idx int) {
	this.idxMap[num] = append(this.idxMap[num], idx)
}

func (this *RandomizedCollection) mapRem(num, idx int) {
	arr := this.idxMap[num]
	for i := 0; i < len(arr); i++ {
		if arr[i] == idx {
			arr = append(arr[:i], arr[i+1:]...)
			break
		}
	}
	this.idxMap[num] = arr
}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
