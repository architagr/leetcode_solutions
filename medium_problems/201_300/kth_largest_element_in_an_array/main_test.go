package kthlargestelementinanarray

import (
	"math/rand"
	"sort"
	"testing"
)

// Cases are the worked examples from the problem statement on LeetCode,
// plus the boundaries of k and the duplicate case the problem calls out:
// the kth largest is a position in sorted order, not the kth distinct value.
func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{name: "example 1", nums: []int{3, 2, 1, 5, 6, 4}, k: 2, want: 5},
		{name: "example 2", nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, k: 4, want: 4},
		{name: "k is 1, the maximum", nums: []int{3, 2, 1, 5, 6, 4}, k: 1, want: 6},
		{name: "k is n, the minimum", nums: []int{3, 2, 1, 5, 6, 4}, k: 6, want: 1},
		{name: "single element", nums: []int{7}, k: 1, want: 7},
		{name: "all equal", nums: []int{2, 2, 2, 2}, k: 3, want: 2},
		{name: "duplicates count separately", nums: []int{5, 5, 4}, k: 2, want: 5},
		{name: "negatives", nums: []int{-1, -5, -3}, k: 2, want: -3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// The input is copied because the heap consumes what it is given.
			in := append([]int(nil), tt.nums...)
			if got := findKthLargest(in, tt.k); got != tt.want {
				t.Errorf("findKthLargest(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}

// Sorting descending and indexing is the definition of the answer, so it
// makes a reference to check the heap against on random input.
func TestFindKthLargestAgainstSort(t *testing.T) {
	r := rand.New(rand.NewSource(7))
	for trial := 0; trial < 2000; trial++ {
		n := 1 + r.Intn(12)
		nums := make([]int, n)
		for i := range nums {
			nums[i] = r.Intn(9) - 4 // a small range, so duplicates are common
		}
		k := 1 + r.Intn(n)

		ref := append([]int(nil), nums...)
		sort.Sort(sort.Reverse(sort.IntSlice(ref)))
		want := ref[k-1]

		in := append([]int(nil), nums...)
		if got := findKthLargest(in, k); got != want {
			t.Fatalf("findKthLargest(%v, %d) = %d, want %d", nums, k, got, want)
		}
	}
}
