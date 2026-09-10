365 Days of LeetCode Challenge — Day 51/365

Deepest Leaves Sum (Medium)
🔗 https://leetcode.com/problems/deepest-leaves-sum/

The obvious solution is two passes: find the maximum depth, then walk again summing
everything at that depth.

One pass is enough, and the trick is being willing to throw work away. Carry the level
down, keep two ints, and at every leaf compare. Deeper than anything seen so far? Replace
the sum — everything accumulated until now belonged to a shallower level. Equal? Add.
Shallower? Ignore it.

That replacement is why the final depth never has to be known in advance. Discovering a
deeper leaf invalidates the previous answer outright, so whatever survives to the end was
accumulated at the true maximum.

Full breakdown in today's newsletter article ⬇

#DSA #LeetCode #100DaysOfCode #CodingInterview #BinaryTree #DFS #Golang
