package dedupeuserlist

import (
	"fmt"
	"testing"
)

// Sized to bracket the crossover rather than to make a point. 100 against 100
// is a form submission; 20,000 against 20,000 is a nightly import, and that is
// not a large number.
var sizes = []int{100, 1_000, 5_000, 20_000}

// Half the incoming emails already exist, which is the realistic case and also
// the fair one: a full miss makes the nested version scan the entire existing
// list every time, which would flatter the comparison.
func build(n int) (incoming, existing []User) {
	existing = make([]User, n)
	for i := range existing {
		existing[i] = User{ID: fmt.Sprint(i), Email: fmt.Sprintf("user%d@example.com", i)}
	}
	incoming = make([]User, n)
	for i := range incoming {
		if i%2 == 0 {
			incoming[i] = User{Email: fmt.Sprintf("user%d@example.com", i)}
		} else {
			incoming[i] = User{Email: fmt.Sprintf("new%d@example.com", i)}
		}
	}
	return incoming, existing
}

func BenchmarkNested(b *testing.B) {
	for _, n := range sizes {
		in, ex := build(n)
		b.Run(fmt.Sprint(n), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				MissingNested(in, ex)
			}
		})
	}
}

func BenchmarkSet(b *testing.B) {
	for _, n := range sizes {
		in, ex := build(n)
		b.Run(fmt.Sprint(n), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				MissingSet(in, ex)
			}
		})
	}
}
