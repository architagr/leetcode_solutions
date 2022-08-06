package string_to_integer_atoi

func MyAtoi(s string) int {
	maxInt := 2147483647
	minInt := -2147483648
	var result int
	var multiplier int
	length := len(s)

	var startFrom int

	for ; startFrom < length; startFrom++ {
		if s[startFrom] != ' ' {
			break
		}
	}

	s = s[startFrom:]
	length = len(s)

	if length == 0 {
		return result
	}

	if (s[0] < '0') || (s[0] > '9') {
		if s[0] == '-' {
			multiplier = -1
		} else if s[0] == '+' {
			multiplier = 1
		} else {
			return result
		}
	} else {
		result = int(s[0] - '0')
		multiplier = 1
	}

	for i := 1; i < length; i++ {
		if (s[i] < '0') || (s[i] > '9') {
			break
		}

		result = (result * 10) + int(s[i]-'0')

		if result >= maxInt {
			break
		}
	}

	result *= multiplier

	if result >= maxInt {
		return maxInt
	}

	if result <= minInt {
		return minInt
	}

	return result
}
