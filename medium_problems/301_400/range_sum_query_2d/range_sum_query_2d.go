package range_sum_query_2d

type NumMatrix struct {
	maxtrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	sumData := make([][]int, len(matrix))
	for ri, row := range matrix {
		rowSum := 0
		sumData[ri] = make([]int, len(row))
		for ci, n := range row {
			rowSum += n
			preRowSum := 0
			if ri != 0 {
				preRowSum = sumData[ri-1][ci]
			}

			sumData[ri][ci] = rowSum + preRowSum
		}
	}
	return NumMatrix{
		maxtrix: sumData,
	}
}

func (obj *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := obj.maxtrix[row2][col2]
	if row1 > 0 {
		sum -= obj.maxtrix[row1-1][col2]
	}
	if col1 > 0 {
		sum -= obj.maxtrix[row2][col1-1]
	}
	if row1 > 0 && col1 > 0 {
		sum += obj.maxtrix[row1-1][col1-1]
	}
	return sum
}
