package intersectionoftwoarray

func intersection(nums1 []int, nums2 []int) []int {
	map1 := make(map[int]bool, len(nums1))
	map2 := make(map[int]bool, len(nums2))
	for _, n := range nums1 {
		map1[n] = true
	}
	for _, n := range nums2 {
		map2[n] = true
	}
	result := make([]int, 0, len(map1))
	for k := range map1 {
		if _, ok := map2[k]; ok {
			result = append(result, k)
		}
	}
	return result
}
