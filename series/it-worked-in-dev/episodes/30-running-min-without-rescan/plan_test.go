package runningminwithoutrescan

import (
	"math/rand"
	"testing"
)

type plan interface {
	Add(int)
	Undo()
	Lowest() int
}

// session builds a plan the way a person does: mostly adding transactions,
// sometimes undoing a few in a row. Big outgoings are rare; salary is monthly.
func session(n int, seed int64) []int {
	r := rand.New(rand.NewSource(seed))
	var steps []int // 0 means undo; anything else is an amount to add
	for i := 0; i < n; i++ {
		switch k := r.Intn(20); {
		case k == 0:
			for j := 0; j < 1+r.Intn(4); j++ {
				steps = append(steps, 0)
			}
		case k == 1:
			steps = append(steps, 250000) // salary
		case k == 2:
			steps = append(steps, -(50000 + r.Intn(100000))) // rent, a big bill
		default:
			steps = append(steps, -(100 + r.Intn(5000))) // everyday spending
		}
	}
	return steps
}

// play runs the session, asking for the lowest balance after every change,
// as the planner's warning does.
func play(p plan, steps []int) int {
	sum := 0
	for _, s := range steps {
		if s == 0 {
			p.Undo()
		} else {
			p.Add(s)
		}
		sum += p.Lowest()
	}
	return sum
}

func TestRescanAndEntryAgree(t *testing.T) {
	steps := session(3000, 1)
	a, b := &PlanByRescan{Opening: 100000}, &PlanByEntry{Opening: 100000}
	for i, s := range steps {
		if s == 0 {
			a.Undo()
			b.Undo()
		} else {
			a.Add(s)
			b.Add(s)
		}
		if a.Lowest() != b.Lowest() {
			t.Fatalf("step %d: rescan %d, entry %d", i, a.Lowest(), b.Lowest())
		}
	}
	for i := 0; i < 5000; i++ { // undo past the start
		a.Undo()
		b.Undo()
	}
	if a.Lowest() != 100000 || b.Lowest() != 100000 {
		t.Fatal("an empty plan's lowest balance is the opening balance")
	}
}

func TestOneMinIsWrongAfterUndo(t *testing.T) {
	p, want := &PlanByOneMin{Opening: 1000}, &PlanByRescan{Opening: 1000}
	for _, a := range []int{-200, -900} { // the second one takes the plan to -100
		p.Add(a)
		want.Add(a)
	}
	p.Undo()
	want.Undo()
	t.Logf("undo the -900: one tracked minimum says %d, the plan's lowest is now %d", p.Lowest(), want.Lowest())
	if p.Lowest() == want.Lowest() {
		t.Fatal("the one-minimum version was right; the episode says it is not")
	}
}
