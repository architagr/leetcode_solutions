package excelsheetcolumntitle

var (
	m = map[int]byte{
		1:  'A',
		2:  'B',
		3:  'C',
		4:  'D',
		5:  'E',
		6:  'F',
		7:  'G',
		8:  'H',
		9:  'I',
		10: 'J',
		11: 'K',
		12: 'L',
		13: 'M',
		14: 'N',
		15: 'O',
		16: 'P',
		17: 'Q',
		18: 'R',
		19: 'S',
		20: 'T',
		21: 'U',
		22: 'V',
		23: 'W',
		24: 'X',
		25: 'Y',
		26: 'Z',
	}
)

func convertToTitle(columnNumber int) string {

	sb := make([]byte, 0, 26)
	for columnNumber > 0 {
		rem := columnNumber % 26
		if rem == 0 {
			rem = 26
			columnNumber -= 1
		}
		sb = append(sb, m[rem])
		columnNumber /= 26
	}
	return string(reverseBytes(sb))
}

func reverseBytes(b []byte) []byte {
	n := len(b)
	for i := 0; i < n/2; i++ {
		b[i], b[n-1-i] = b[n-1-i], b[i]
	}
	return b
}
