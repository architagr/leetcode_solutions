package number_after_a_double_reversal

func IsSameAfterReversals(num int) bool {
	if num == 0 {
		return true
	}

	return num%10 != 0
}
