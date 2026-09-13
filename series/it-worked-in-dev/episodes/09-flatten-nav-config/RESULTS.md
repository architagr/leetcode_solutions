# Measured

```
Apple M1 Pro · go1.26.4 darwin/arm64
go test -bench=. -benchtime=300ms
```

Reproduce with the command above in this directory. Time-based, not a fixed
iteration count - see episode 7's RESULTS for why that matters here.

| shape | items | concat | appending | ratio |
|---|---:|---:|---:|---:|
| a nav menu, 2 deep | 43 | 2.41 µs | 729 ns | 3.3x |
| a docs sidebar, 3 deep | 156 | 11.9 µs | 1.88 µs | 6.3x |
| 100 levels | 101 | 86.6 µs | 2.31 µs | 37.6x |
| 600 levels | 601 | **2.41 ms** | 11.0 µs | **219.7x** |

Raw ns: 2,413 / 729.3 · 11,933 / 1,882 · 86,589 / 2,306 · 2,411,063 / 10,974

## Depth, not size, again

101 items take **7 times longer** than 156 items, because those 101 are a
hundred levels deep and the 156 are three.

100 levels to 600 is six times the depth for **27.8 times the work**.

## Memory is where it is starkest

| shape | concat | appending |
|---|---|---|
| a nav menu, 43 | 6.30 KB / 67 allocs | 3.00 KB / 3 allocs |
| a docs sidebar, 156 | 27.1 KB / 263 allocs | 7.00 KB / 4 allocs |
| 600 levels, 601 | **11.4 MB / 5,357 allocs** | **31.0 KB / 6 allocs** |

**376x the memory** at 600 deep, across **893x** as many allocations, to
produce a 601-row slice.

## Presizing the slice: worth it, until it is not

`FlattenPresized` counts the items first and allocates once at exactly the right
size, so `append` never grows the slice.

| shape | appending | presized | |
|---|---:|---:|---|
| a nav menu, 43 | 729 ns | 393 ns | 1.9x faster |
| a docs sidebar, 156 | 1.88 µs | 1.43 µs | 1.3x faster |
| 100 levels | 2.31 µs | 2.10 µs | 1.1x faster |
| 600 levels | 11.0 µs | 11.9 µs | **1.08x slower** |

It is always one allocation instead of three to six, which is the point of
doing it. But it pays for that with **an extra full pass over the tree**, and by
600 levels that pass costs more than the handful of re-allocations it avoids.

A nice reminder that "allocate once" is a heuristic, not a law: the counting
pass is real work, and `append`'s doubling is already close to free.

## Row copies

`TestConcatRowCopyCount` asserts a row is copied once per ancestor, so on a
chain the total is the sum of the depths - 180,300 copies at 600 deep. That is
what 2.41 ms is spent on.
