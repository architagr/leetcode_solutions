## Intuition

The "BST" label is a bit of a red herring for the approach actually used here. A mode is
just "whichever value (or values) show up most often" — and the most direct way to
answer that is to count how many times every value appears, then read off whichever
value(s) hit the highest count. That's a frequency-table problem, and it doesn't
actually require the tree to be sorted at all; it works the same way on any binary tree.

So the solution splits into two clean phases:

1. **Tally every value.** Walk the whole tree once and build a `map[int]int` from value
   → how many times it occurs. This ignores left/right ordering entirely — every node
   just increments its own value's counter.
2. **Find the max, then collect the winners.** Once every value's count is known, scan
   the map for the highest count, then scan it again to collect every value that hits
   that count (there can be ties — multiple values can be equally frequent, which is why
   the problem says "modes," plural).

The BST property (`left <= node <= right`) would let you do this with an in-order
traversal and O(1) extra space by comparing each value to the previous one in sorted
order (the follow-up question in the problem statement is hinting at exactly that). This
implementation trades that space savings for simplicity: a plain hashmap tally is easier
to reason about and get right, at the cost of O(n) extra space for the map.

**Complexity:**
- Time: O(n) — every node is visited once to build the map, and the map (size ≤ n) is
  scanned twice more to find the max count and collect matching keys.
- Space: O(n) for the frequency map, plus O(h) for the recursion stack (h = tree height).
