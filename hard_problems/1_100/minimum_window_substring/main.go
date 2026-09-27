package minimumwindowsubstring

func minWindow(s string, t string) string {
	// What the window still needs from t. Entering characters decrement,
	// leaving ones increment; the window covers t when nothing is above 0.
	// Negative means a surplus the window can afford to lose.
	frequencyMapT := make(map[byte]int, 26)
	for i := 0; i < len(t); i++ {
		frequencyMapT[t[i]]++
	}
	result := ""

	for l, r := 0, 0; r < len(s); r++ {
		if _, ok := frequencyMapT[s[r]]; !ok {
			continue
		}
		frequencyMapT[s[r]]--
		// Covered: shrink from the left as long as it stays covered, keeping
		// the shortest window. The continue below still runs l++.
		for ; found(frequencyMapT) && l <= r; l++ {
			if len(result) == 0 || r-l+1 < len(result) {
				result = s[l : r+1]
			}
			if _, ok := frequencyMapT[s[l]]; !ok {
				continue
			}
			frequencyMapT[s[l]]++
		}
	}
	return result
}
// found reports whether nothing is still needed. It scans the map (at most
// 52 letters); a "missing" counter updated as entries cross 0 would make it O(1).
func found(frequencyMapT map[byte]int) bool {
	for _, c := range frequencyMapT {
		if c > 0 {
			return false
		}
	}
	return true
}
