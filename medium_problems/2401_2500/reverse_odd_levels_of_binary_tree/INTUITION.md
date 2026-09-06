# Intuition

Reverse the values on every odd level of a perfect binary tree. The obvious approach is BFS: collect each level into a slice, reverse the odd ones, write the values back. That works and it's easy to reason about, but it allocates a slice per level and the widest level holds half the tree.

There's a way to do it with no extra storage at all, and it comes from restating what "reverse a level" means.

## Reversing a level is just swapping mirror pairs

Take a level with eight nodes. Reversing it sends position 0 to position 7, 1 to 6, 2 to 5, 3 to 4. Every element trades places with the element the same distance in from the other end. So reversing a level is exactly: swap all four mirror pairs.

That reframing is the whole solution, because mirror pairs are something recursion can produce directly. Instead of materialising a level and reversing it, descend the tree two nodes at a time, always holding a pair that mirrors across the centre, and swap each pair you're handed when the level is odd.

The tree being perfect is what makes this safe. Every level is completely full, so every node genuinely has a mirror partner and the two sides of any pair descend in lockstep.

## Keeping the pair mirrored on the way down

This is the part that's easy to get subtly wrong. Given a mirror pair `(l, r)`, their four children need to be regrouped into two mirror pairs, and the pairing is not the obvious one.

The outermost children are `l.Left` and `r.Right`. The innermost are `l.Right` and `r.Left`. So the two recursive calls have to cross on one side:

```
rev(l.Left,  r.Right, d+1)
rev(l.Right, r.Left,  d+1)
```

Pair left-with-left and right-with-right instead and you end up comparing nodes that aren't mirrors, and the level comes out permuted rather than reversed. I find it worth drawing once, because the correct version looks asymmetric and the wrong version looks tidy.

## Depth, and why the swap is only a test

The recursion starts at the root's two children, which are level 1, so the depth argument matches the real level number. `d%2 == 1` then picks out exactly the odd levels.

Every level still gets traversed, even the even ones, because the odd levels below them have to be reached. On even levels the function does nothing but recurse. That's fine, since visiting a node is cheap and every node has to be visited anyway to get to the bottom.

One edge case falls out for free: a single-node tree. `root.Left` is `nil`, the guard fires immediately, and nothing happens, which is correct because there are no odd levels to reverse.

## Complexity

Time is O(n). Every node is visited once and the work per node is a comparison and maybe a swap.

Space is O(h) for the recursion stack. For a perfect binary tree the height is log2(n), so this is O(log n) with no allocation. The BFS version, by contrast, needs O(n) for the level buffers.
