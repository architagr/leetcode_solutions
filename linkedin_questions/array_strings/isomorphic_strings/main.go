package isomorphicstrings

import "strings"

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	ms := make(map[rune]int, 26)
	mt := make(map[rune]int, 26)
	for _, r := range s {
		ms[r]++
	}
	for _, r := range t {
		mt[r]++
	}
	if len(ms) != len(mt) {
		return false
	}
	sb := strings.Builder{}
	sb.Grow(len(s))
	replace := make(map[rune]rune)
	for i, r := range s {
		if rt, ok := replace[r]; ok {
			if rt == rune(t[i]) {
				sb.WriteRune(rt)
			} else {
				return false
			}
			continue
		}
		if mt[rune(t[i])] == ms[r] {
			replace[r] = rune(t[i])
			sb.WriteRune(replace[r])
		} else {
			return false
		}
	}
	return t == sb.String()
}
