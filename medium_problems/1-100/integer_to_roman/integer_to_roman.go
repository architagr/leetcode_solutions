package integer_to_roman

type IntegerRoman struct {
	integer int
	symbol  string
}

var intRomanArray = []IntegerRoman{
	{
		integer: 1,
		symbol:  "I",
	},
	{
		integer: 5,
		symbol:  "V",
	},
	{
		integer: 10,
		symbol:  "X",
	},
	{
		integer: 50,
		symbol:  "L",
	},
	{
		integer: 100,
		symbol:  "C",
	},
	{
		integer: 500,
		symbol:  "D",
	},
	{
		integer: 1000,
		symbol:  "M",
	},
}

func IntToRoman(num int) string {
	if num == 0 {
		return ""
	}
	ans := ""
	for i := 0; i < len(intRomanArray); i = i + 2 {

		x := num / intRomanArray[i].integer
		if i <= len(intRomanArray)-2 {
			if num < intRomanArray[i+2].integer && num >= intRomanArray[i].integer {
				if x <= 3 {
					for j := 0; j < x; j++ {
						ans += intRomanArray[i].symbol
					}
				} else if x == 4 {
					ans = intRomanArray[i].symbol + intRomanArray[i+1].symbol
				} else if x <= 8 {
					x -= 5
					ans = intRomanArray[i+1].symbol
					for j := 0; j < x; j++ {
						ans += intRomanArray[i].symbol
					}
				} else if x == 9 {
					ans = intRomanArray[i].symbol + intRomanArray[i+2].symbol
				} else {
					ans = intRomanArray[i+2].symbol
				}
				num %= intRomanArray[i].integer
			}
		} else {
			for j := 0; j < x; j++ {
				ans += intRomanArray[i].symbol
			}
			num %= intRomanArray[i].integer
		}

	}
	return ans + IntToRoman(num)
}
