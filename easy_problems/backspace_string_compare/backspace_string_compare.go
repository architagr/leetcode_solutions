package backspace_string_compare

func BackspaceCompare(s, t string) bool {
	sA := getByteArr(s)
	tA := getByteArr(t)

	return string(sA) == string(tA)
}
func getByteArr(s string) []byte {
	arr := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '#' && len(arr) > 0 {
			arr = arr[:len(arr)-1]
		}
		if s[i] != '#' {
			arr = append(arr, s[i])
		}
	}
	return arr
}
