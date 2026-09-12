package valid_parentheses

func IsValid(A string) bool {
	// Maps a closing bracket to the opening bracket it must match. Keying by
	// the closer is what lets the loop below decide with one lookup.
	pairs := map[byte]byte{')': '(', ']': '[', '}': '{'}

	// At most one opener per character, so the stack never has to grow.
	stack := make([]byte, 0, len(A))

	for i := 0; i < len(A); i++ {
		open, isCloser := pairs[A[i]]
		if !isCloser {
			stack = append(stack, A[i])
			continue
		}
		// A closer with nothing open, or with the wrong thing open, can never
		// be rescued by the rest of the string.
		if len(stack) == 0 || stack[len(stack)-1] != open {
			return false
		}
		stack = stack[:len(stack)-1]
	}

	// Anything still open was never closed.
	return len(stack) == 0
}
