package pointsthatintersectwithcars

func numberOfPoints(nums [][]int) int {
	line := make([]int, 102)
	for _, val := range nums {
		line[val[0]]++
		line[val[1]+1]--
	}
	var counter int
	var result int
	for i := 0; i < 102; i++ {
		counter += line[i]
		if counter > 0 {
			result++
		}
	}
	return result
}
