package lastninonepass

import (
	"fmt"
	"testing"
)

// How long a log is when somebody asks for the last 200 lines of it.
var sizes = []struct {
	name  string
	lines int
}{
	{"dev_200", 200},         // the log in front of you while you write the code
	{"request_10k", 10000},   // one request's worth of a chatty service
	{"pod_1m", 1000000},      // a pod's log since the last restart
	{"rotation_2m", 2000000}, // what a rotation period holds
}

const n = 200

func runLog(b *testing.B, fn func(open func() Log) []string) {
	for _, s := range sizes {
		open := logOf(s.lines)
		b.Run(s.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				fn(open)
			}
		})
	}
}

func BenchmarkCollecting(b *testing.B) {
	runLog(b, func(open func() Log) []string { return LastNByCollecting(open(), n) })
}

func BenchmarkCollectingAndCopying(b *testing.B) {
	runLog(b, func(open func() Log) []string { return LastNByCollectingAndCopying(open(), n) })
}

func BenchmarkCountingTwice(b *testing.B) {
	runLog(b, func(open func() Log) []string { return LastNByCountingTwice(open, n) })
}

func BenchmarkRing(b *testing.B) {
	runLog(b, func(open func() Log) []string { return LastNByRing(open(), n) })
}

// The same question on a chain that is already in memory, where the gap costs
// two variables and the counting version costs a second walk.
var chainSizes = []int{1000, 100000, 1000000}

func BenchmarkNthFromEndByCounting(b *testing.B) {
	for _, size := range chainSizes {
		head := chainOf(size)
		b.Run(fmt.Sprintf("entries_%d", size), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				NthFromEndByCounting(head, n)
			}
		})
	}
}

func BenchmarkNthFromEndByGap(b *testing.B) {
	for _, size := range chainSizes {
		head := chainOf(size)
		b.Run(fmt.Sprintf("entries_%d", size), func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				NthFromEndByGap(head, n)
			}
		})
	}
}
