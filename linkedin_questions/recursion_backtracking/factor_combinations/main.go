package factorcombinations

func getFactors(n int) [][]int {
	if n < 4 {
		return [][]int{}
	}
	factors := make([][]int, 0)
	current := []int{n}
	factor(current, &factors)
	return factors
}

func factor(current []int, factors *[][]int) {
	if len(current) > 1 {
		add(current, factors)
	}
	last := current[len(current)-1]
	current = current[:len(current)-1]
	start := 2
	if len(current) > 0 {
		start = current[len(current)-1]
	}
	for i := start; i <= last/i; i++ {
		if last%i == 0 {
			current = append(current, i)
			x := last / i
			current = append(current, x)
			factor(current, factors)
			current = current[:len(current)-2]
		}
	}
	current = append(current, last)
}

func add(current []int, factors *[][]int) {
	x := append([]int{}, current...)
	*factors = append(*factors, x)
}
