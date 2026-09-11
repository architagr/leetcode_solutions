package add_binary

import "strings"

func AddBinary(a, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	lenA, lenB := len(a), len(b)

	reminder := 0
	result := make([]byte, 0, lenA+1)

	for i := 0; i < lenA; i++ {
		if a[lenA-1-i] == '1' {
			reminder++
		}

		if lenB-1-i >= 0 && b[lenB-1-i] == '1' {
			reminder++
		}

		switch reminder {
		case 0:
			result = append(result, '0')
		case 1:
			result = append(result, '1')
			reminder = 0
		case 2:
			result = append(result, '0')
			reminder = 1
		case 3:
			result = append(result, '1')
			reminder = 1
		}
	}

	if reminder > 0 {
		result = append(result, '1')
	}

	var sb strings.Builder
	for i := len(result) - 1; i >= 0; i-- {
		sb.WriteByte(result[i])
	}

	return sb.String()
}

func AddBinaryApproch1(a, b string) string {

	n, m := len(a)-1, len(b)-1
	c := byte('0')
	result := ""
	for n >= 0 && m >= 0 {
		result = string((c^a[n])^b[m]) + result

		if a[n] == '1' && b[m] == '1' || a[n] == '1' && c == '1' || c == '1' && b[m] == '1' {
			c = byte('1')
		} else {
			c = byte('0')
		}
		n--
		m--
	}
	for n >= 0 {
		result = string((c^a[n])^'0') + result

		if a[n] == '1' && c == '1' {
			c = byte('1')
		} else {
			c = byte('0')
		}
		n--
	}
	for m >= 0 {
		result = string((c^'0')^b[m]) + result

		if c == '1' && b[m] == '1' {
			c = byte('1')
		} else {
			c = byte('0')
		}
		m--
	}
	if c == '1' {
		result = string(c) + result
	}
	return result
}
