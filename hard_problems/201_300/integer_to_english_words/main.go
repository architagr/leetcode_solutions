package integertoenglishwords

import "strings"

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	thousand, million, billion, trillion := "", "", "", ""

	hundred := numberToWordsForUnder1000(num % 1000)
	num /= 1000

	for i := 1; num > 0; i++ {
		switch i {
		case 1:
			thousand = numberToWordsForUnder1000(num % 1000)
		case 2:
			million = numberToWordsForUnder1000(num % 1000)
		case 3:
			billion = numberToWordsForUnder1000(num % 1000)
		case 4:
			trillion = numberToWordsForUnder1000(num % 1000)
		}
		num /= 1000
	}
	result := ""
	if len(trillion) > 0 {
		result += " " + trillion + " Trillion"
	}
	if len(billion) > 0 {
		result += " " + billion + " Billion"
	}
	if len(million) > 0 {
		result += " " + million + " Million"
	}
	if len(thousand) > 0 {
		result += " " + thousand + " Thousand"
	}
	return strings.Trim(result+" "+hundred, " ")
}

func numberToWordsForUnder1000(num int) string {
	if num > 999 {
		return ""
	}
	result := []string{}
	if num > 99 {
		result = append(result, oncePlace(num/100), "Hundred")
		num %= 100
	}
	if num > 9 {
		if num > 19 {
			result = append(result, tensePlace(num/10))
			num %= 10
		} else {
			str := ""
			switch num {
			case 10:
				str = "Ten"
			case 11:
				str = "Eleven"
			case 12:
				str = "Twelve"
			case 13:
				str = "Thirteen"
			case 15:
				str = "Fifteen"
			case 18:
				str = "Eighteen"
			case 14, 16, 17, 19, 20:
				result = append(result, oncePlace(num%10)+"teen")
			}
			result = append(result, str)
			num = 0
		}
	}
	if num > 0 {
		result = append(result, oncePlace(num))
	}

	return strings.Trim(strings.Join(result, " "), " ")
}

func tensePlace(num int) string {
	words := map[int]string{
		2: "Twenty",
		3: "Thirty",
		4: "Forty",
		5: "Fifty",
		6: "Sixty",
		7: "Seventy",
		8: "Eighty",
		9: "Ninety",
	}
	return words[num]
}

func oncePlace(num int) string {
	words := map[int]string{
		1: "One",
		2: "Two",
		3: "Three",
		4: "Four",
		5: "Five",
		6: "Six",
		7: "Seven",
		8: "Eight",
		9: "Nine",
	}
	return words[num]
}
