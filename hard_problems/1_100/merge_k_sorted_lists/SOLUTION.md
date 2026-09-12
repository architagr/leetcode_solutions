# Merge k Sorted Lists — solution walkthrough

From `merge_k_sorted_lists.go`:

```go
type nodeHeap []*ListNode

func (h nodeHeap) Len() int           { return len(h) }
func (h nodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h nodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *nodeHeap) Push(x any)        { *h = append(*h, x.(*ListNode)) }
func (h *nodeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func MergeKLists(lists []*ListNode) *ListNode {
	h := &nodeHeap{}
	for _, l := range lists {
		if l != nil {
			*h = append(*h, l)
		}
	}
	heap.Init(h)

	dummy := &ListNode{}
	tail := dummy

	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		tail.Next = node
		tail = node
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}

	tail.Next = nil
	return dummy.Next
}
```

## Two arcs meet at this problem

Day 22 merged **two** sorted lists: compare the heads, take the smaller, advance that list. The comparison was an `if`, because with exactly two candidates that is all a comparison needs to be.

With `k` lists, "take the smallest head" is no longer a single comparison. It is a repeated query for the minimum of a collection that changes after every answer — which is the definition of what a heap is for.

So this is day 22's merge with the `if` replaced by a heap. The dummy head, the splicing, and the shape of the loop all come across unchanged.

## The heap holds heads, not nodes

```go
for _, l := range lists {
	if l != nil {
		*h = append(*h, l)
	}
}
heap.Init(h)
```

![Step 1](images/walkthrough-1.png)

Only one node per list goes in.

This is the observation the complexity rests on. A list's second element cannot be the next smallest overall while its first is still unplaced — the list is sorted, so its own head is smaller. It is not a candidate yet, and it becomes one at exactly the moment its predecessor is placed.

So the heap is size `k`, not size `N`. That is what separates O(N log k) from O(N log N), and on the constraints here — k up to 10⁴, N up to 10⁴ — it matters.

## Two details in that initialisation

**Empty lists are filtered.** A `nil` entry must not be pushed: `Less` would dereference it immediately during the first sift. The problem's examples include `[[]]` and `[]`, so this is a real input, not a hypothetical.

**`Init` rather than k pushes.** Appending everything and heapifying once is O(k) by Floyd's method. Pushing one at a time would be O(k log k). Day 76 flagged this distinction and did not use it; here it is used.

## The loop

```go
for h.Len() > 0 {
	node := heap.Pop(h).(*ListNode)
	tail.Next = node
	tail = node
	if node.Next != nil {
		heap.Push(h, node.Next)
	}
}
```

![Step 2](images/walkthrough-2.png)

Pop the smallest head, splice it on, and push its successor — which has just become a candidate.

![Step 3](images/walkthrough-3.png)

![Step 4](images/walkthrough-4.png)

The heap stays at size k throughout. One node leaves, at most one enters.

## Why not scan the k heads instead

Looking at all `k` heads and taking the minimum is correct. It costs O(k) per placed node, so O(N·k) overall.

The waste is that the scan re-examines the same heads every round. Between two rounds only **one** head changed — the list that was just advanced. A heap keeps the ordering between rounds and pays O(log k) to absorb that single change.

## The dummy head, again

```go
dummy := &ListNode{}
tail := dummy
```

Straight from day 22. Without it the loop needs a branch for "is this the first node?" on every iteration, to decide whether to set `head` or `tail.Next`.

## Splicing, and the cycle it can create

```go
tail.Next = node
tail = node
```

The nodes already exist. Copying their values into fresh nodes would cost O(N) memory to produce something that pointer rewriting produces for free — the rule from day 22, and the one 206 and 21 were corrected for.

But splicing has a consequence:

```go
tail.Next = nil
```

Every spliced node arrives with a `Next` still pointing into its original list. In practice the final node popped is the overall maximum, which is the tail of its own list, so its `Next` is already `nil`.

The explicit cut is belt and braces — and the test for this file walks the result with a step cap so that a cycle fails the test rather than hanging it. A merge that reuses nodes is exactly the shape of code where an unterminated list becomes an infinite one.

![Step 5](images/walkthrough-5.png)

## Complexity

- **Time: O(N log k).** Each of N nodes is pushed once and popped once, on a heap that never exceeds k. Plus O(k) to heapify.
- **Space: O(k)** for the heap. Nothing is allocated per node; the output reuses the input.

Compare: scanning is O(N·k), and concatenating everything and sorting is O(N log N).

## Test

`merge_k_sorted_lists_test.go` has the worked example. `heap_test.go` covers no lists, all lists empty, empties interleaved, all-equal values, and 1000 random trials against sorting — plus an assertion that the result reuses the input nodes rather than copying them.
