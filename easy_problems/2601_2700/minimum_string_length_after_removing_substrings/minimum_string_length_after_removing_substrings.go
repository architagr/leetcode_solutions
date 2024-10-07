package minimumstringlengthafterremovingsubstrings

// [0,1,2,3,4,5,6,7]
// [C,C,C,C,D,D,D,D]
func minLength(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	isRemoved := false
	for i := 0; i < len(s)-1; i++ {
		if i < 0 {
			continue
		}
		if s, isRemoved = remove(s, i); isRemoved {
			i -= 2
		}
	}
	if len(s) >= 2 {
		s, isRemoved = remove(s, len(s)-2)
		for isRemoved {
			s, isRemoved = remove(s, len(s)-2)
		}
	}
	return len(s)
}

func remove(str string, i int) (string, bool) {
	s := string(str[i : i+2])
	if s == "AB" || s == "CD" {
		str = str[:i] + str[i+2:]
		return str, true
	}
	return str, false
}
