package binarysubarrayswithsum

func numSubarraysWithSum(nums []int, goal int) int {
	mapData := make(map[int]int)
	sum := 0
	count := 0
	for _, x := range nums {
		sum += x
		// (j, i] sums to goal when prefix[j] == sum - goal. The empty prefix
		// (j before index 0) is counted here instead of seeding mapData[0].
		if sum == goal {
			count++
		}
		// Every earlier prefix equal to sum-goal starts a subarray ending here.
		if sum >= goal {
			count += mapData[sum-goal]
		}
		mapData[sum]++
	}
	return count
}
