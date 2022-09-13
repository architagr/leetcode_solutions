package utf_8_validation

func ValidateUTF8(data []int) bool {
	n, t := len(data), 0

	for i := 0; i < n; {
		if data[i] >= 0 && data[i] <= 127 {
			t = 0
		} else if data[i] >= 192 && data[i] <= 223 {
			t = 1
		} else if data[i] >= 224 && data[i] <= 239 {
			t = 2
		} else if data[i] >= 240 && data[i] <= 247 {
			t = 3
		} else {
			return false
		}
		i++

		for i < n && t > 0 {
			t--
			if data[i] < 128 || data[i] > 191 {
				return false
			}
			i++
		}
	}
	return t == 0
}
