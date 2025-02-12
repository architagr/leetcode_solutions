package valid_parentheses

func IsValid(A string) bool {
	n := len(A)
	stack := make([]byte, n)
	lastElement := -1
	m := map[byte]byte{'[': ']', '{': '}', '(': ')'}

	for i := 0; i < n; i++ {
		if lastElement < 0 {
			lastElement++
			stack[lastElement] = A[i]
			continue
		}

		v := m[stack[lastElement]]
		if v == A[i] {
			lastElement--
		} else {
			lastElement++
			stack[lastElement] = A[i]
		}

	}
	return lastElement == -1
}
