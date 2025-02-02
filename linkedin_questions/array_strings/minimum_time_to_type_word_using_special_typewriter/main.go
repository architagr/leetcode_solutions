package minimumtimetotypewordusingspecialtypewriter

func minTimeToType(word string) int {
	current := byte('a')
	ans := 0
	for i := 0; i < len(word); i++ {
		ans += findMinDistance(current, word[i])
		current = word[i]
	}
	return ans
}

func findMinDistance(a, b byte) int {
	inta := int(a - 'a')
	intb := int(b - 'a')
	if inta > intb {
		inta, intb = intb, inta
	}
	return minVal((intb-inta), (26-intb+inta)) + 1
}

func minVal(a, b int) int {
	if a < b {
		return a
	}
	return b
}
