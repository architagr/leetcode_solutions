package maximumdifferencebetweenincreasingelements

func maximumDifference(nums []int) int {
	min, diff := nums[0], -1
	for i := 1; i < len(nums); i++ {
		if nums[i] <= min {
			min = nums[i]
			continue
		}
		d := nums[i] - min
		if d > diff {
			diff = d
		}
	}

	return diff
}
