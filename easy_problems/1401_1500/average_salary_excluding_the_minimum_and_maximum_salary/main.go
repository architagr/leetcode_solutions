package averagesalaryexcludingtheminimumandmaximumsalary

func average(salary []int) float64 {
	max, min, n := salary[0], salary[0], len(salary)
	sum := 0
	for i := 0; i < n; i++ {
		sum += salary[i]
		if max < salary[i] {
			max = salary[i]
		}

		if min > salary[i] {
			min = salary[i]
		}

	}

	sum -= max + min
	return float64(sum) / float64(n-2)
}
