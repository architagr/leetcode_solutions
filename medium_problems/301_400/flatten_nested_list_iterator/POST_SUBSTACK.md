---
meta_title: "HasNext is the method that does the work"
meta_description: "Flattening into a slice returns the right values and is not an iterator. The lazy version keeps a stack of cursors and advances only when asked."
tags: [golang, stack, iterators, dsa]
---

![Day 63](HERO.png)

*365 Days of LeetCode Challenge — Day 63/365*

**[341. Flatten Nested List Iterator](https://leetcode.com/problems/flatten-nested-list-iterator/)** (Medium)

You are given a nested list of integers — a list whose elements are either integers or other nested lists, to any depth. Implement an iterator that yields the integers in order.

## The solution that passes and misses the point

Here is the first thing most people write, and it is worth taking seriously rather than dismissing:

```go
func Constructor(nestedList []*NestedInteger) *NestedIterator {
	obj := &NestedIterator{}
	obj.flattenList(nestedList)   // walks everything, copies every int into a slice
	return obj
}
func (this *NestedIterator) Next() int { ... this.list[this.index] ... }
```

It is short, it is obviously correct, and it returns exactly the right values in exactly the right order. On LeetCode it is accepted.

It is also not an iterator, and the gap between "returns the right values" and "is an iterator" is the whole reason this problem exists.

An iterator is a promise about *when* work happens. This does all of it in the constructor, before the caller has asked for a single element. And it is a promise about *what is retained*: this holds a complete copy of every integer for as long as the object lives.

Neither costs anything on the inputs in the test cases. Both are wrong in precisely the situations iterators are for:

- A structure far too large to copy — or one being read from somewhere it cannot all be held at once.
- A caller that takes three elements and stops. The eager version has already paid for all of them.

If you write the eager version in an interview, expect the follow-up to be "what if the input does not fit in memory", and the answer is the rest of this piece.

## What an iterator actually has to do

Two things: **do nothing until asked**, and **keep only enough state to resume**.

The second is the interesting half. What is "enough state" for a nested list?

If the cursor is three levels deep, you need to know which list you are in and how far into it — and the same for the list that contains it, and the one containing that, because when the inner list finishes you have to resume the outer one where it left off.

That is a path from the outer list down to the cursor, one entry per level:

```go
type frame struct {
	list []*NestedInteger
	i    int
}

type NestedIterator struct {
	stack []*frame
}
```

And that is a stack, for the usual reason: you always resume the *most recent* unfinished list first.

## The same object as day 55

Worth pausing on, because it is the reason this problem sits where it does in this series.

Day 55 was Binary Search Tree Iterator. Its state is a stack holding the path from the root to the current node, so that when a subtree is exhausted, the iterator can climb back to the parent and continue.

Today's state is a stack holding the path from the outer list to the current element, so that when a nested list is exhausted, the iterator can climb back to the parent list and continue.

Different data, identical design. Once you have written one of these, you have written both — and the general shape is "an iterator over a recursive structure keeps the recursion's call stack explicitly, because it has to be able to stop in the middle of it."

## Construction does nothing

```go
func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{stack: []*frame{{list: nestedList}}}
}
```

![Step 1](images/walkthrough-1.png)

One frame wrapping the outer list, cursor at zero. Nothing is walked and nothing is allocated in proportion to the input. O(1), whatever the structure looks like.

## HasNext is where the work happens

This is the part that surprises people, and it is worth saying loudly: **`HasNext` is not a passive question here. It is the method that moves the iterator.**

```go
func (this *NestedIterator) HasNext() bool {
	for len(this.stack) > 0 {
		top := this.stack[len(this.stack)-1]

		if top.i == len(top.list) {
			this.stack = this.stack[:len(this.stack)-1]
			continue
		}

		item := top.list[top.i]
		if item.IsInteger() {
			return true
		}

		top.i++
		this.stack = append(this.stack, &frame{list: item.GetList()})
	}
	return false
}
```

Three cases, and the loop keeps going until one of them settles the question:

**The current list is finished.** Its frame is popped and the parent resumes.

**The cursor is on an integer.** Stop. There is something to return, and the cursor is sitting on it.

**The cursor is on a nested list.** Descend into it.

Only the middle case returns true, and only the empty stack returns false. Everything else is movement.

![Step 2](images/walkthrough-2.png)

![Step 3](images/walkthrough-3.png)

There is a nice consequence of structuring it this way. Because the loop returns *immediately* when the cursor is already on an integer, calling `HasNext` ten times in a row does the advancing work once and then nine no-ops. It is idempotent without any explicit "have I already advanced" flag, which is the kind of property worth designing for rather than patching in.

![Step 4](images/walkthrough-4.png)

![Step 5](images/walkthrough-5.png)

## The line that prevents an infinite loop

```go
top.i++
this.stack = append(this.stack, &frame{list: item.GetList()})
```

The parent's cursor is advanced past the nested list **before** the iterator descends into it.

Get that order wrong and here is what happens. You descend into the inner list with the parent's cursor still pointing at it. You consume the inner list. Its frame is popped. Control returns to the parent, which looks at its cursor — still pointing at the nested list it just finished — and descends into it again. And again.

Stepping the parent first means that when control comes back, the parent resumes *after* the list, which is what step 5 shows happening.

This is now the third appearance of the same idea in this series, and I want to name the pattern because it generalises well beyond these three problems:

- Day 37 sank a grid cell before recursing into its neighbours.
- Day 42 recorded a cloned node in the map before cloning its neighbours.
- Today steps the parent cursor before descending.

All three are the same rule: **mark the thing as handled before you go into it, not after.** In every case, doing it afterwards means a re-entrant path finds the thing still looking unhandled.

## Next positions itself

```go
func (this *NestedIterator) Next() int {
	this.HasNext()
	top := this.stack[len(this.stack)-1]
	v := top.list[top.i].GetInteger()
	top.i++
	return v
}
```

That first line is deliberate.

The problem guarantees that a value exists whenever `Next` is called. It does not guarantee that `HasNext` was called first — and those are different promises. A caller who already knows there are three elements might reasonably call `Next` three times in a row.

Without the call, the second of those would index into a frame that has already been exhausted and panic. With it, `Next` works regardless of how the caller chooses to drive it, and the cost is nothing when the cursor is already positioned, because `HasNext` returns on its first check.

Removing an ordering requirement between two public methods is almost always worth one redundant call.

## Complexity

- **Time:** O(1) for construction. Across a full iteration every element is visited once and every frame is pushed and popped once, so `Next` is **amortised** O(1). An individual call can do more work than that — if it has to climb out of five finished nested lists before finding the next integer, it pops five frames.
- **Space:** O(d), where d is the maximum nesting depth. Not O(n). The elements are never copied, which is the entire difference from the eager version.

## Builds on

- [Day 55: Binary Search Tree Iterator](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/binary_search_tree_iterator/) — the same design: a stack holding the path to the cursor, advanced only when the caller asks

Full code and the step-by-step walkthrough:
[flatten_nested_list_iterator](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/301_400/flatten_nested_list_iterator/SOLUTION.md)

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
