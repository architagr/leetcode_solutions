package string_to_integer_atoi

import "strings"

func MyAtoi(s string) int {
	hasNegitive := false
	s = strings.TrimSpace(s)
	num := int64(0)
	lastSign := false
	for i := 0; i < len(s); i++ {
		a, ok := validChar(s[i])
		if ok && a < 0 && num >= 0 && i > 0 && s[i-1] != ' ' {
			break
		}
		if ok && a == -1 {
			hasNegitive = true
			if lastSign {
				break
			}
			lastSign = true
		} else if ok && a == -2 {
			hasNegitive = false
			if lastSign {
				break
			}
			lastSign = true
		} else if ok && a >= 0 {
			lastSign = false
			num = (num * 10) + int64(a)
			if hasNegitive && num >= (1<<31) {
				return -(1 << 31)
			}

			if num > ((1 << 31) - 1) {
				return ((1 << 31) - 1)
			}
		} else if !ok {
			break
		}
	}
	if hasNegitive {
		num = num * -1
	}

	return int(num)
}

func validChar(b byte) (int, bool) {
	if b >= '0' && b <= '9' {
		return int(b - '0'), true
	} else if b == '-' {
		return -1, true
	} else if b == '+' {
		return -2, true
	}

	return 0, false

}
