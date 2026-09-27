---
meta_title: "Max sum of distinct subarrays of length k: sum plus a map"
meta_description: "A fixed window carrying a running sum and a count map that deletes at zero. The window is all-distinct when the map has k keys; then compare its sum."
tags: [golang, sliding-window, hash-map, arrays, leetcode]
---

# Maximum Sum of Distinct Subarrays With Length K

*365 Days of LeetCode Challenge — Day 113/365*

🔗 [LeetCode #2461](https://leetcode.com/problems/maximum-sum-of-distinct-subarrays-with-length-k/) · Difficulty: Medium

This medium is two easies from the start of the arc wearing one coat. If the first week
went in, this should take minutes.

Given an array and `k`, consider every subarray of length `k` whose elements are all
different. Return the largest sum among them, or 0 if there are none.

## Two things to carry

Each window needs two facts: its sum, and whether its elements are distinct.

The sum is Day 92: keep a running total, add the element entering, subtract the one
leaving.

Distinctness is Day 93: keep a count per value, delete a value from the map when its count
drops to zero, and then the number of keys is the number of different values. A window of
`k` elements is all-distinct exactly when the map has `k` keys.

## Builds on

- [Day 93: Substrings of Size Three with Distinct Characters](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/1801_1900/substrings_of_size_three_with_distinct_characters/) — a fixed window whose map deletes keys at zero, so all-distinct is len(map) == window size
- [Day 92: Maximum Average Subarray I](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/601_700/maximum_average_subarray_i/) — the running sum over a fixed window, with the k-- offset

## The code

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

The familiar `k--` offset means "k keys" is written `len(m) == k+1`.

Tracing `[1,5,4,2,9,9,9]` with `k = 3`. `[1,5,4]`: three keys, sum 10.

![Step 1: [1,5,4] is distinct](https://raw.githubusercontent.com/architagr/leetcode_solutions/main/medium_problems/2401_2500/maximum_sum_of_distinct_subarrays_with_length_k/images/walkthrough-1.png)

`[5,4,2]`: sum 11.

![Step 2: [5,4,2], sum 11](https://raw.githubusercontent.com/architagr/leetcode_solutions/main/medium_problems/2401_2500/maximum_sum_of_distinct_subarrays_with_length_k/images/walkthrough-2.png)

`[4,2,9]`: sum 15.

![Step 3: [4,2,9], sum 15](https://raw.githubusercontent.com/architagr/leetcode_solutions/main/medium_problems/2401_2500/maximum_sum_of_distinct_subarrays_with_length_k/images/walkthrough-3.png)

`[2,9,9]` has only two keys, so it doesn't count. Neither does `[9,9,9]`.

![Step 4: windows with a repeat are skipped](https://raw.githubusercontent.com/architagr/leetcode_solutions/main/medium_problems/2401_2500/maximum_sum_of_distinct_subarrays_with_length_k/images/walkthrough-4.png)

## Details

`result` starts at 0, and since every value is positive, any qualifying window beats it.
If no window qualifies, 0 is also the answer the problem asks for, so no special case.

The sum is an `int64` because the signature says so, but it's also necessary: `10^5`
elements of up to `10^5` each can add up to `10^10`, past what 32 bits hold.

O(n) time, O(k) space for the map.

Full code and the step-by-step walkthrough:
[maximum_sum_of_distinct_subarrays_with_length_k](https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/2401_2500/maximum_sum_of_distinct_subarrays_with_length_k/SOLUTION.md)

---

*Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the code and problem statement.*
