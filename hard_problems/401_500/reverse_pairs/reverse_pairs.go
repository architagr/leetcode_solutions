package reverse_pairs

func ReversePairs(nums []int) int {
	return mergeSort(&nums, 0, len(nums)-1)
}
func mergeSort(A *[]int, start, end int) int {
	if start == end {
		return 0
	}
	mid := (start + end) / 2
	a := mergeSort(A, start, mid)
	b := mergeSort(A, mid+1, end)
	ab := merge(A, start, mid, end)
	return (a + b + ab)
}
func merge(A *[]int, l, y, r int) int {
	n, m := y-l+1, r-y
	B := make([]int, n)
	C := make([]int, m)
	for i := l; i <= y; i++ {
		B[i-l] = (*A)[i]
	}
	for i := y + 1; i <= r; i++ {
		C[i-(y+1)] = (*A)[i]
	}
	i, j, x, y, c := 0, 0, 0, 0, 0
	for k := l; k <= r; k++ {
		if i == n {
			(*A)[k] = C[j]
			j += 1
		} else if j == m {
			(*A)[k] = B[i]
			i += 1
		} else if B[i] <= C[j] {
			(*A)[k] = B[i]
			i += 1
		} else {
			(*A)[k] = C[j]
			j += 1
		}
		if x < n && y < m && B[x] > 2*C[y] {
			c += (n - x)
			y += 1
		} else {
			x++
		}
	}
	return c
}
