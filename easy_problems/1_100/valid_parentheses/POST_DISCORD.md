**365 Days of LeetCode Challenge — Day 58/365**
**Valid Parentheses** (Easy)
🔗 https://leetcode.com/problems/valid-parentheses/

New topic: stacks.

Start with what does NOT work. Counting openers against closers accepts `)(`. Tracking a depth that rises and falls accepts `([)]` - depth never goes negative and ends at zero.

Both fail for the same reason: a single number records how MANY brackets are open, never WHICH. This problem is entirely about which.

```go
func IsValid(A string) bool {
	pairs := map[byte]byte{')': '(', ']': '[', '}': '{'}
	stack := make([]byte, 0, len(A))

	for i := 0; i < len(A); i++ {
		open, isCloser := pairs[A[i]]
		if !isCloser {
			stack = append(stack, A[i])
			continue
		}
		if len(stack) == 0 || stack[len(stack)-1] != open {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}
```

Key the map by the CLOSING bracket. One lookup then answers both questions: is this a closer, and what must be on top if it is.

The emptiness test has to come first in that condition - Go short-circuits `||`, and the index after it would read past the end of an empty slice.

And the last line is not a formality: reaching the end proves every closer matched, not that every opener was closed. `(((` never enters the closer branch.

O(n) time and space.

Full walkthrough: https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/1_100/valid_parentheses/SOLUTION.md
