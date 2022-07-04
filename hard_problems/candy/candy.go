package candy

const (
	Equal      = 0
	Increasing = 1
	Decreasing = -1
)

func Candy(ratings []int) int {
	n := len(ratings)
	if n <= 1 {
		return n
	}
	candies, up, down, oldSlope := 0, 0, 0, 0
	for i := 1; i < n; i++ {
		newSlope := Equal
		if ratings[i] > ratings[i-1] {
			newSlope = Increasing
		} else if ratings[i] < ratings[i-1] {
			newSlope = Decreasing
		}
		if (oldSlope == Increasing && newSlope == Equal) || (oldSlope == Decreasing && (newSlope == Increasing || newSlope == Equal)) {
			candies += sumOfN(up) + sumOfN(down) + maxValue(up, down)
			up = 0
			down = 0
		}
		if newSlope == Increasing {
			up++
		} else if newSlope == Decreasing {
			down++
		} else {
			candies++
		}

		oldSlope = newSlope
	}
	candies += sumOfN(up) + sumOfN(down) + maxValue(up, down) + 1
	return candies
}

func sumOfN(n int) int {
	return n * (n + 1) / 2
}
func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
