## Intuition

Every root-to-leaf path is just a sequence of node values from the top of the tree down
to some leaf. The natural way to enumerate all of them is a **DFS from the root**,
carrying along the "path so far" as you descend, and recording that path the moment you
hit a leaf (a node with no children).

There's a small wrinkle: the root itself needs to be handled specially, since the path
string starts as just the root's value with no arrow before it, while every node after
that gets appended as `"->value"`. That's why this solution has two functions instead
of one recursive helper: `binaryTreePaths` seeds the path with the root's value and
kicks off the recursion into its children, while `foo` handles every node after the
root, where the `"->"` separator is always needed.

**Complexity:**
- Time: O(n²) in the worst case — building each path string involves copying the prefix
  so far, and across a skewed tree of depth n, the total work across all paths is
  O(1+2+...+n) = O(n²). For a balanced tree it's closer to O(n log n).
- Space: O(n) for the recursion stack in the worst case (a skewed tree), plus O(n) to
  store the output paths themselves.
