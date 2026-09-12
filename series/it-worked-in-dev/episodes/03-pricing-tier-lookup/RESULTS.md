# Measured

```
Apple M1 Pro · go1.26.4 darwin/arm64
go test -bench=. -benchtime=200x -run=XXX
```

Reproduce with the command above in this directory.

Each figure is **1,000 lookups**, because one lookup is never the workload -
this is a month of usage rows being priced.

| tiers | scan | binary search | winner |
|---:|---:|---:|---|
| 6 | 2.56 µs | 4.26 µs | **scan, by 1.7x** |
| 50 | 16.6 µs | 14.8 µs | binary, by 1.1x |
| 500 | 98.1 µs | 22.4 µs | binary, by 4.4x |
| 5,000 | 901 µs | 39.4 µs | binary, by 22.9x |

Raw ns: 2,556 / 4,256 · 16,589 / 14,794 · 98,062 / 22,402 · 901,077 / 39,386

## The crossover is the whole result

**At six tiers - which is what a price list actually has - the scan is
1.7 times faster than binary search.** Not equal. Faster.

A six-element scan is six sequential reads of one cache line. Binary search on
six elements is three jumps to unpredictable offsets, with a branch the
predictor cannot learn, and it pays that overhead to avoid work that was never
expensive.

The two meet at **about fifty tiers**. Below that the clever version is a
pessimisation; above it the gap opens fast.

## Growth

Ten times the tiers, from 500 to 5,000:

- scan: 98.1 µs to 901 µs, a factor of **9.2** - linear, as expected
- binary: 22.4 µs to 39.4 µs, a factor of **1.8**, because each ten-fold
  increase adds about three and a third comparisons, not ten times as many

## Neither allocates

Both report `0 B/op` and `0 allocs/op` at every size. The scan holds one
variable; the binary search holds three integers. Nothing is built, so there is
no memory trade here at all - which makes this episode different from the two
before it, where the faster version bought speed with allocations.

The cost is elsewhere: the list must stay sorted, and somebody has to keep it
that way.
