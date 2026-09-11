package minimum_XOR_sum_of_two_arrays

import "sort"

func MinimumXORSum(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	nums1Map := make(map[int][]int)
	nums2Map := make(map[int][]int)

	for i := 0; i < len(nums1); i++ {
		val, ok := nums1Map[nums1[i]]

		if ok {
			val = append(val, i)
		} else {
			val = []int{i}
		}
		nums1Map[nums1[i]] = val
	}

	for i := 0; i < len(nums2); {
		val, ok := nums2Map[nums2[i]]
		val1, ok1 := nums1Map[nums2[i]]
		if ok1 {
			if len(val1) == 1 {
				delete(nums1Map, nums2[i])
			} else {
				val1 = val1[1:]
				nums1Map[nums2[i]] = val1
			}

			nums1 = append(nums1[:val1[0]], nums1[val1[0]+1:]...)
			if i+1 == len(nums2) {
				nums2 = nums2[:i]
			} else {
				nums2 = append(nums2[:i], nums2[i+1:]...)
			}

		} else {
			if ok {
				val = append(val, i)
			} else {
				val = []int{i}
			}
			nums2Map[nums2[i]] = val
			i++
		}
	}
	ans := 0

	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i]%2 == 0 && nums2[j]%2 == 0 {
			ans += nums1[i] ^ nums2[j]
			nums1 = append(nums1[:i], nums1[i+1:]...)
			nums2 = append(nums2[:j], nums2[j+1:]...)
		} else if nums1[i]%2 != 0 {
			i++
		} else if nums2[j]%2 != 0 {
			j++
		}
	}

	for i := 0; i < len(nums1); i++ {
		ans += nums1[i] ^ nums2[i]
	}

	return ans
}
