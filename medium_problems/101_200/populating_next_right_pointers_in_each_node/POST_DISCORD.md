**365 Days of LeetCode Challenge — Day 50/365**
**Populating Next Right Pointers in Each Node** (Medium)
🔗 https://leetcode.com/problems/populating-next-right-pointers-in-each-node/

"Same level" makes it BFS, and needing to know where a level ends makes it the nil-sentinel BFS from Day 16.

The obvious version walks each level left to right, remembers the previous node, and sets `prev.Next = current`. That works.

This one does something neater — it pushes children **right before left**, so the BFS walks every level backwards. And once the walk runs backwards, the node to your right is simply the node visited just before you:

```go
current.Next = prev
prev = current
```

No lookahead, nothing written into a node already passed. The assignment reads exactly like the requirement.

The boundary falls out too: popping the sentinel sets `prev = current`, and current is nil — so the first node of the next level, which is the *rightmost* one, gets `Next = nil`. That's the rule for the rightmost node, arrived at without a branch testing for it.

Worth noting what it doesn't use: the tree is perfect and nothing here depends on that, so this works unchanged on 117 where it isn't.

O(n) time, O(w) space.

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/101_200/populating_next_right_pointers_in_each_node/SOLUTION.md
