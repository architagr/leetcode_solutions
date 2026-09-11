package countnumberofdistinctintegersafterreverseoperations

func countDistinctIntegers(nums []int) int {
	m := make(map[int]bool)
	for _, val := range nums {
		m[val] = true
		m[reverse(val)] = true
	}
	return len(m)
}

func reverse(a int) int {
	x := 0
	for a > 0 {
		x = (x * 10) + (a % 10)
		a /= 10
	}
	return x
}
