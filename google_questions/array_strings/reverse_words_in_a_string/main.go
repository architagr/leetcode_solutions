package reversewordsinastring

import "strings"

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	b := make([]byte, 0, len(s))
	start, end := 0, 0

	for i := len(s) - 1; i >= 0; i-- {
		if len(b) > 0 && s[i] == ' ' && b[end-1] == ' ' {
			continue
		}
		b = append(b, s[i])

		if s[i] == ' ' {
			reverse(b, start, end-1)
			start = end + 1
		}
		end++

	}
	reverse(b, start, end-1)
	return string(b)

}

func reverse(str []byte, start, end int) []byte {
	for start <= end {
		str[start], str[end] = str[end], str[start]
		start++
		end--
	}

	return str
}
