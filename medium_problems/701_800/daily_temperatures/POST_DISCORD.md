**365 Days of LeetCode Challenge — Day 61/365**
**Daily Temperatures** (Medium)
🔗 https://leetcode.com/problems/daily-temperatures/

Yesterday's monotonic stack with two changes.

The stack holds INDICES, not values - the answer is a distance, and a distance needs positions.

And values repeat. Day 59 keyed its answers by value, safe only because uniqueness was guaranteed. Temperatures repeat constantly; indices are unique by construction.

```go
for i := len(temperatures) - 1; i >= 0; i-- {
	for !stack.IsEmpty() && temperatures[stack.Top()] <= temperatures[i] {
		stack.Pop()
	}
	if stack.IsEmpty() {
		results[i] = 0
	} else {
		results[i] = stack.Top() - i
	}
	stack.Push(i)
}
```

One character decides whether duplicates break this: `<=`, not `<`.

Equal temperatures are not WARMER. A day with the same temperature can never be anyone's answer, so it must be popped. With `<` equal days stay on the stack and the first one found reports a warmer day that is merely equal. On `[70,70]` the answer is `[0,0]`; `<` gives `[1,0]`.

The stack holds days still capable of being somebody's answer. A day stops qualifying the moment a warmer-or-equal day appears to its left.

O(n) time - each index pushed once, popped at most once.

Full walkthrough: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/701_800/daily_temperatures/SOLUTION.md
