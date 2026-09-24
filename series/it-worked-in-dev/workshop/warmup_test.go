package workshop

import "testing"

// The setup check, and the first lesson, in one file.
//
// `make verify` runs the four benchmarks below. Each pair calls the same
// function with the same argument, once throwing the result away and once
// keeping it. Same work, so each pair should cost the same.
//
// One pair does. The other does not.
//
//	MixDiscarded-8     0.3184n ± 1%
//	MixKept-8           1.054n ± 1%   <- 3.3x apart
//	SumToDiscarded-8     324.9n ± 0%
//	SumToKept-8          324.4n ± 1%   <- the same
//
// Ten runs through benchstat, which is what `make verify` does. A single run
// on a busy machine reported the SumTo pair 1.3x apart when they are the same,
// which would teach the opposite of the lesson. M1 Pro; yours will differ.
//
// Both functions are inlined. `go test -gcflags=-m` prints "inlining call to
// Mix" and "inlining call to SumTo" alike, so inlining is not what separates
// them - which is worth knowing, because "it got inlined" is the usual
// explanation and it is wrong.
//
// What separates them is that Go's SSA eliminates unused *values* and nothing
// in it eliminates an unused *loop*. Mix is straight-line arithmetic, so
// discarding the result leaves nothing to compute. SumTo's thousand additions
// run either way.
//
// The number is its own tell. 0.315 ns on a 3.2 GHz machine is about one clock
// cycle, and five operations do not happen in one cycle. Any benchmark result
// near a single cycle is a benchmark of nothing.
//
// The lesson is not "the compiler deletes your benchmarks" - it deletes far
// less than the folklore says. It is that whether it does depends on the shape
// of what you wrote, the explanation everyone reaches for first is the wrong
// one, and you cannot tell by reading the code. Which is the premise of the
// two days.
//
// Your numbers will differ. Bring them; the ratio is what matters.

// The sink. Package-level, so the compiler cannot prove nobody reads it.
var (
	SinkU uint64
	SinkI int
)

const sumN = 1_000

// --- inlinable: the discarded version really is deleted ---

func BenchmarkMixDiscarded(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mix(uint64(i)) // result dropped on the floor
	}
}

func BenchmarkMixKept(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SinkU = Mix(uint64(i))
	}
}

// --- has a loop, so it is not inlined, so nothing is deleted ---

func BenchmarkSumToDiscarded(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumTo(sumN)
	}
}

func BenchmarkSumToKept(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SinkI = SumTo(sumN)
	}
}

// A plain test as well, so `make verify` proves the module compiles and runs
// before it reaches a benchmark. A failing benchmark and a failing build look
// the same from the terminal otherwise.
func TestSumTo(t *testing.T) {
	cases := []struct{ n, want int }{
		{0, 0},
		{1, 1},
		{10, 55},
		{1_000, 500_500},
	}
	for _, c := range cases {
		if got := SumTo(c.n); got != c.want {
			t.Errorf("SumTo(%d) = %d, want %d", c.n, got, c.want)
		}
	}
}

func TestMixIsDeterministic(t *testing.T) {
	if Mix(0) != Mix(0) {
		t.Fatal("Mix is not deterministic")
	}
	if Mix(1) == Mix(2) {
		t.Fatal("Mix collides on 1 and 2")
	}
}
