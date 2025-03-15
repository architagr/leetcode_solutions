package robotreturntoorigin

func judgeCircle(moves string) bool {
	vertical, horizontal := 0, 0

	for i := 0; i < len(moves); i++ {
		if moves[i] == 'U' {
			vertical++
		} else if moves[i] == 'D' {
			vertical--
		} else if moves[i] == 'R' {
			horizontal++
		} else if moves[i] == 'L' {
			horizontal--
		}
	}

	return vertical == 0 && horizontal == 0
}
