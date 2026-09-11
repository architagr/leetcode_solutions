package mergesortedarray

func merge(nums1 []int, m int, nums2 []int, n int) {
	y := 0
	if m == 0 {
		copy(nums1, nums2)
	} else if n == 0 {
		return
	} else {
		for x := 0; x < n; x++ {
			y := findInsertPos(nums1[y:m+x], nums2[x])
			if y < m+x {
				copy(nums1[y+1:], nums1[y:])
			}
			nums1[y] = nums2[x]
		}
	}
}
func findInsertPos(num1 []int, val int) int {
	len := len(num1)
	mid := len / 2
	if num1[0] >= val {
		return 0
	} else if num1[len-1] <= val {
		return len
	} else if num1[mid] > val {
		return findInsertPos(num1[0:mid], val)
	} else {
		return mid + findInsertPos(num1[mid:len], val)
	}
}
