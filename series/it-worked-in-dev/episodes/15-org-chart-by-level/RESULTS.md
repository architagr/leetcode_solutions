# Measured

```
Apple M1 Pro · go1.26.4 darwin/arm64
go test -bench=. -benchtime=200x
go test -run TestVisitsBeforeLevelComplete -v
```

Reproduce with the commands above in this directory.

Four shapes. Two of them are here because the queue wins nothing on them.

| shape | people | what it is |
|---|---:|---|
| `org_5k` | 5,461 | six levels, four reports each |
| `org_56k` | 55,987 | six levels, six reports each |
| `flat_50k` | 50,001 | one person at the top, everybody else under them |
| `chain_2k` | 2,001 | the reorg nobody meant to ship: one report each |

## Producing the whole chart

Both functions return the identical `[][]string`.

| shape | people | one depth-keyed walk | a queue | ratio |
|---|---:|---:|---:|---:|
| `org_5k` | 5,461 | 86.8 µs | 67.7 µs | 1.28x |
| `org_56k` | 55,987 | 894 µs | 895 µs | 1.00x |
| `flat_50k` | 50,001 | 944 µs | 386 µs | 2.45x |
| `chain_2k` | 2,001 | 92.7 µs | 102 µs | 0.91x |

Raw ns: 86,811 / 67,688 · 893,794 / 894,899 · 943,751 / 385,701 ·
92,730 / 101,836

On `org_56k` it is a dead heat: **1.00x**. On `chain_2k` the recursive walk is
the faster one - the queue is **1.10x** slower there, because a chain makes it
allocate a new frontier per level, 4,013 allocations against 2,014.

Allocated, same runs:

| shape | one depth-keyed walk | a queue |
|---|---:|---:|
| `org_5k` | 312,216 B, 56 allocs | 251,073 B, 25 allocs |
| `org_56k` | 3,951,516 B, 78 allocs | 3,032,751 B, 33 allocs |
| `flat_50k` | 4,344,268 B, 28 allocs | 1,204,288 B, 4 allocs |
| `chain_2k` | 189,688 B, 2,014 allocs | 221,664 B, 4,013 allocs |

## Producing the first three levels

What the page asks for on first paint. `FirstLevelsByDepthKeyedWalk` has no way
to stop early, so it is the whole walk followed by a slice.

| shape | people | one depth-keyed walk | a queue | ratio |
|---|---:|---:|---:|---:|
| `org_5k` | 5,461 | 71.2 µs | 366 ns | 194.8x |
| `org_56k` | 55,987 | 917 µs | 385 ns | 2383.3x |
| `flat_50k` | 50,001 | 965 µs | 431 µs | 2.24x |
| `chain_2k` | 2,001 | 87.4 µs | 195 ns | 447.8x |

Raw ns: 71,211 / 365.6 · 916,616 / 384.6 · 964,598 / 431,180 · 87,404 / 195.2

Every ratio compares `FirstLevelsByDepthKeyedWalk(root, 3)` against
`FirstLevelsByQueue(root, 3)` on the same chart.

`flat_50k` is the shape where it does not help: level 1 is everybody, so
finishing level 1 means looking at all 50,001 people either way. **2.24x**, and
that is the allocation pattern rather than the algorithm.

Allocated, same runs:

| shape | one depth-keyed walk | a queue |
|---|---:|---:|
| `org_5k` | 312,216 B, 56 allocs | 680 B, 9 allocs |
| `org_56k` | 3,951,515 B, 78 allocs | 1,608 B, 10 allocs |
| `flat_50k` | 4,344,265 B, 28 allocs | 1,605,738 B, 6 allocs |
| `chain_2k` | 189,688 B, 2,014 allocs | 144 B, 6 allocs |

## When each walk finishes a level, counted rather than timed

`TestVisitsBeforeLevelComplete`. Three counts per row: the visit on which the
depth-first walk reaches the **last member** of that level, the visit on which
it can **say** the level is finished, and what the queue looks at.

| shape | level | walk has it at | walk can say so at | queue |
|---|---:|---:|---:|---:|
| `org_5k` | 1 | 4,097 | 5,461 | 5 |
| `org_5k` | 2 | 5,121 | 5,461 | 21 |
| `org_56k` | 1 | 46,657 | 55,987 | 7 |
| `org_56k` | 2 | 54,433 | 55,987 | 43 |
| `flat_50k` | 1 | 50,001 | 50,001 | 50,001 |
| `chain_2k` | 1 | 2 | 2,001 | 2 |

The `chain_2k` row is the one worth staring at. The depth-first walk has all of
level 1 after **two** visits and cannot declare it for another 1,999, because
nothing in a depth-first walk rules out another node turning up at depth 1.

The `flat_50k` row is the counter-example: every number is 50,001, and no
ordering helps.
