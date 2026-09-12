**365 Days of LeetCode Challenge — Day 81/365**
**Ugly Number II** (Medium)
🔗 https://leetcode.com/problems/ugly-number-ii/

Testing each integer is hopeless - the 1690th ugly number is 2,123,366,400, so almost everything you test is a miss. Stop testing candidates, start generating them.

Read the definition forwards: the only way to make an ugly number is to multiply a smaller one by 2, 3 or 5.

The obvious structure is a heap - this whole arc has been about "repeatedly take the smallest". Push 1, pop the smallest, push its three multiples. O(n log n), and it needs a SET, because 6 arrives as both 2x3 and 3x2.

That set is the tell: the structure generates work it then throws away.

```go
val_2 = 2 * dp[count_2]
val_3 = 3 * dp[count_3]
val_5 = 5 * dp[count_5]
min_val = min(val_2, min(val_3, val_5))
if min_val == val_2 { count_2++ }
if min_val == val_3 { count_3++ }
if min_val == val_5 { count_5++ }
```

Three separate ifs, NOT an else-if. When 6 comes up, val_2 is 2x3 and val_3 is 3x2 - both equal 6, so both pointers advance. That is the deduplication: no set, just the fact that a value produced twice consumed both producers.

O(n), no heap, no set.

Days 76-78 made the case for the heap. Day 79 found sorting better. Today an array with three counters wins. Knowing when to put a technique down is half the skill.

Full walkthrough: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/201_300/ugly_number_ii/SOLUTION.md
