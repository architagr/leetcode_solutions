package numberofstepstoreduceanumbertozero

func numberOfSteps(num int) int {
	steps := 0
	for num > 0 {
		if num&1 == 1 {
			num--
		} else {
			num >>= 1
		}
		steps++
	}
	return steps
}
