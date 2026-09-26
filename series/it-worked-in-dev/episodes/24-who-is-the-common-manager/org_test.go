package whoisthecommonmanager

import (
	"math/rand"
	"testing"
)

// org builds a company of about n employees: each manager has between lo and
// hi direct reports, filled level by level, IDs in hiring order.
func org(n, lo, hi int, seed int64) (*Employee, []*Employee) {
	r := rand.New(rand.NewSource(seed))
	root := &Employee{ID: 0}
	all := []*Employee{root}
	for i := 0; len(all) < n; i++ {
		m := all[i]
		k := lo + r.Intn(hi-lo+1)
		for j := 0; j < k && len(all) < n; j++ {
			e := &Employee{ID: len(all)}
			m.Reports = append(m.Reports, e)
			all = append(all, e)
		}
	}
	return root, all
}

// pick chooses k distinct attendees, the same ones for the same seed.
func pick(all []*Employee, k int, seed int64) []int {
	r := rand.New(rand.NewSource(seed))
	ids := make([]int, 0, k)
	for _, i := range r.Perm(len(all))[:k] {
		ids = append(ids, all[i].ID)
	}
	return ids
}

// oneTeam picks k attendees from inside a single manager's reporting line,
// the usual meeting.
func oneTeam(all []*Employee, k int, seed int64) []int {
	r := rand.New(rand.NewSource(seed))
	for {
		m := all[r.Intn(len(all))]
		var line []*Employee
		var walk func(e *Employee)
		walk = func(e *Employee) {
			line = append(line, e)
			for _, x := range e.Reports {
				walk(x)
			}
		}
		walk(m)
		if len(line) >= k*2 && len(line) < len(all)/4 {
			return pick(line, k, seed)
		}
	}
}

func TestAllThreeAgree(t *testing.T) {
	root, all := org(3000, 2, 7, 1)
	byID := Number(root)
	cases := [][]int{{0}, {5}, {0, 2999}, {17, 17}, {10, 11}}
	for k := 1; k <= 300; k *= 3 {
		cases = append(cases, pick(all, k, int64(k)), oneTeam(all, k, int64(k)))
	}
	// one attendee is the other's manager: the manager is the answer
	p := PathTo(root, 2500)
	cases = append(cases, []int{p[2].ID, 2500})
	for _, ids := range cases {
		a := ManagerByPaths(root, ids)
		b := ManagerByCount(root, ids)
		c := ManagerByNumbers(root, byID, ids)
		if a != b || b != c {
			t.Fatalf("attendees %v: paths %v, count %v, numbers %v", ids, id(a), id(b), id(c))
		}
	}
	if got := ManagerByCount(root, []int{p[2].ID, 2500}); got != p[2] {
		t.Fatal("a manager and their report should meet at the manager")
	}
}

func id(e *Employee) int {
	if e == nil {
		return -1
	}
	return e.ID
}
