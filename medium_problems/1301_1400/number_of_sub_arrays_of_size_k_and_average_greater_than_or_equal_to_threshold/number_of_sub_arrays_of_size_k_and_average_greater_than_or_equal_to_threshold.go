package numberofsubarraysofsizekandaveragegreaterthanorequaltothreshold

func numOfSubarrays(arr []int, k int, threshold int) int {
	threshold *= k
	k--
	sum, count := 0, 0

	for i := 0; i < k; i++ {
		sum += arr[i]
	}
	for i := k; i < len(arr); i++ {
		sum += arr[i]
		if sum >= threshold {
			count++
		}
		sum -= arr[i-k]
	}
	return count
}
