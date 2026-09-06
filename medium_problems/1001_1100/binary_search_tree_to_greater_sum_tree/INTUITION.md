# Intuition

Every node has to become itself plus the sum of every key greater than it. Read that literally and it sounds like an O(n^2) problem: for each node, go find all the bigger nodes and add them up.

The trick is noticing that a BST already knows what "bigger" means.

## The ordering you already have

An in-order traversal of a BST hands you the keys in ascending order. Walk it backwards, right subtree before left subtree, and you get them in descending order instead.

That changes everything. If I'm visiting nodes from largest to smallest, then by the time I arrive at any node, I have already visited exactly the nodes that are greater than it. There is nothing to search for. I just need to have been keeping a total.

So the whole problem collapses to: traverse right, root, left, carry a running sum, and at each node do `node.Val += runningSum` before moving on.

## Carrying the sum without a shared variable

The usual way to hold that running total is a pointer or a closure variable that every call mutates. This solution does it differently: the sum is a parameter going down and a return value coming back up.

`parse(node, parentSum)` returns the running total as it stands after that entire subtree has been processed. The base case, `parse(nil, parentSum)`, returns `parentSum` unchanged, which is what makes the threading work at the edges.

There is a small thing I like here. After `node.Val += right`, the node's new value *is* the running sum, because the new value is the node's key plus everything above it. So the code can pass `node.Val` straight into the left-subtree call rather than tracking a separate accumulator. The tree ends up storing the state the traversal needs.

## Why the right subtree gets the parent's sum unchanged

`parse(node.Right, parentSum)` passes the incoming sum down untouched. That's correct because everything in the right subtree is greater than `node`, so those nodes' totals shouldn't include `node` yet. They get processed first, `node` folds their result into itself afterward, and only then does the left subtree see a sum that includes `node`.

Get that order wrong and the values drift by exactly one subtree's worth, which is an annoying bug to spot because the tree still looks plausible.

## Complexity

Time is O(n). Each node is visited once and does constant work.

Space is O(h) for the recursion stack, where h is the height. A balanced BST gives O(log n); a fully skewed one degrades to O(n). No extra structures beyond the stack, and the conversion happens in place.
