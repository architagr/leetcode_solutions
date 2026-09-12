# Min Stack — solution walkthrough

From `main.go`:

```go
type dataNode struct {
	val, min int
}
type MinStack struct {
	dataArr []dataNode
	len     int
}

func Constructor() MinStack {
	return MinStack{dataArr: make([]dataNode, 0), len: 0}
}

func (this *MinStack) Push(val int) {
	defer func() { this.len++ }()
	if len(this.dataArr) == 0 {
		this.dataArr = append(this.dataArr, dataNode{val: val, min: val})
		return
	}
	currMin := this.GetMin()
	if currMin > val {
		currMin = val
	}
	this.dataArr = append(this.dataArr, dataNode{val: val, min: currMin})
}

func (this *MinStack) Pop() {
	defer func() { this.len-- }()
	this.dataArr = this.dataArr[:this.len-1]
}

func (this *MinStack) Top() int    { return this.dataArr[this.len-1].val }
func (this *MinStack) GetMin() int { return this.dataArr[this.len-1].min }
```

## Where a single `min` field breaks

Keeping one minimum and updating it on push works right up until the first pop.

Pop the element that *is* the current minimum, and the field is stale. There is nothing to recompute it from except a scan, and a scan is O(n).

That failure is worth sitting with, because it says what the structure really needs. The minimum is not a property of "the stack". It is a property of **the stack at a given height**, and popping returns you to a height whose minimum you have already discarded.

## Storing it per entry

```go
type dataNode struct {
	val, min int
}
```

Each entry carries the value pushed and the minimum of everything at or below it.

![Step 1](images/walkthrough-1.png)

The first entry is its own minimum; there is nothing below it.

![Step 2](images/walkthrough-2.png)

```go
currMin := this.GetMin()
if currMin > val {
	currMin = val
}
```

0 does not beat -2, so the existing minimum is copied up into the new entry. Every entry has a correct answer for its own height, even when that answer is the same as its neighbour's.

![Step 3](images/walkthrough-3.png)

-3 does beat it, so this entry records a new minimum.

![Step 4](images/walkthrough-4.png)

```go
func (this *MinStack) GetMin() int { return this.dataArr[this.len-1].min }
```

`getMin` is a field read. No scan, no comparison, no bookkeeping.

![Step 5](images/walkthrough-5.png)

And this is the step that shows why the design works. `Pop` does nothing to maintain the minimum:

```go
func (this *MinStack) Pop() {
	defer func() { this.len-- }()
	this.dataArr = this.dataArr[:this.len-1]
}
```

Removing the top exposes an entry whose `min` was computed when it was on top and is still correct, because nothing below it ever changed. The previous minimum is restored for free, because it was never overwritten.

## The `defer` in `Push`

```go
func (this *MinStack) Push(val int) {
	defer func() { this.len++ }()
	...
	currMin := this.GetMin()
```

`Push` calls `GetMin`, which reads `dataArr[this.len-1]`. The deferred increment means `this.len` still describes the stack *before* this push at the moment `GetMin` runs, so it reads the old top. That is what makes the call correct.

It works, and it is doing something subtle with ordering that a reader has to reconstruct. Incrementing explicitly at the end of each branch would be more obvious; the `defer` earns its place by covering the early `return` in the empty case without duplicating the line.

## The redundant length field

`this.len` always equals `len(this.dataArr)`. The struct keeps both.

Nothing is wrong with it and everything stays consistent, but it is two things that must agree rather than one thing that cannot disagree. Dropping the field and using `len(this.dataArr)` throughout would remove the possibility entirely.

## The other standard solution

The common alternative is two stacks: one for values, one holding only the minima, pushed to when a value is less than or equal to the current minimum.

It uses less space when new minima are rare. It also needs care about ties — push on strictly-less-than and duplicate minima get popped too early.

The pair-per-entry version has no such edge case, which is why it is the one to reach for first.

## Complexity

- **Time: O(1)** for all four operations, with no amortisation.
- **Space: O(n)**, two integers per element instead of one.

## Test

`main_test.go` runs the sequence from the problem statement: push -2, 0, -3, read the minimum, pop, then read the top and the minimum again. That last pair is the assertion that matters — it checks the minimum reverted to -2 after the pop.
