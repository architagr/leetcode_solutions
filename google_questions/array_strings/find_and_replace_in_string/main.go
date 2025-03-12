package findandreplaceinstring

import "strings"

func findReplaceString(s string, indices []int, sources, targets []string) string {
	n := len(s)
	match := make([]int, n)
	for i := 0; i < n; i++ {
		match[i] = -1
	}

	for i := 0; i < len(indices); i++ {
		ix := indices[i]
		end := ix + len(sources[i])
		if end <= len(s) && s[ix:end] == sources[i] {
			match[ix] = i
		}

	}

	ans := strings.Builder{}
	ix := 0
	for ix < n {
		if match[ix] >= 0 {
			ans.WriteString(targets[match[ix]])
			ix += len(sources[match[ix]])
		} else {
			ans.WriteByte(s[ix])
			ix++
		}
	}
	return ans.String()
}
