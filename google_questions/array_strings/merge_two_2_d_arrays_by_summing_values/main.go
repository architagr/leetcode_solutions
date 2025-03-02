package mergetwo2darraysbysummingvalues

const (
	idIndex    = 0
	valueIndex = 1
)

func mergeArrays(nums1, nums2 [][]int) [][]int {

	n, m := len(nums1), len(nums2)
	result := make([][]int, 0, n+m)
	for i, j := 0, 0; i < n || j < m; {
		if j >= m || (i < n && nums1[i][idIndex] < nums2[j][idIndex]) {
			result = append(result, []int{nums1[i][idIndex], nums1[i][valueIndex]})
			i++
		} else if i >= n || (j < m && nums1[i][idIndex] > nums2[j][idIndex]) {
			result = append(result, []int{nums2[j][idIndex], nums2[j][valueIndex]})
			j++
		} else if i < n && j < m && nums1[i][idIndex] == nums2[j][idIndex] {
			result = append(result, []int{nums2[j][idIndex], nums1[i][valueIndex] + nums2[j][valueIndex]})
			i++
			j++
		}
	}

	return result
}
