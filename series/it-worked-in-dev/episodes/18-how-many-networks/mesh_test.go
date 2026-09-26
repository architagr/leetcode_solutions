package howmanynetworks

import (
	"math/rand"
	"testing"
)

func link(m Mesh, a, b int) {
	m[a] = append(m[a], b)
	m[b] = append(m[b], a)
}

// mesh is one network of n services: each new service links to one earlier
// service, so everything is connected, plus `extra` random links on top so it
// is not a tree.
func mesh(n, extra int, seed int64) Mesh {
	r := rand.New(rand.NewSource(seed))
	m := make(Mesh, n)
	for i := 1; i < n; i++ {
		link(m, i, r.Intn(i))
	}
	for k := 0; k < extra; k++ {
		a, b := r.Intn(n), r.Intn(n)
		if a != b {
			link(m, a, b)
		}
	}
	return m
}

// teams is `count` separate networks of `size` services each: every team runs
// its own cluster and nothing links across.
func teams(count, size int) Mesh {
	m := make(Mesh, count*size)
	for t := 0; t < count; t++ {
		base := t * size
		for i := 1; i < size; i++ {
			link(m, base+i, base+i-1)
		}
		link(m, base, base+size-1)
	}
	return m
}

// isolated is n services with no links at all: every service is its own
// network.
func isolated(n int) Mesh { return make(Mesh, n) }

// joined is teams(count, size) plus one link from each team to the next: the
// same services and the same clusters, and count-1 cross-team routes that
// make them all one network.
func joined(count, size int) Mesh {
	m := teams(count, size)
	for t := 1; t < count; t++ {
		link(m, t*size, (t-1)*size)
	}
	return m
}

// visits counts what each version touches, so the claim below is a count
// rather than an estimate.
func visitsByReach(m Mesh) int {
	n := 0
	for s := range m {
		n += len(Reachable(m, s))
	}
	return n
}

func TestBothAgree(t *testing.T) {
	cases := []Mesh{
		{}, isolated(1), isolated(40), teams(1, 3), teams(30, 7),
		mesh(1, 0, 3), mesh(200, 150, 4), joined(20, 5),
	}
	// Two networks plus three loners, linked in an order that makes the
	// smallest service in a network appear last in the scan.
	odd := make(Mesh, 8)
	link(odd, 7, 3)
	link(odd, 3, 0)
	link(odd, 6, 5)
	cases = append(cases, odd)
	for _, s := range shapes {
		cases = append(cases, s.mesh)
	}
	for _, m := range cases {
		if a, b := NetworksByReach(m), NetworksBySweep(m); a != b {
			t.Fatalf("%d services: reach says %d networks, sweep says %d", len(m), a, b)
		}
	}
	if got := NetworksBySweep(odd); got != 5 {
		t.Fatalf("odd mesh: %d networks, want 5", got)
	}
}

func TestVisits(t *testing.T) {
	for _, s := range shapes {
		t.Logf("%-12s %6d services, %5d networks  |  services visited: "+
			"reach per service %11d, one sweep %6d",
			s.name, len(s.mesh), NetworksBySweep(s.mesh), visitsByReach(s.mesh), len(s.mesh))
	}
}
