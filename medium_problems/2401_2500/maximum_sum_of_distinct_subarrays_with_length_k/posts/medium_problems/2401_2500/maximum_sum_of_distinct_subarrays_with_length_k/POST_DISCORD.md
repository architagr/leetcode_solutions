**365 Days of LeetCode Challenge — Day 113/365**
**Maximum Sum of Distinct Subarrays With Length K** (Medium)
🔗 https://leetcode.com/problems/maximum-sum-of-distinct-subarrays-with-length-k/

Day 92's running sum + Day 93's count map, on one fixed window of length k.

```go
m[nums[i]]++
sum += int64(nums[i])
if len(m) == k+1 { // k was decremented
	result = maxVal(result, sum)
}
m[nums[i-k]]--
if m[nums[i-k]] == 0 {
	delete(m, nums[i-k])
}
sum -= int64(nums[i-k])
```

Deleting at zero makes `len(m)` the distinct count, so "all distinct" is one comparison. `int64` because sums can hit 10^10.

`[1,5,4,2,9,9,9]`, k = 3 → 15.

O(n) time, O(k) space.

Full walkthrough with step-by-step diagrams: https://github.com/architagr/leetcode_solutions/blob/main/medium_problems/2401_2500/maximum_sum_of_distinct_subarrays_with_length_k/SOLUTION.md
