package matchingbracketsconfig

import (
	"math/rand"
	"strings"
	"testing"
)

// rule builds one realistic rule line with some nesting.
func rule(r *rand.Rand) string {
	var b strings.Builder
	var emit func(depth int)
	emit = func(depth int) {
		b.WriteString("fn")
		kinds := []string{"()", "[]", "{}"}
		k := kinds[r.Intn(3)]
		b.WriteByte(k[0])
		for i := 0; i < 1+r.Intn(3); i++ {
			if i > 0 {
				b.WriteString(", ")
			}
			if depth < 4 && r.Intn(3) == 0 {
				emit(depth + 1)
			} else {
				b.WriteString("cart_total_42")
			}
		}
		b.WriteByte(k[1])
	}
	emit(0)
	return b.String()
}

// rules is a file of n rule lines.
func rules(n int, seed int64) string {
	r := rand.New(rand.NewSource(seed))
	lines := make([]string, n)
	for i := range lines {
		lines[i] = rule(r)
	}
	return strings.Join(lines, "\n")
}

// nested is one expression n levels deep, the shape a generated config takes
// when a tool wraps every condition in another group.
func nested(n int) string {
	kinds := "([{"
	closers := ")]}"
	var b strings.Builder
	for i := 0; i < n; i++ {
		b.WriteByte(kinds[i%3])
		b.WriteString("x ")
	}
	for i := n - 1; i >= 0; i-- {
		b.WriteByte(closers[i%3])
	}
	return b.String()
}

// crossed swaps two closers deep inside an otherwise valid file: every count
// still balances, the nesting does not.
func crossed(s string) string {
	b := []byte(s)
	for i := len(b) / 2; i < len(b)-1; i++ {
		if (b[i] == ')' || b[i] == ']' || b[i] == '}') && strings.IndexByte(")]}", b[i+1]) >= 0 && b[i] != b[i+1] {
			b[i], b[i+1] = b[i+1], b[i]
			return string(b)
		}
	}
	panic("no adjacent closers to swap")
}

func TestCountingPassesCrossedBrackets(t *testing.T) {
	for _, s := range []string{"([)]", "{(})", crossed(nested(30)), crossed(rules(200, 1))} {
		if !BalancedByCounting(s) {
			t.Fatalf("counting rejected %.20q; the episode says it does not", s)
		}
		if BalancedByErasing(s) || BalancedByStack(s) {
			t.Fatalf("a correct check accepted %.20q", s)
		}
	}
	t.Log(`"([)]": counting says balanced, the stack says not`)
}

func TestCorrectVersionsAgree(t *testing.T) {
	cases := []string{"", "x", "()", ")(", "(", ")", "([])", "([)]", "{[()()]}", "((((", "}}}}", "(]"}
	cases = append(cases, rules(50, 2), crossed(rules(50, 3)), nested(40), crossed(nested(40)))
	for _, s := range shapes {
		cases = append(cases, s.text)
	}
	for _, s := range cases {
		if a, b := BalancedByErasing(s), BalancedByStack(s); a != b {
			t.Fatalf("%.30q: erasing %v, stack %v", s, a, b)
		}
	}
}

func TestPasses(t *testing.T) {
	for _, s := range shapes {
		t.Logf("%-14s %8d bytes, deepest nesting %5d", s.name, len(s.text), depth(s.text))
	}
}

func depth(s string) int {
	d, most := 0, 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '[', '{':
			d++
			most = max(most, d)
		case ')', ']', '}':
			d--
		}
	}
	return most
}
