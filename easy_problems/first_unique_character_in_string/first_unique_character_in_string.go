package first_unique_character_in_string

func firstUniqChar(s string) int {
	type countPos struct {
		count int
		pos   int
	}
	var counts [26]countPos
	for i, r := range s {
		counts[r-'a'].count++
		counts[r-'a'].pos = i
	}

	pos := len(s)
	for _, l := range counts {
		if l.count == 1 && l.pos < pos {
			pos = l.pos
		}
	}
	if pos == len(s) {
		return -1
	}
	return pos
}
