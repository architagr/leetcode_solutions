package rearrange_array_elements_by_sign

func RearrangeArray(nums []int) []int {
	n := len(nums)
	positive := make([]int, 0, n/2)
	negitive := make([]int, 0, n/2)
	k, pk, nk := 0, 0, 0

	for i := 0; i < n; i++ {
		if nums[i] >= 0 {
			positive = append(positive, nums[i])
		} else {
			negitive = append(negitive, nums[i])
		}

		if k%2 == 0 {
			if len(positive) > pk {
				nums[k] = positive[pk]
				pk++
				k++
			}
		} else {
			if len(negitive) > nk {
				nums[k] = negitive[nk]
				nk++
				k++
			}
		}
	}
	for ; k < n; k++ {
		if k%2 == 0 {
			nums[k] = positive[pk]
			pk++
		} else {
			nums[k] = negitive[nk]
			nk++
		}
	}

	return nums
}
