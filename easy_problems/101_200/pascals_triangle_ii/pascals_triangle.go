package pascalstriangle2

func GeneratePascaleTriangle(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	} else if rowIndex == 1 {
		return []int{1, 1}
	}
	arr := []int{1, 1}
	return genrateTriangle(&arr, rowIndex)
}

func genrateTriangle(arr *[]int, num int) []int {
	n := len(*arr)
	if n-1 == num {
		return *arr
	}

	row := make([]int, n+1)

	row[0] = 1
	row[n] = 1
	for i := 1; i < n; i++ {
		row[i] = (*arr)[i-1] + (*arr)[i]
	}

	return genrateTriangle(&row, num)
}
