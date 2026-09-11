package palindrome_number

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	} else if x == 0 {
		return true
	} else if x%10 == 0 {
		return false
	}

	//start approx 1
	// divisor, length := getFirstDivisor(x)

	// for i := 0; i < length/2; i++ {
	// 	rightnumber := (x / divisor) % 10
	// 	leftNumber := x % 10
	// 	x %= divisor
	// 	x /= 10
	// 	divisor /= 100

	// 	if rightnumber != leftNumber {
	// 		return false
	// 	}

	// }
	// return true
	//end approch 1

	// start approch 2

	reserveNum := 0
	for x > reserveNum {
		reserveNum = reserveNum*10 + x%10
		x /= 10
	}
	return x == reserveNum || x == reserveNum/10
	//end approch 2
}

func getFirstDivisor(x int) (int, int) {
	divisor, length := 1, 1
	for x >= 10 {
		divisor *= 10
		x /= 10
		length++
	}

	return divisor, length
}
