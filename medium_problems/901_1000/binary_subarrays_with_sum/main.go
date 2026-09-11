package binarysubarrayswithsum

func numSubarraysWithSum(nums []int, goal int) int {
	mapData := make(map[int]int)
	sum := 0
	count := 0
	for _, x := range nums {
		sum += x
		if sum == goal {
			count++
		}
		if sum >= goal {
			count += mapData[sum-goal]
		}
		mapData[sum]++
	}
	return count
}
