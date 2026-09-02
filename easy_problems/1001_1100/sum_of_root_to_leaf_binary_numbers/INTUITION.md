## Intuition

Each root-to-leaf path spells out a binary number one bit at a time, most significant
bit first — exactly the order a top-down (pre-order-ish) recursion visits nodes in. So
instead of collecting the bits into a list and converting them to a number only once a
leaf is reached, the recursion can maintain the number *as it descends*: at each step,
shift the value built so far one bit to the left (multiply by 2) and drop in the
current node's bit (`0` or `1`) in the units place. That's just how binary numbers are
built digit by digit — `1101` in binary is `((1*2+1)*2+0)*2+1` — and it maps directly
onto "go one level deeper, append one more bit."

Carrying that running value down as a function parameter means each recursive call
only needs to know "the number formed by the path so far," not the whole path itself.
When a leaf is hit, that running value *is* the complete binary number for that path,
so it can be returned directly. Internal nodes don't contribute a value of their own —
they just pass the (now-extended) running value down to their children and sum
whatever comes back up from the left and right subtrees.

**Complexity:**
- Time: O(n) — every node is visited exactly once.
- Space: O(h) for the recursion stack, where h is the tree's height (O(log n) for a
  balanced tree, O(n) for a completely skewed one).
