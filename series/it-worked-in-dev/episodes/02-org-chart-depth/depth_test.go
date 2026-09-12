package orgchartdepth

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

// chain builds a straight reporting line: p0 <- p1 <- ... <- p(n-1).
func chain(n int) []Person {
	out := make([]Person, n)
	for i := range out {
		out[i] = Person{ID: fmt.Sprint(i)}
		if i > 0 {
			out[i].ManagerID = fmt.Sprint(i - 1)
		}
	}
	return out
}

// bushy builds a tree where every manager has `width` reports.
func bushy(levels, width int) []Person {
	people := []Person{{ID: "0"}}
	frontier := []string{"0"}
	next := 1
	for l := 0; l < levels; l++ {
		var kids []string
		for _, m := range frontier {
			for w := 0; w < width; w++ {
				id := fmt.Sprint(next)
				next++
				people = append(people, Person{ID: id, ManagerID: m})
				kids = append(kids, id)
			}
		}
		frontier = kids
	}
	return people
}

func TestBothAgree(t *testing.T) {
	cases := []struct {
		name   string
		people []Person
	}{
		{"single person", chain(1)},
		{"straight line of 5", chain(5)},
		{"bushy 3x3", bushy(3, 3)},
		{"wide and shallow", bushy(1, 40)},
		{"deep and narrow", chain(60)},
		{"empty", nil},
		// Two separate trees: an org after an acquisition, before anybody
		// wires the two CEOs together. Both versions have to cope.
		{"two roots", append(chain(3), []Person{{ID: "x"}, {ID: "y", ManagerID: "x"}}...)},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			up := LevelsByWalkingUp(c.people)
			down := LevelsByLevelOrder(c.people)
			if !reflect.DeepEqual(up, down) {
				t.Fatalf("disagreed:\n  up   %v\n  down %v", up, down)
			}
		})
	}
}

func TestKnownLevels(t *testing.T) {
	people := []Person{
		{ID: "ceo"},
		{ID: "vp", ManagerID: "ceo"},
		{ID: "dir", ManagerID: "vp"},
		{ID: "eng", ManagerID: "dir"},
	}
	want := map[string]int{"ceo": 0, "vp": 1, "dir": 2, "eng": 3}
	if got := LevelsByLevelOrder(people); !reflect.DeepEqual(got, want) {
		t.Errorf("level order = %v, want %v", got, want)
	}
	if got := LevelsByWalkingUp(people); !reflect.DeepEqual(got, want) {
		t.Errorf("walking up = %v, want %v", got, want)
	}
}

func TestBothAgreeOnRandomOrgs(t *testing.T) {
	rng := rand.New(rand.NewSource(11))
	for trial := 0; trial < 1000; trial++ {
		n := rng.Intn(60) + 1
		people := make([]Person, n)
		for i := range people {
			people[i] = Person{ID: fmt.Sprint(i)}
			if i > 0 {
				// a random earlier person, so the graph is always a forest
				people[i].ManagerID = fmt.Sprint(rng.Intn(i))
			}
		}
		rng.Shuffle(len(people), func(a, b int) { people[a], people[b] = people[b], people[a] })
		if up, down := LevelsByWalkingUp(people), LevelsByLevelOrder(people); !reflect.DeepEqual(up, down) {
			t.Fatalf("trial %d disagreed", trial)
		}
	}
}

// The write-up claims the upward version does sum-of-depths hops. Counted,
// rather than asserted in prose: a straight line of n costs 0+1+...+(n-1).
func TestUpwardHopCount(t *testing.T) {
	for _, n := range []int{10, 50, 200} {
		people := chain(n)
		byID := make(map[string]Person, len(people))
		for _, p := range people {
			byID[p.ID] = p
		}
		hops := 0
		for _, p := range people {
			for cur := p; cur.ManagerID != ""; hops++ {
				cur = byID[cur.ManagerID]
			}
		}
		if want := n * (n - 1) / 2; hops != want {
			t.Errorf("n=%d: %d hops, want %d", n, hops, want)
		}
	}
}
