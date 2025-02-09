package lettercombinationsofaphonenumber

var numberPad = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	byteArr := make([]byte, len(digits))
	result := make([]string, 0)
	combination(0, digits, byteArr, &result)
	return result
}
func combination(index int, digits string, byteArr []byte, result *[]string) {
	if index == len(digits) {
		*result = append(*result, string(byteArr))
		return
	}
	for _, c := range numberPad[digits[index]] {
		byteArr[index] = c
		combination(index+1, digits, byteArr, result)
	}
}
