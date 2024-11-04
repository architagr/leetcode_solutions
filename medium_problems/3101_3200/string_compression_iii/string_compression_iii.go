package stringcompressioniii

func compressedString(word string) string {
	c := word[0]
	output := make([]byte, 0, 2*len(word))
	count := 1
	for i := 1; i < len(word); i++ {
		for ; i < len(word) && c == word[i]; i++ {
			count++
		}
		for count > 9 {
			output = append(output, '9', c)
			count -= 9
		}
		output = append(output, getByte(count), c)

		count = 0
		if i < len(word) {
			c = word[i]
			count = 1
		}
	}
	if count > 0 {
		output = append(output, getByte(count), c)
	}
	return string(output)
}
func getByte(i int) byte {
	switch i {
	case 1:
		return byte('1')
	case 2:
		return byte('2')
	case 3:
		return byte('3')
	case 4:
		return byte('4')
	case 5:
		return byte('5')
	case 6:
		return byte('6')
	case 7:
		return byte('7')
	case 8:
		return byte('8')
	}
	return byte('9')
}
