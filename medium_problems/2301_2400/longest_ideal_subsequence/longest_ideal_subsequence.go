package longest_ideal_subsequence

// LongestIdealString returns the length of the longest subsequence of s
// whose adjacent letters sit within k of each other in the alphabet.
//
// Only the last letter of a candidate subsequence constrains what can
// follow it, so tracking the best length per ending letter is enough:
// best[c] is the longest ideal subsequence ending in letter c.
func LongestIdealString(s string, k int) int {
	var best [26]int
	answer := 0

	for i := 0; i < len(s); i++ {
		c := int(s[i] - 'a')
		longest := 0
		lo, hi := c-k, c+k
		if lo < 0 {
			lo = 0
		}
		if hi > 25 {
			hi = 25
		}
		for j := lo; j <= hi; j++ {
			if best[j] > longest {
				longest = best[j]
			}
		}
		if longest+1 > best[c] {
			best[c] = longest + 1
		}
		if best[c] > answer {
			answer = best[c]
		}
	}
	return answer
}
