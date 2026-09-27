package runningminwithoutrescan

import "testing"

func TestSessionSizes(t *testing.T) {
	for _, s := range shapes {
		adds, undos, peak, size := 0, 0, 0, 0
		for _, st := range s.steps {
			if st == 0 {
				undos++
				if size > 0 {
					size--
				}
			} else {
				adds++
				size++
			}
			peak = max(peak, size)
		}
		t.Logf("%-12s %6d changes: %6d added, %5d undone, plan peaks at %6d transactions",
			s.name, len(s.steps), adds, undos, peak)
	}
}
