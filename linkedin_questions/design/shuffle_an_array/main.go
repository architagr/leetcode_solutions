package shuffleanarray

import "math/rand"

type Solution struct {
	original []int
	current  []int
}

func Constructor(nums []int) Solution {
	obj := Solution{
		original: nums,
	}
	_ = obj.Reset()
	return obj
}

func (this *Solution) Reset() []int {
	c := make([]int, len(this.original))
	copy(c, this.original)
	this.current = c
	return c
}

func (this *Solution) Shuffle() []int {
	for i := 0; i < 10; i++ {
		x, y := rand.Intn(len(this.current)), rand.Intn(len(this.current))
		this.current[x], this.current[y] = this.current[y], this.current[x]
	}
	return this.current
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
