**365 Days of LeetCode Challenge — Day 7/365**
**Path Sum** (Easy)
🔗 https://leetcode.com/problems/path-sum/

Here's the idea: "root-to-leaf" means the check only happens at a genuine leaf. A
node with children is never compared against the target, even when the running sum
matches it exactly along the way. So the running total rides down the recursion as
an extra argument, and it only gets checked against `targetSum` once you hit a node
with nothing left underneath it.

![Example 1](images/1.jpg)

**Full solution:**
```go
func hasPathSum(root *TreeNode, targetSum int) bool {
	return sum(root, targetSum, 0)
}

func sum(root *TreeNode, target, current int) bool {
	if root == nil {
		return false
	} else if root.Left != nil && root.Right != nil {
		return sum(root.Left, target, current+root.Val) || sum(root.Right, target, current+root.Val)
	} else if root.Left != nil {
		return sum(root.Left, target, current+root.Val)
	} else if root.Right != nil {
		return sum(root.Right, target, current+root.Val)
	}
	return target == root.Val+current
}
```

**Walkthrough** on `[5,4,8,11,null,13,4,7,2,null,null,null,1]`, `targetSum = 22`:
```
            5
          /   \
         4     8
        /     / \
      11     13  4
      / \          \
     7   2          1
```

- `sum(5,22,0)` → two children, descends left into `4` (`current=5`)
- `sum(4,22,5)` → only left child, descends into `11` (`current=9`)
- `sum(11,22,9)` → two children, tries `7` and `2`, each with `current=20`

![Step 1: current sum descends 5 -> 4 -> 11, current becomes 9](images/walkthrough-1.svg)

- `sum(7,22,20)` → leaf, `22 == 7+20` (27)? no, so `false`

![Step 2: at leaf 7, 22 != 27, returns false](images/walkthrough-2.svg)

- `sum(2,22,20)` → leaf, `22 == 2+20` (22)? yes, so `true`

![Step 3: at leaf 2, 22 == 22, returns true](images/walkthrough-3.svg)

`true` climbs back up `11` → `4` → `5`. `||` short-circuits along the way, so the
entire right subtree (`8`, `13`, `4`, `1`) never gets visited. That's my favorite
part of this one: half the tree just doesn't matter once you've found your answer.
Returns `true`.

![Step 4: true bubbles up 2 -> 11 -> 4 -> 5, right subtree never visited](images/walkthrough-4.svg)

O(n) time, O(h) space (tree height).
