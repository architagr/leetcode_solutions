## Intuition

The last three problems rebuilt trees from traversals, where the order of the list carried
the structure. Here the structure is given explicitly, one edge at a time, but in any order.
Two things are needed: somewhere to find a node by its value, and a way to tell which node is
the root.

A map from value to node handles the first. For each description, look up (or create) the
parent and the child, and attach the child on the side `isLeft` says. Because nodes are
created the first time their value appears, whether as a parent or as a child, edges can
arrive in any order and still join up correctly.

The root is the only node that never appears as a child. So collect every child value in a
set, and the one node whose value isn't in it is the root.

## Builds on

This is the tree-shaped version of building an adjacency map from an edge list, which is how
the graph problems started back on
[Day 36: Find if Path Exists in Graph](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/1901_2000/find_if_path_exists_in_graph/).
In graph terms, the root is the one node with no incoming edge.

The code finds the root by deleting every child from the node map and taking whatever is left.
That mutates the map it built, which is fine since it's local. Iterating a Go map has no
defined order, but with exactly one entry left it doesn't matter.

**Complexity:**
- Time: O(n) for n descriptions.
- Space: O(n) for the node map and the child set.
