package decode_the_message

func DecodeMessage(key, message string) string {
	ans := make([]byte, len(message))

	// get map table of the key
	mapKey := make(map[byte]byte)
	x := byte('a')
	for i := 0; i < len(key); i++ {
		_, ok := mapKey[key[i]]
		if !ok && key[i] != ' ' {
			mapKey[key[i]] = x
			x++
		}
	}
	//decode
	for i := 0; i < len(message); i++ {
		if message[i] == ' ' {
			ans[i] = message[i]
		} else {
			ans[i] = mapKey[message[i]]
		}
	}
	return string(ans)
}
