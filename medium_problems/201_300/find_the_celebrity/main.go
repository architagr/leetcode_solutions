package findthecelebrity

/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */
func solution(knows func(a int, b int) bool) func(n int) int {
	return func(n int) int {
		m := make(map[int]int8)
		for i := 0; i < n; i++ {
			m[i] = -1
		}
		start := 0
		i := 1
		for len(m) > 1 {
			next := (start + i) % n
			if start == next {
				m = map[int]int8{start: -1}
				break
			}
			if knows(start, next) {
				// start can not be a celibrity as he knows someone
				delete(m, start)
				// we want to get a new number who is not yet explored
				start = getNextPossible(next, n, m)
				i = 1
			} else {
				// next can not be celibrity as he is not known by start
				delete(m, next)
				i++
			}
		}
		result := -1
		if len(m) == 1 {
			// one final validatation to confirn the ans
			for r := range m {
				result = r
			}
			for i := 0; i < n; i++ {
				if i == result {
					continue
				}
				if knows(result, i) || !knows(i, result) {
					result = -1
					break
				}
			}
		}
		return result
	}
}
func getNextPossible(current, n int, m map[int]int8) int {

	for i := 0; i < n; i++ {
		res := (current + i) % n
		if _, ok := m[res]; ok {
			return res
		}
	}
	return current
}
