# Flatten Nested List Iterator — solution walkthrough

From `main.go`:

```go
type frame struct {
	list []*NestedInteger
	i    int
}

type NestedIterator struct {
	stack []*frame
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{stack: []*frame{{list: nestedList}}}
}

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

func (this *NestedIterator) Next() int {
	this.HasNext()
	top := this.stack[len(this.stack)-1]
	v := top.list[top.i].GetInteger()
	top.i++
	return v
}
```

## What this is not

The obvious solution walks the whole structure in the constructor, copies every integer into a slice, and indexes it in `Next`.

That returns the right values and is not an iterator. It does all the work up front whether the caller wants one element or all of them, and it keeps a copy of every integer alive for the lifetime of the object.

The distinction only bites in the cases iterators exist for: a structure too large to copy, or a caller that stops early. It is worth writing the lazy version anyway, because the lazy version is the one that transfers.

## The state

```go
type frame struct {
	list []*NestedInteger
	i    int
}
```

One frame per level of nesting: which list, and how far into it.

The stack of frames is the path from the outer list down to the cursor. That is the same thing a BST iterator keeps — the path from the root to the current node — which is why day 55 and this are the same design.

## Construction does nothing

```go
func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{stack: []*frame{{list: nestedList}}}
}
```

![Step 1](images/walkthrough-1.png)

One frame wrapping the outer list, cursor at zero. No walking, no allocation proportional to the input. O(1) however deep or wide the structure is.

## `HasNext` is where the work happens

This is the part that surprises people. `HasNext` sounds like a passive question and is in fact the method that moves the iterator.

```go
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
```

Three cases, and the loop runs until one settles it:

- **The current list is finished.** Pop the frame and continue in the parent.
- **The cursor is on an integer.** Stop. There is something to return.
- **The cursor is on a nested list.** Descend into it.

![Step 2](images/walkthrough-2.png)

![Step 3](images/walkthrough-3.png)

Once the cursor is on an integer the loop returns immediately, which is what makes repeated calls to `HasNext` free. It does not re-advance anything, because the first case it examines already satisfies it.

![Step 4](images/walkthrough-4.png)

The inner list is exhausted, so its frame is popped and the parent resumes.

![Step 5](images/walkthrough-5.png)

## The line that prevents an infinite loop

```go
top.i++
this.stack = append(this.stack, &frame{list: item.GetList()})
```

The parent's cursor is stepped past the nested list **before** descending into it.

Descend first and the parent's cursor still points at the list just entered. When that inner list finishes and its frame is popped, the parent looks at its cursor, sees the same nested list, and descends again. The iterator never terminates.

Stepping first means the parent resumes *after* the list when control returns to it, which is exactly what step 5 shows.

This is the third appearance of one idea. Day 37 sank a grid cell before recursing into neighbours. Day 42 recorded a cloned node in the map before cloning its neighbours. All three work because the thing is marked as handled before it is entered.

## `Next` positions itself

```go
func (this *NestedIterator) Next() int {
	this.HasNext()
	...
}
```

`Next` calls `HasNext` rather than assuming the caller did.

The problem promises that a value exists when `Next` is called — it does not promise `HasNext` was asked first. A caller who knows there are three elements may well call `Next` three times directly, and without this line the second call would index past the end of a finished frame.

It costs nothing when the cursor is already positioned, because `HasNext` returns on its first check.

## Complexity

- **Time: O(1)** for construction. Over a full iteration each element is visited once and each frame pushed and popped once, so `Next` is amortised O(1). A single call can do more work when it has to climb out of several finished lists.
- **Space: O(d)** where `d` is the maximum nesting depth. Not O(n) — the elements are never copied.

## Test

`main_test.go` covers the worked examples. `lazy_test.go` pins the two properties the design exists for: the constructor advances nothing, and `HasNext` can be called repeatedly without consuming a value.
