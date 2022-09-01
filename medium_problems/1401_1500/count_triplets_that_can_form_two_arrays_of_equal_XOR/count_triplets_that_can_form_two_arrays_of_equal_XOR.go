package count_triplets_that_can_form_two_arrays_of_equal_XOR

func CountTriplets(arr []int) int {
	var ans int

	seen := map[int][]int{0: {-1}}

	for k, x := 0, 0; k < len(arr); k++ {
		x ^= arr[k]
		for _, i := range seen[x] {
			ans += k - i - 1
		}
		seen[x] = append(seen[x], k)
	}

	return ans
}
