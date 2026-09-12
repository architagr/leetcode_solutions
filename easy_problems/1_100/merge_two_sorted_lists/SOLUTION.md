# Merge Two Sorted Lists — solution walkthrough

The whole function, from `merge_sorted_list.go`:

```go
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}

	return dummy.Next
}
```

## The dummy node

```go
dummy := &ListNode{}
tail := dummy
```

`dummy` is never part of the answer and its `Val` is never read. It exists so that `tail` is something real to append to on the very first iteration.

Without it the function needs a separate block before the loop to decide which list supplies the head, duplicating the same comparison that the loop already performs. With it, the first node is handled by the loop like every other node, and the return is `dummy.Next`.

## The loop

The example is `list1 = [1,2,4]`, `list2 = [1,3,4]`.

![Step 1](images/walkthrough-1.png)

Nothing placed yet. `tail` points at the dummy.

![Step 2](images/walkthrough-2.png)

`1 <= 1` is true, so `list1`'s node is taken. Note which one: with equal values the first list wins, which is what keeps the merge stable.

![Step 3](images/walkthrough-3.png)

Now `list1` is on 2 and `list2` is on 1, so `list2`'s node goes next.

![Step 4](images/walkthrough-4.png)

`2 < 3`, so `list1` advances again.

![Step 5](images/walkthrough-5.png)

`list1` runs out after its 4 is placed, and the loop ends.

## Why the loop condition is `&&`

```go
for list1 != nil && list2 != nil {
```

The loop only runs while *both* lists still have nodes, because the comparison inside needs two values to compare. The moment either is empty there is nothing left to decide.

## Attaching the remainder

```go
if list1 != nil {
	tail.Next = list1
} else {
	tail.Next = list2
}
```

Whichever list still has nodes gets attached whole. Not walked, not copied: one pointer assignment.

This is correct because of two facts that are true together at this point. The leftover list is sorted, because it always was. And every node in it is at least as large as everything already placed, because the only way its head survived the loop is by losing every comparison it entered.

The `else` branch covers the case where `list2` is the non-empty one, and it also correctly handles both being empty, since `tail.Next = nil` is exactly right there.

## The three nil checks that are not there

An earlier version of this function opened with:

```go
if list1 == nil && list2 == nil { return nil }
if list1 == nil { return list2 }
if list2 == nil { return list1 }
```

None of these are needed. If both inputs are `nil`, the loop never runs, the `else` branch sets `tail.Next = nil`, and `dummy.Next` is `nil`. If exactly one is `nil`, the loop never runs and the other list is attached whole and returned. The general path already produces the right answer for every empty case, and the guards were three extra branches restating it.

## `<=` versus `<`

```go
if list1.Val <= list2.Val {
```

On ties this takes from `list1`. The result is sorted either way, so a test that only checks sortedness will not catch the difference.

What changes is stability: with `<=`, equal values keep their original relative order, with `list1`'s ahead of `list2`'s. That matters when this merge is the combining step inside a merge sort, which is the usual reason to have written it.

## Complexity

- **Time: O(n + m).** Each node is inspected at most once, and the remainder is attached without traversal.
- **Space: O(1).** One dummy node and two pointers, regardless of input size. Nothing is allocated per element.

## Test

`merge_sorted_list_test.go` covers the worked example `[1,2,4]` with `[1,3,4]`, along with the empty cases from the problem statement: both lists empty, and one empty with the other holding a single node.
