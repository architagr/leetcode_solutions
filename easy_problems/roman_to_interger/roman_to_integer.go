package roman_to_integer

var ref map[byte]int = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {

	output, length := 0, len(s)
	last := 0
	for i := length - 1; i >= 0; i-- {
		r := ref[s[i]]
		if r >= last {
			output += r
			last = r
		} else {
			output -= r
		}
	}
	return output
}
