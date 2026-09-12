package pricingtierlookup

import (
	"fmt"
	"math/rand"
	"testing"
)

// A price list has a handful of tiers. A rate-limit table or a geo lookup
// built the same way has thousands, which is the point of the larger sizes.
var sizes = []int{6, 50, 500, 5_000}

func listOf(n int) []Tier {
	mins := make([]int, n)
	for i := range mins {
		mins[i] = i * 1_000
	}
	return mk(mins...)
}

// A batch of lookups, because one lookup is never the workload. This is a
// month of usage rows being priced.
const lookups = 1_000

func probes(n, max int) []int {
	rng := rand.New(rand.NewSource(5))
	out := make([]int, n)
	for i := range out {
		out[i] = rng.Intn(max)
	}
	return out
}

func BenchmarkScan(b *testing.B) {
	for _, n := range sizes {
		tiers, ps := listOf(n), probes(lookups, n*1_000)
		b.Run(fmt.Sprint(n), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				for _, u := range ps {
					TierForScan(tiers, u)
				}
			}
		})
	}
}

func BenchmarkBinary(b *testing.B) {
	for _, n := range sizes {
		tiers, ps := listOf(n), probes(lookups, n*1_000)
		b.Run(fmt.Sprint(n), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				for _, u := range ps {
					TierForBinary(tiers, u)
				}
			}
		})
	}
}
