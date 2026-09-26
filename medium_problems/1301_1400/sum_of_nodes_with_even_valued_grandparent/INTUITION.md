## Intuition

Nodes don't know their parents in this tree, let alone their grandparents. But the
recursion does, because it just came from there. So each call can hand its children the
two facts they need: "I am your parent" and "my parent is your grandparent".

That's the downward-state pattern from the last two mediums again. The difference is what
travels: not a running max or min, but the last two nodes on the path. Every call shifts
the window by one:

```go
sum += compute(node.Left, node, parent)
```

The child's parent is this node, and the child's grandparent is this node's parent.

## Builds on

- [Day 87: Count Good Nodes in Binary Tree](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/1401_1500/count_good_nodes_in_binary_tree/) — information about the path above a node, passed down as parameters instead of looked up
- [Day 89: Maximum Difference Between Node and Ancestor](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/1001_1100/maximum_difference_between_nodes_and_ancestor/) — the same pass carrying more than one value at once

The grandparent is passed as a pointer, and nil means "there isn't one". That's what makes
the root and its children safe without a special case: the root gets `(nil, nil)`, its
children get `(root, nil)`, and only from the third level down is there a grandparent to
check. The check is `grandParent != nil && grandParent.Val%2 == 0`, with the nil test
first so the value read never happens on nil.

There's a mirror-image way to solve it that some people find more natural: have each
*even* node add up its grandchildren directly, reaching two levels down. That needs no
extra parameters but has to guard four possible grandchildren. Passing the window down
keeps each call looking only at itself.

Only the values are ever read from `parent` and `grandParent`, so passing two ints (with a
sentinel for "none") would work too. Pointers make "none" obvious.

**Complexity:**
- Time: O(n), each node visited once.
- Space: O(h) for the recursion stack.
