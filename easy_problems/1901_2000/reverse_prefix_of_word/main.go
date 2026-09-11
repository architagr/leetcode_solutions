package reverseprefixofword

func reversePrefix(word string, ch byte) string {
	b := []byte(word)
	i := 0
	for ; i < len(b); i++ {
		if b[i] == ch {
			break
		}
	}
	if i < len(b) {
		for l := 0; l <= i; l, i = l+1, i-1 {
			b[i], b[l] = b[l], b[i]
		}
	}
	return string(b)
}
