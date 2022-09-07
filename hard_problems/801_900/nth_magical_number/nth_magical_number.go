package nth_magical_number

func NthMagicalNumber(A int, B int, C int) int {
	MOD := int64(1000000007)
	min := B

	if B > C {
		min = C
	}
	//Ath element can not go beyong Ath multiple of Minimum of B,C.
	//low would be the first multiple of Minimum of B,C.
	high, low := int64(A)*int64(min), int64(min)

	//To calculate the Ath position divisble by B or C, we need to (imagine) think of the A multiples of B and C atleast(say 5)
	// But we don't need to calculate the common multiple twice.
	//So we need to calculate lcm of B & C, to be able to reduce the common multiple of B and C.
	lcm := int64((B * C) / gcd(B, C))
	ans := low
	if A == 1 {
		return int(ans % MOD)
	}

	//Now we are ready to run Binary Search for the search space defined above.
	for low <= high {
		mid := (low + high) / 2
		countB := mid / int64(B)
		countC := mid / int64(C)
		//To Reduce count of common factors.
		commonFactors := mid / lcm
		position := (countB + countC - commonFactors)
		//As we need the Ath Smallest,
		if position >= int64(A) {
			ans = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if ans < 0 {
		ans += MOD
	}
	return int(ans % MOD)
}
func gcd(a, b int) int {
	if b > a {
		a, b = b, a
	}
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
