package pow_x_n

func MyPow(x float64, n int) float64 {

	isPowerNegitive := false
	if n < 0 {
		isPowerNegitive = true
		n *= -1
	}
	val := pow(x, n)
	if isPowerNegitive {
		return 1 / val
	}
	return val
}

func pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n&1 == 0 {
		return pow(x*x, n/2)
	}
	return pow(x*x, n/2) * x
}
