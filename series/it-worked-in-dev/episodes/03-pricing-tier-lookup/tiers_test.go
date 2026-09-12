package pricingtierlookup

import (
	"fmt"
	"math/rand"
	"testing"
)

func mk(mins ...int) []Tier {
	out := make([]Tier, len(mins))
	for i, m := range mins {
		out[i] = Tier{MinUnits: m, PricePerK: 100 - i, Name: fmt.Sprintf("tier%d", i)}
	}
	return out
}

func TestBothAgree(t *testing.T) {
	tiers := mk(0, 1_000, 10_000, 100_000)
	cases := []struct {
		units int
		want  string
		ok    bool
	}{
		{0, "tier0", true},
		{1, "tier0", true},
		{999, "tier0", true},
		{1_000, "tier1", true},      // exactly on a boundary belongs to the new tier
		{1_001, "tier1", true},
		{9_999, "tier1", true},
		{10_000, "tier2", true},
		{5_000_000, "tier3", true},  // past the end stays in the last tier
	}
	for _, c := range cases {
		t.Run(fmt.Sprint(c.units), func(t *testing.T) {
			s, sok := TierForScan(tiers, c.units)
			b, bok := TierForBinary(tiers, c.units)
			if s.Name != c.want || !sok {
				t.Errorf("scan = %q/%v, want %q", s.Name, sok, c.want)
			}
			if b.Name != c.want || !bok {
				t.Errorf("binary = %q/%v, want %q", b.Name, bok, c.want)
			}
		})
	}
}

// A price list that does not start at zero has a gap below its first tier, and
// both versions must report "no tier" rather than inventing one.
func TestBelowTheFirstTier(t *testing.T) {
	tiers := mk(500, 1_000)
	for _, units := range []int{0, 1, 499} {
		if _, ok := TierForScan(tiers, units); ok {
			t.Errorf("scan found a tier for %d units", units)
		}
		if _, ok := TierForBinary(tiers, units); ok {
			t.Errorf("binary found a tier for %d units", units)
		}
	}
}

func TestEmptyList(t *testing.T) {
	if _, ok := TierForScan(nil, 5); ok {
		t.Error("scan found a tier in an empty list")
	}
	if _, ok := TierForBinary(nil, 5); ok {
		t.Error("binary found a tier in an empty list")
	}
}

// The scan is the reference implementation: it is obviously correct, so the
// binary version is checked against it rather than against hand-written
// expectations, across every size and every boundary.
func TestBinaryMatchesScan(t *testing.T) {
	rng := rand.New(rand.NewSource(23))
	for trial := 0; trial < 3000; trial++ {
		n := rng.Intn(40) + 1
		mins := make([]int, n)
		cur := rng.Intn(5)
		for i := range mins {
			mins[i] = cur
			cur += rng.Intn(50) + 1
		}
		tiers := mk(mins...)
		// probe every boundary and either side of it, plus random points
		probes := []int{-1, 0, cur + 10}
		for _, m := range mins {
			probes = append(probes, m-1, m, m+1)
		}
		for i := 0; i < 10; i++ {
			probes = append(probes, rng.Intn(cur+20))
		}
		for _, u := range probes {
			s, sok := TierForScan(tiers, u)
			b, bok := TierForBinary(tiers, u)
			if sok != bok || s != b {
				t.Fatalf("trial %d units=%d: scan=%v/%v binary=%v/%v tiers=%v",
					trial, u, s, sok, b, bok, mins)
			}
		}
	}
}
