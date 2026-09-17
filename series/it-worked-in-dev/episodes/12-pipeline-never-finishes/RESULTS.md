# Measured

```
Apple M1 Pro · go1.26.4 darwin/arm64
go test -bench=. -benchtime=300ms
```

Reproduce with the command above in this directory.

## The map of seen stages, on its own

| shape | stages | loops | map of seen |
|---|---:|---|---:|
| a hand-written pipeline | 12 | no | 383 ns |
| generated, one step per rule | 200 | no | 9.09 µs |
| a long chain that terminates | 100,000 | no | 5.48 ms |
| a long chain looping near the top | 100,000 | yes | 5.35 ms |
| a long chain looping in the last 1% | 100,000 | yes | 5.38 ms |

Raw ns: 382.7 · 9,089 · 5,475,021 · 5,351,342 · 5,375,737

The three 100,000-stage rows cost the same whether the chain loops after two
stages or after ninety-nine thousand. The map is built either way.

## Both implementations

| shape | stages | map of seen | two pointers | ratio |
|---|---:|---:|---:|---:|
| a hand-written pipeline | 12 | 383 ns | 3.17 ns | 120.6x |
| generated | 200 | 9.09 µs | 195 ns | 46.7x |
| a long chain, no loop | 100,000 | 5.48 ms | 115 µs | 47.8x |
| looping near the top | 100,000 | 5.35 ms | 238 µs | 22.5x |
| looping in the last 1% | 100,000 | 5.38 ms | 208 µs | 25.8x |

Raw ns: 382.7 / 3.172 · 9,089 / 194.6 · 5,475,021 / 114,592 ·
5,351,342 / 237,749 · 5,375,737 / 207,961

Every ratio compares `HasCycleBySeen` against `HasCycleByTwoPointers` on the
same pipeline.

## Memory

| shape | map of seen | two pointers |
|---|---|---|
| 12 stages | 328 B / 3 allocs | 0 B / 0 allocs |
| 200 stages | 9.35 KB / 11 allocs | 0 B / 0 allocs |
| 100,000 stages | **4.73 MB** / 530 allocs | 0 B / 0 allocs |

4.73 megabytes to return a boolean. The two-pointer version holds two
variables and allocates nothing at any size.

## The 12-stage pipeline is the surprising row

120.6x on twelve stages, because the map pays for a map before it walks
anything: 328 bytes and three allocations for a pipeline that fits on a
screen. Two pointers is 3.17 ns, which is a handful of instructions.

In absolute terms 383 ns is still nothing, and no profiler will ever point at
it. The ratio is only interesting where this runs per request.

## Finding where the loop starts

| shape | map of seen | two pointers | ratio |
|---|---:|---:|---:|
| looping near the top | 4.61 ms | 235 µs | 19.6x |
| looping in the last 1% | 4.65 ms | 320 µs | 14.5x |

Raw ns: 4,611,735 / 235,012 · 4,646,652 / 319,764

`EntryByTwoPointers` costs a second walk and still allocates nothing. The gap
narrows — 14.5x against 47.8x — because the second walk is real work the map
version never has to do: it already knew.

## The hop limit

`HasCycleByHopLimit` is what most HTTP clients ship for redirects. On the
150-stage pipeline in `TestHopLimitLiesAboutLongChains` it reports a cycle with
a limit of 100, and there is no cycle. It is not benchmarked here, because a
check that answers the wrong question quickly is not a faster check.

## The trace in the diagrams

`TestDrawnTrace` asserts the five-stage pipeline the walkthrough images draw:
the cursors sit on validate/enrich, enrich/retry, notify/enrich, and meet on
retry at turn 4, with the loop starting at validate. The first draft of those
images drew the wrong turn, which is why the trace is now a test.
