**365 Days of LeetCode Challenge — Day 59/365**
**Next Greater Element I** (Easy)
🔗 https://leetcode.com/problems/next-greater-element-i/

The brute force scans right from each element. O(n^2), and what it wastes is specific: the scans repeat each other across the same descending runs.

Flip the question. Instead of "what is to the right of me", ask as each value arrives: which earlier values does this one answer?

```go
for _, v := range nums2 {
	for len(stack) > 0 && stack[len(stack)-1] < v {
		next[stack[len(stack)-1]] = v
		stack = stack[:len(stack)-1]
	}
	stack = append(stack, v)
}
```

The stack holds values still waiting for something bigger, and it is always DECREASING - if an earlier value were smaller than a later one, the later one would have resolved it on arrival. A stack maintaining an ordering invariant like that is a monotonic stack.

Obvious objection: does that inner loop make it quadratic? No. Every value is pushed once and popped at most once, so across the whole run the inner loop runs at most n times total. The brute force RE-READS values it rejected; this CONSUMES them.

One thing to check before reusing this: the answers are keyed by value, which is safe only because the problem guarantees values are unique.

O(n + m) time.

Full walkthrough: https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/401_500/next_greater_element_i/SOLUTION.md
