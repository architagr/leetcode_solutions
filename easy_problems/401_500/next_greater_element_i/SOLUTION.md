# Next Greater Element I — solution walkthrough

From `main.go`:

```go
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	next := make(map[int]int, len(nums2))
	stack := make([]int, 0, len(nums2))

	for _, v := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < v {
			next[stack[len(stack)-1]] = v
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, v)
	}

	for _, v := range stack {
		next[v] = -1
	}

	result := make([]int, len(nums1))
	for i, num := range nums1 {
		result[i] = next[num]
	}
	return result
}
```

## The question, asked the other way round

The natural reading is "for each element, look right until something is bigger". That is O(n²), and it re-reads the same descending runs over and over.

This asks the opposite question. As each value arrives, **which earlier values does it answer?**

A value resolves every earlier value smaller than it that is still waiting, and a resolved value is never looked at again.

## What the stack holds

Values seen so far that have not yet found anything larger.

That set has a property worth stating explicitly, because it is where the name comes from: it is always **decreasing from bottom to top**. If some earlier value were smaller than a later one, the later one would have resolved and removed it on arrival. So a smaller value can never sit below a larger one.

A stack that maintains an ordering invariant like this is a monotonic stack.

## The loop

![Step 1](images/walkthrough-1.png)

With nothing to compare against, the first value just waits.

```go
for len(stack) > 0 && stack[len(stack)-1] < v {
	next[stack[len(stack)-1]] = v
	stack = stack[:len(stack)-1]
}
stack = append(stack, v)
```

![Step 2](images/walkthrough-2.png)

3 arrives and beats the waiting 1, so 1's answer is 3 and it leaves the stack permanently.

![Step 3](images/walkthrough-3.png)

Same again for 3 and 4. Notice the stack never holds a small value underneath a larger one.

![Step 4](images/walkthrough-4.png)

2 beats nothing, so the inner loop does not run and 2 stacks on top of 4. The stack is `[4, 2]` — decreasing, as promised.

![Step 5](images/walkthrough-5.png)

5 clears both waiters in a single iteration: 2 first, then 4.

This is the step that shows why the inner loop is not a hidden O(n). It ran twice here, and both iterations removed a value for good.

## Why this is linear

Each value in `nums2` is pushed exactly once. Each value is popped at most once, because a popped value is never pushed again.

So across the whole outer loop, the inner loop runs at most `n` times in total, however uneven its distribution. That is the difference between this and the brute force: the brute force *re-reads* values it has already rejected, and this one *consumes* them.

## The leftovers

```go
for _, v := range stack {
	next[v] = -1
}
```

Anything still waiting when `nums2` runs out never met a larger value.

## Why a map and not an array

`nums1` is a subset of `nums2` in a different order, so positions do not line up. The answers are computed once over `nums2` and then read back by value.

That is safe only because the problem guarantees all values are unique across both arrays. Without that guarantee, two equal values would share one map entry and collide. It is the first thing to check before reusing this shape elsewhere.

## Complexity

- **Time: O(n + m).** Every value in `nums2` is pushed and popped at most once; then one pass over `nums1`.
- **Space: O(n)** for the stack and the map.

## Test

`main_test.go` covers four cases, including `nums2 = [1,3,4,2,5]` where the final 5 has to resolve two pending values in one iteration, and a case where the answer for some elements is `-1`.
