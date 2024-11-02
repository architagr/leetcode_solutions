package circularsentence

func isCircularSentence(sentence string) bool {
	start, end := 0, len(sentence)-1

	for ; start < end; start++ {
		if sentence[start] != ' ' {
			break
		}
	}

	for ; end >= start; end-- {
		if sentence[end] != ' ' {
			break
		}
	}
	if sentence[end] != sentence[start] {
		return false
	}

	for ; start < end; start++ {
		if sentence[start] == ' ' {
			if sentence[start-1] != sentence[start+1] {
				return false
			}
		}
	}
	return true
}
