package calculate_digit_sum_of_a_string

import "strconv"

func DigitSum(s string, k int) string {
	if len(s) <= k {
		return s
	}
	news := ""
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += int(s[i] - '0')
		if (i+1)%k == 0 || i+1 == len(s) {
			news += strconv.Itoa(sum)
			sum = 0
		}
	}
	return DigitSum(news, k)
}
