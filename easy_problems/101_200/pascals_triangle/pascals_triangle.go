package pascalstriangle

func GeneratePascaleTriangle(numRows int) [][]int {
	arr := make([][]int, 0)
	if numRows == 1 {
		arr = append(arr, []int{1})
	} else {
		arr = append(arr, []int{1})
		arr = append(arr, []int{1, 1})
		genrateTriangle(&arr, numRows)
	}
	return arr
}

func genrateTriangle(arr *[][]int, num int) {
	n := len(*arr)
	if n == num {
		return
	}

	row := make([]int, n+1)

	row[0] = 1
	row[n] = 1
	for i := 1; i < n; i++ {
		row[i] = (*arr)[n-1][i-1] + (*arr)[n-1][i]
	}

	*arr = append((*arr), row)
	genrateTriangle(arr, num)
}
