// Package workshop holds the setup check for the 2-day Go performance
// workshop, and later the labs.
//
// The two functions below are the setup check's subjects. They are chosen to
// disagree with each other about whether the compiler can delete them, which
// is the point - see warmup_test.go.
package workshop

// Mix is the splitmix64 finalizer. Five operations, no branches, and - the part
// that matters - no loop.
//
// Straight-line arithmetic is what the compiler's dead-value elimination can
// actually remove. Inline it, discard the result, and there is nothing left to
// compute, so a benchmark of it measures an empty loop.
func Mix(x uint64) uint64 {
	x ^= x >> 33
	x *= 0xff51afd7ed558ccd
	x ^= x >> 33
	x *= 0xc4ceb9fe1a85ec53
	x ^= x >> 33
	return x
}

// SumTo adds the integers from 1 to n.
//
// The loop is the whole difference, and not for the reason most people guess.
// SumTo is inlined too - `go test -gcflags=-m` prints "inlining call to SumTo"
// at both benchmark call sites. Inlining is not the discriminator.
//
// The discriminator is that Go's SSA removes unused *values*; nothing in it
// removes an unused *loop*. So the thousand additions run whether or not you
// read the total, and discarding the result of this one costs you nothing.
func SumTo(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i
	}
	return total
}
