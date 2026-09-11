package add_digits

func AddDigits(num int) int {
	if num <= 9 {
		return num
	}

	sum := 0

	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return AddDigits(sum)
}
