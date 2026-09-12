# Measured

```
Apple M1 Pro · go1.26.4 darwin/arm64
go test -bench=. -benchtime=200x
```

Reproduce with the command above in this directory.

## Balanced trees — the difference barely shows

| shape | directories | brute force | postorder | ratio |
|---|---:|---:|---:|---:|
| depth 3, breadth 4 | 85 | 4.96 µs | 3.91 µs | 1.3x |
| depth 6, breadth 3 | 1,093 | 81.3 µs | 63.7 µs | 1.3x |
| depth 12, breadth 2 | 8,191 | 865 µs | 631 µs | 1.4x |

Raw: 4,960 / 3,909 · 81,281 / 63,661 · 864,951 / 630,819 ns.

Eight thousand directories, and the naive version is 40% slower. That is not a
reason to change anything.

## Deep, narrow trees — it falls apart

| shape | directories | brute force | postorder | ratio |
|---|---:|---:|---:|---:|
| chain, depth 200 | 201 | 157 µs | 12.5 µs | **12.6x** |
| chain, depth 800 | 801 | 2,633 µs | 76.3 µs | **34.5x** |

Raw: 157,175 / 12,497 · 2,633,461 / 76,335 ns.

The input ratio between those two shapes is **10.2x** - 8,191 directories
against 801 - which is the figure the hero card and the posts use when they say
ten times smaller.

Two different ratios live in these tables and they must not be run together.

The **34.5x** is brute force against postorder on the same 800-deep chain - one
shape, two implementations.

The cross-shape comparison is a different number: **801 directories in a chain
took 2.63 ms against 0.86 ms for 8,191 directories in a balanced tree,
which is 3.0x slower on a tenth of the input.** That is the one that says size
is not what hurts and depth is.

## The growth rates

Going from depth 200 to depth 800 — four times the directories:

- brute force: 157 µs → 2,633 µs (157,175 → 2,633,461 ns), a factor of **16.8**. Four times the input
  for sixteen times the work is quadratic.
- postorder: 12.5 µs → 76.3 µs (12,497 → 76,335 ns), a factor of **6.1**, against 4x the input. Linear,
  plus the map growing.

## Allocations are identical

Both report the same `B/op` and `allocs/op` at every shape. The output map is the
same size either way, so the difference is entirely repeated arithmetic and pointer
chasing — not memory pressure.
