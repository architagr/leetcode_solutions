package addstrings

func addStrings(num1, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	carry := 0

	result := make([]byte, 0, max(len(num1), len(num2))+1)

	for i >= 0 || j >= 0 || carry > 0 {
		d1 := getDigit(num1, i)
		d2 := getDigit(num2, j)

		sum := d1 + d2 + carry
		result = append(result, byte(sum%10)+'0')
		carry = sum / 10

		i--
		j--
	}

	reverseBytes(result)
	return string(result)
}

func getDigit(num string, idx int) int {
	if idx < 0 {
		return 0
	}
	return int(num[idx] - '0')
}

func reverseBytes(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
