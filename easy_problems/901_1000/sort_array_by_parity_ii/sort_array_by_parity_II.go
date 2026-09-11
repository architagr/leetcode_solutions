package sort_array_by_parity_II

func SortArrayByParityII(nums []int) []int {
	n := len(nums)
	even := make([]int, 0, n/2)
	odd := make([]int, 0, n/2)
	k, ek, ok := 0, 0, 0

	for i := 0; i < n; i++ {
		if nums[i]%2 == 0 {
			even = append(even, nums[i])
		} else {
			odd = append(odd, nums[i])
		}

		if k%2 == 0 {
			if len(even) > ek {
				nums[k] = even[ek]
				ek++
				k++
			}
		} else {
			if len(odd) > ok {
				nums[k] = odd[ok]
				ok++
				k++
			}
		}
	}
	for ; k < n; k++ {
		if k%2 == 0 {
			nums[k] = even[ek]
			ek++
		} else {
			nums[k] = odd[ok]
			ok++
		}
	}

	return nums
}
