package matchingbracketsconfig

import "strings"

// A rules file for the pricing engine: expressions like
//
//	discount(max(cart[total], 500), {tier: gold})
//
// The linter rejects a file whose (), [] and {} do not nest properly, before
// the parser gets a chance to produce an error on line 4,000 about something
// on line 12.

// BalancedByCounting is the check that ships first: every kind of bracket
// opens as often as it closes, and never closes more than it has opened.
// It is fast, and it passes "([)]".
func BalancedByCounting(s string) bool {
	var round, square, curly int
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			round++
		case ')':
			round--
		case '[':
			square++
		case ']':
			square--
		case '{':
			curly++
		case '}':
			curly--
		}
		if round < 0 || square < 0 || curly < 0 {
			return false
		}
	}
	return round == 0 && square == 0 && curly == 0
}

// BalancedByErasing is what I would write once the counting check is known to
// be wrong: a properly nested file always contains an adjacent pair - (), []
// or {} - somewhere, so keep deleting them. If nothing is left, it nested.
func BalancedByErasing(s string) bool {
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		if strings.IndexByte("()[]{}", s[i]) >= 0 {
			b.WriteByte(s[i])
		}
	}
	t := b.String()
	for {
		u := strings.ReplaceAll(t, "()", "")
		u = strings.ReplaceAll(u, "[]", "")
		u = strings.ReplaceAll(u, "{}", "")
		if u == t {
			return t == ""
		}
		t = u
	}
}

// BalancedByStack keeps the brackets that are open and not yet closed. A
// closer must match the one opened most recently, which is the top.
func BalancedByStack(s string) bool {
	stack := make([]byte, 0, 64)
	for i := 0; i < len(s); i++ {
		var open byte
		switch c := s[i]; c {
		case '(', '[', '{':
			stack = append(stack, c)
			continue
		case ')':
			open = '('
		case ']':
			open = '['
		case '}':
			open = '{'
		default:
			continue
		}
		// Nothing open, or the wrong thing open: no later character can fix it.
		if len(stack) == 0 || stack[len(stack)-1] != open {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0 // anything still open was never closed
}
