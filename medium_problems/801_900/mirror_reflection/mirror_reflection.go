package mirror_reflection

func MirrorReflection(p, q int) int {
	for q%2 == 0 && p%2 == 0 {
		q /= 2
		p /= 2
	}

	if q%2 == 0 && p%2 != 0 {
		return 0
	} else if q%2 != 0 && p%2 == 0 {
		return 2
	} else if q%2 != 0 && p%2 != 0 {
		return 1
	}

	return -1
}
