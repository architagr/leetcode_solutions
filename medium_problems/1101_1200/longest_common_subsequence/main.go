package longestcommonsubsequesnce

func longestCommonSubsequence(text1 string, text2 string) int {

	if len(text1) > len(text2) {
		text1, text2 = text2, text1
	}
	prev, current := make([]int, len(text1)+1), make([]int, len(text1)+1)
	for i := len(text2) - 1; i >= 0; i-- {
		for j := len(text1) - 1; j >= 0; j-- {
			if text1[j] == text2[i] {
				current[j] = 1 + prev[j+1]
			} else {
				current[j] = maxValue(prev[j], current[j+1])
			}
		}
		current, prev = prev, current
	}

	return prev[0]
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
