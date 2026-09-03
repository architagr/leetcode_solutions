**365 Days of LeetCode Challenge — Day 17/365**

**Two Sum IV - Input is a BST** (Easy)
🔗 https://leetcode.com/problems/two-sum-iv-input-is-a-bst/

Strip away the "BST" part and this is plain Two Sum. Walk the tree, keep a hash
map of complements, check each new node against it. The part I like: because
the map is shared across the whole recursion, a match can turn up between two
branches that never even touch. And the code still checks the right subtree
even after the left one already found an answer, since it's not short-circuited.

Full breakdown in today's newsletter article.

#DSA #LeetCode #100DaysOfCode #BinarySearchTree #HashSet #Golang #CodingInterview
