**365 Days of LeetCode Challenge — Day 78/365**
**Top K Frequent Elements** (Medium)
🔗 https://leetcode.com/problems/top-k-frequent-elements/

Count the values, then select the top k by count. Only the second step has a decision in it.

Sorting the counts is O(d log d), and the follow-up asks for better than O(n log n). A heap capped at k is O(d log k):

```go
heap.Push(h, entry{ele: ele, count: count})
if h.Len() > k {
	heap.Pop(h)
}
```

Now the part that reads backwards. To keep the k MOST frequent, you want a MIN-heap.

```go
func (h minHeap) Less(i, j int) bool { return h[i].count < h[j].count }
```

Ask what the structure does on each arrival: decide which incumbent the new entry displaces. That is always the LEAST frequent of the ones currently kept - the weakest member, the one on the boundary. So it has to be at the root.

A max-heap of size k holds the most frequent at the root. That entry is in no danger, and finding the one to evict would mean scanning the heap - O(k), which throws away the reason for using one.

Get it backwards and nothing crashes. It silently returns the k LEAST frequent elements.

Also note Less compares `count`, not `ele` - the value is cargo.

O(n + d log k) time.

Full walkthrough: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/301_400/top_k_frequent_elements/SOLUTION.md
