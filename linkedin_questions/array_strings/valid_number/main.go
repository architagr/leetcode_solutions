package validnumber

import "strings"

type numberModel struct {
	integer    strings.Builder
	hasDecimal bool
	decimal    strings.Builder
	sign       rune
}

func isNumber(s string) bool {
	// this if parsing till we get 'e' or 'E'
	n, eindex, ok := parseIntegerDecimal(s)
	if ok && eindex != -1 {
		// if 'e' or 'E' is the last chat in string
		if eindex+1 == len(s) {
			return false
		}
		// this if parsing string after 'e' or 'E'
		n, eindex, ok = parseIntegerDecimal(s[eindex+1:])
		if ok {
			// this part of string should only be integer so no decimal
			// and should also not have additional 'e' or 'E'
			if n.hasDecimal || eindex != -1 {
				return false
			}
		}
	}
	return ok
}

func parseIntegerDecimal(s string) (result *numberModel, eIndex int, ok bool) {
	result = &numberModel{}
	result.decimal = strings.Builder{}
	result.decimal.Grow(len(s))

	result.integer = strings.Builder{}
	result.integer.Grow(len(s))
	eIndex = -1
	ok = true
	errorResponse := func() {
		result = nil
		ok = false
	}
	for i, r := range s {
		if r == 'e' || r == 'E' {
			eIndex = i
			break
		} else if r == '-' || r == '+' {
			// sign shoulf not appear again and it should only be on the index 0
			if result.sign == '-' || result.sign == '+' || i != 0 {
				errorResponse()
				return
			}
			result.sign = r
		} else if r == '.' {
			// we should not have multiple decimal
			if result.hasDecimal {
				errorResponse()
				return
			}
			result.hasDecimal = true
		} else if r >= '0' && r <= '9' {
			if result.hasDecimal {
				result.decimal.WriteRune(r)
			} else {
				result.integer.WriteRune(r)
			}
		} else {
			// if we have any extra char then '+', '-', '.', 'e', 'E', '0' to '9' then it is invalid
			errorResponse()
			return
		}
	}
	if eIndex == 0 || result.decimal.Len()+result.integer.Len() == 0 {
		errorResponse()
		return
	}
	return
}
