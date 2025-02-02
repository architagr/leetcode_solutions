package longpressedname

func isLongPressedName(name string, typed string) bool {
	if len(typed) < len(name) {
		return false
	}
	n, m := len(name), len(typed)
	if name[n-1] != typed[m-1] {
		return false
	}
	i, j := 0, 0
	for i < len(name) && j < len(typed) {
		if name[i] != typed[j] {
			return false
		}
		countName := countConsecutiveChar(name, i)
		countTyped := countConsecutiveChar(typed, j)
		if countName > countTyped {
			return false
		}
		i += countName
		j += countTyped
	}
	return j >= len(typed) && i >= len(name)
}
func countConsecutiveChar(str string, start int) int {
	count := 1
	c := str[start]
	for start = start + 1; start < len(str); start, count = start+1, count+1 {
		if str[start] != c {
			break
		}
	}
	return count
}
