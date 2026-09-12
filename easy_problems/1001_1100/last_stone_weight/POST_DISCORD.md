**365 Days of LeetCode Challenge — Day 76/365**
**Last Stone Weight** (Easy)
🔗 https://leetcode.com/problems/last-stone-weight/

Heaps start today.

Why sorting is the wrong shape: sort descending, smash the first two, put the remainder back in the right place - that insert is O(n), paid every round. The deeper mismatch is that sorting answers a question nobody asked. This never needs the third-heaviest stone or the order of the rest.

A heap keeps exactly ONE promise: the largest is at the root. Nothing about the order of anything else. Maintaining less is what makes it cheap.

```go
func (h *MaxHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }

for maxHeap.Len() > 1 {
	x := heap.Pop(maxHeap).(int)
	y := heap.Pop(maxHeap).(int)
	if z := x - y; z != 0 {
		heap.Push(maxHeap, z)
	}
}
```

Two things worth knowing about Go's `container/heap`.

It always builds a MIN-heap. You get a max-heap by lying to it - `Less` defined as greater-than. That one inverted comparison is the whole difference.

And your `Pop` removes the LAST element, not the root. The package swaps the root to the end and sifts down first, then calls yours. Returning `old[0]` is a classic bug that fails silently.

O(n log n) time.

Full walkthrough: https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/1001_1100/last_stone_weight/SOLUTION.md
