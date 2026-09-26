package invalidationwavefront

import (
	"math/rand"
	"testing"
)

func link(c Cluster, a, b int) {
	c[a] = append(c[a], b)
	c[b] = append(c[b], a)
}

// gossip is n cache nodes where each new node peers with one earlier node,
// so everything is connected, plus `extra` random peerings on top.
func gossip(n, extra int, seed int64) Cluster {
	r := rand.New(rand.NewSource(seed))
	c := make(Cluster, n)
	for i := 1; i < n; i++ {
		link(c, i, r.Intn(i))
	}
	for k := 0; k < extra; k++ {
		a, b := r.Intn(n), r.Intn(n)
		if a != b {
			link(c, a, b)
		}
	}
	return c
}

// ring is n nodes each peering with the next: the slowest shape for gossip.
func ring(n int) Cluster {
	c := make(Cluster, n)
	for i := 0; i < n; i++ {
		link(c, i, (i+1)%n)
	}
	return c
}

// spread picks k distinct origins from n nodes, evenly, so the same k always
// means the same origins.
func spread(n, k int) []int {
	out := make([]int, k)
	for i := range out {
		out[i] = i * n / k
	}
	return out
}

func TestBothAgree(t *testing.T) {
	type tc struct {
		c       Cluster
		origins []int
	}
	cases := []tc{
		{ring(1), []int{0}},
		{ring(12), []int{0}},
		{ring(12), []int{0, 6}},
		{ring(12), []int{3, 3, 3}}, // the same origin, three times
		{gossip(300, 300, 1), spread(300, 1)},
		{gossip(300, 300, 1), spread(300, 30)},
		{gossip(300, 300, 1), spread(300, 300)}, // everyone is an origin
	}
	// Two clusters with no link between them: one is never reached.
	split := make(Cluster, 6)
	link(split, 0, 1)
	link(split, 1, 2)
	link(split, 3, 4)
	link(split, 4, 5)
	cases = append(cases, tc{split, []int{0}}, tc{split, []int{0, 5}})
	for _, s := range shapes {
		if len(s.origins) <= 100 {
			cases = append(cases, tc{s.c, s.origins})
		}
	}
	for _, x := range cases {
		a, b := RoundsBySearchEach(x.c, x.origins), RoundsByOneWave(x.c, x.origins)
		if a != b {
			t.Fatalf("%d nodes, %d origins: search each says %d, one wave says %d",
				len(x.c), len(x.origins), a, b)
		}
	}
	if got := RoundsByOneWave(split, []int{0}); got != -1 {
		t.Fatalf("split cluster from one side: %d rounds, want -1", got)
	}
	if got := RoundsByOneWave(split, []int{0, 5}); got != 2 {
		t.Fatalf("split cluster from both sides: %d rounds, want 2", got)
	}
}

func TestShapes(t *testing.T) {
	for _, s := range shapes {
		searched := len(s.origins) * len(s.c)
		t.Logf("%-14s %6d nodes, %5d origins, %3d rounds  |  nodes visited: "+
			"search each %10d, one wave %6d",
			s.name, len(s.c), len(s.origins), RoundsByOneWave(s.c, s.origins),
			searched, len(s.c))
	}
}
