package decodeways

import "strconv"

func numDecodings(s string) int {
	dpTable := make(map[int]int, len(s))
	return dfs(0, s, dpTable)
}

func dfs(index int, str string, dpTable map[int]int) int {
	if val, ok := dpTable[index]; ok {
		return val
	}
	if index == len(str) {
		return 1
	}
	if str[index] == '0' {
		return 0
	}
	if index == len(str)-1 {
		return 1
	}
	ans := dfs(index+1, str, dpTable)
	substrNum, _ := strconv.Atoi(str[index : index+2])
	if substrNum <= 26 {
		ans += dfs(index+2, str, dpTable)
	}
	dpTable[index] = ans
	return ans
}
