package is_subsequence

func IsSubsequence(s string, t string) bool {
	k, n, m := 0, len(t), len(s)
	if m == 0 {
		return true
	}
	for i := 0; i < n; i++ {
		if s[k] == t[i] {
			k++
			if k == m {
				return true
			}
		}
	}
	return false
}
