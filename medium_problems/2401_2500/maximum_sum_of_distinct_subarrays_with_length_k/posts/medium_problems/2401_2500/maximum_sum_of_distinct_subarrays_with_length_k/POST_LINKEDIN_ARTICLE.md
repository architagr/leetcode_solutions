---
meta_title: "Max sum of distinct subarrays of length k: sum plus a map"
meta_description: "A fixed window carrying a running sum and a count map that deletes at zero. The window is all-distinct when the map has k keys; then compare its sum."
---

## 365 Days of LeetCode Challenge — Day 113/365

# Maximum Sum of Distinct Subarrays With Length K

🔗 https://leetcode.com/problems/maximum-sum-of-distinct-subarrays-with-length-k/ · Difficulty: Medium

### The problem

Among subarrays of length `k` whose elements are all different, return the largest sum
(or 0 if there are none).

### The intuition

Two of the very first problems in this arc, stacked. Day 92 slid a running sum over
windows of length `k`. Day 93 slid a count map over windows of three letters and called a
window good when the map had one key per letter.

This window carries both. The sum gives the score. The map, with keys deleted at zero,
tells you whether the window's `k` elements are all different: they are exactly when the
map has `k` keys.

### Builds on

- [Day 93: Substrings of Size Three with Distinct Characters](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/1801_1900/substrings_of_size_three_with_distinct_characters/) — a fixed window whose map deletes keys at zero, so all-distinct is len(map) == window size
- [Day 92: Maximum Average Subarray I](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/601_700/maximum_average_subarray_i/) — the running sum over a fixed window, with the k-- offset

### The solution

```go
func maximumSubarraySum(nums []int, k int) int64 {
	k--
	m := make(map[int]int)
	sum, result := int64(0), int64(0)
	for i := 0; i < k; i++ {
		m[nums[i]]++
		sum += int64(nums[i])
	}
	for i := k; i < len(nums); i++ {
		m[nums[i]]++
		sum += int64(nums[i])
		if len(m) == k+1 {
			result = maxVal(result, sum)
		}
		m[nums[i-k]]--
		if m[nums[i-k]] == 0 {
			delete(m, nums[i-k])
		}
		sum -= int64(nums[i-k])
	}
	return result
}
```

Tracing `[1,5,4,2,9,9,9]` with `k = 3`. `[1,5,4]`: three keys, sum 10.

![Step 1: [1,5,4] is distinct](https://raw.githubusercontent.com/architagr/leetcode_solutions/main/medium_problems/2401_2500/maximum_sum_of_distinct_subarrays_with_length_k/images/walkthrough-1.png)

`[5,4,2]`: sum 11.

![Step 2: [5,4,2], sum 11](https://raw.githubusercontent.com/architagr/leetcode_solutions/main/medium_problems/2401_2500/maximum_sum_of_distinct_subarrays_with_length_k/images/walkthrough-2.png)

`[4,2,9]`: sum 15.

![Step 3: [4,2,9], sum 15](https://raw.githubusercontent.com/architagr/leetcode_solutions/main/medium_problems/2401_2500/maximum_sum_of_distinct_subarrays_with_length_k/images/walkthrough-3.png)

`[2,9,9]` has only two keys, so it doesn't count. Neither does `[9,9,9]`.

![Step 4: windows with a repeat are skipped](https://raw.githubusercontent.com/architagr/leetcode_solutions/main/medium_problems/2401_2500/maximum_sum_of_distinct_subarrays_with_length_k/images/walkthrough-4.png)

Because of the `k--` offset, the check reads `len(m) == k+1`. `result` starting at 0 doubles
as the answer when no window qualifies. The sum is an `int64` to match the return type, and
it needs to be: `10^5` values of up to `10^5` can reach `10^10`.

O(n) time, O(k) space.

Full code and the step-by-step walkthrough:
[maximum_sum_of_distinct_subarrays_with_length_k](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/2401_2500/maximum_sum_of_distinct_subarrays_with_length_k/SOLUTION.md)

#DSA #LeetCode #100DaysOfCode #CodingInterview #SlidingWindow #HashMap #Golang

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
