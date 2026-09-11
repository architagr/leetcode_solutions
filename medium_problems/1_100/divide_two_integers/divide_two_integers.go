package divide_two_integers

import "math"

func Divide(A int, B int) int {
	if (A == math.MaxInt32 && B == 1) || (A == -(1<<31) && B == -1) {
		return math.MaxInt32
	}
	sign := false

	if (A >= 0 && B >= 0) || (A < 0 && B < 0) {
		sign = true
	}
	A = int(math.Abs(float64(A)))
	B = int(math.Abs(float64(B)))
	ans := 0
	for A-B >= 0 {
		count := 0
		for A-(B<<1<<count) >= 0 {
			count++
		}
		ans += 1 << count
		A -= B << count
	}
	if !sign {
		ans *= -1
	}
	return ans
}
