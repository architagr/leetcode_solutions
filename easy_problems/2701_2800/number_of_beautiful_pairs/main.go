package numberofbeautifulpairs

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func countBeautifulPairs(nums []int) int {
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		FirstDigit := nums[i]
		for FirstDigit >= 10 { //divide by ten to get the first digit
			FirstDigit /= 10
		}
		for j := i + 1; j < len(nums); j++ {
			LastDigit := nums[j] % 10            // to get the last digit
			if gcd(FirstDigit, LastDigit) == 1 { // to identify if gcd==1
				count++
			}
		}
	}
	return count
}
