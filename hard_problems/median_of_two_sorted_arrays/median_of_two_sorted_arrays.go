package median_of_two_sorted_arrays

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	arr := mergeArray(nums1, nums2)

	n := len(arr)

	if n%2 == 0 {
		return (float64(arr[n/2]) + float64(arr[(n-2)/2])) / 2.0
	}
	return float64(arr[n/2])
}

func mergeArray(nums1 []int, nums2 []int) []int {
	n, m, k := len(nums1), len(nums2), 0
	arr := make([]int, n+m)
	i, j := 0, 0
	for i < n && j < m {
		if nums1[i] < nums2[j] {
			arr[k] = nums1[i]
			i++
			k++
		} else if nums1[i] > nums2[j] {
			arr[k] = nums2[j]
			j++
			k++
		} else if nums1[i] == nums2[j] {
			arr[k] = nums1[i]
			k++
			arr[k] = nums2[j]
			k++
			j++
			i++
		}
	}
	if i < len(nums1) {
		for ; i < len(nums1); i, k = i+1, k+1 {
			arr[k] = nums1[i]
		}
	}
	if j < len(nums2) {
		for ; j < len(nums2); j, k = j+1, k+1 {
			arr[k] = nums2[j]
		}
	}

	return arr
}
