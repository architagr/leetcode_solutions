package numberofsubarraysofsizekandaveragegreaterthanorequaltothreshold

func numOfSubarrays(arr []int, k int, threshold int) int {
	// average >= threshold  is  sum/k >= threshold  is  sum >= threshold*k.
	// Scaling once here keeps every window check an integer comparison.
	threshold *= k
	// k becomes an offset: arr[i-k] is the oldest element in the window.
	k--
	sum, count := 0, 0

	// Prime one short, so the loop is add, check, remove every time.
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
