package majority_element

func MajorityElement(A []int) int {
	me := A[0]
	count := 1
	n := len(A)
	for i := 1; i < n; i++ {
		if count == 0 {
			me = A[i]
			count = 1
		} else {
			if me == A[i] {
				count++
			} else {
				count--
			}
		}
	}
	count = getCount(A, me)
	if count > (len(A) / 2) {
		return me
	} else {
		return -1
	}
}

func getCount(A []int, nm int) int {
	count := 0
	for i := 0; i < len(A); i++ {
		if A[i] == nm {
			count++
		}
	}
	return count
}
