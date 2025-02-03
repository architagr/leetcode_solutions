package randompickwithweight

import "math/rand"

type Solution struct {
	w    []int
	psum []int // prefixsum
	sm   int
}

func Constructor(w []int) Solution {
	sm := 0
	psum := make([]int, len(w)) // space: O(N)
	for i, v := range w {
		sm += v
		psum[i] = sm
	}
	s := Solution{
		w:    w,
		sm:   sm,
		psum: psum,
	}
	return s
}

// lookup in prefix sum
func (this *Solution) findIndex(n int) int {
	start := 0
	end := len(this.psum) - 1
	for start < end {
		mid := (start + end) / 2
		if this.psum[mid] < n {
			start = mid + 1
		} else if this.psum[mid] > n {
			end = mid
		} else {
			return mid
		}
	}

	return start
}

func (this *Solution) PickIndex() int {
	n := rand.Intn(this.sm) + 1
	i := this.findIndex(n)
	return i

}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
