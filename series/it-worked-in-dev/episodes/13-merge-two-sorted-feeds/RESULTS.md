# Measured

```
Apple M1 Pro · go1.26.4 darwin/arm64
go test -bench=. -benchtime=200x
```

Reproduce with the command above in this directory.

Every shape uses `tick=1`, which gives about one event in two a timestamp
shared with another event. That is what a coarse clock and a bulk insert
produce, and it is the case the tie-breaking argument is about, so the
benchmark runs on it rather than on an input with no ties in it.

## Concatenate and sort, on its own

| shape | events | sort.Slice | allocations |
|---|---:|---:|---:|
| one screen of a feed | 20 + 20 | 1.67 µs | 1,512 B, 4 |
| a day for an active account | 500 + 500 | 63.7 µs | 32,872 B, 4 |
| the month view | 20,000 + 20,000 | 3.75 ms | 1,286,274 B, 4 |
| an export nobody paginated | 200,000 + 200,000 | 40.1 ms | 12,804,200 B, 4 |
| one busy source, one quiet one | 200,000 + 50 | 80.4 ms | 6,406,248 B, 4 |

Raw ns: 1,670 · 63,679 · 3,750,073 · 40,056,674 · 80,427,295

The last row is half the data of the row above it and takes **2.0x** as long.

## All four implementations

| shape | events | sort.Slice | sort.SliceStable | slices.SortFunc | merge |
|---|---:|---:|---:|---:|---:|
| one screen | 20 + 20 | 1.67 µs | 1.54 µs | 967 ns | 396 ns |
| a day | 500 + 500 | 63.7 µs | 55.6 µs | 34.2 µs | 5.31 µs |
| the month view | 20,000 + 20,000 | 3.75 ms | 3.36 ms | 2.27 ms | 280 µs |
| an export | 200,000 + 200,000 | 40.1 ms | 48.3 ms | 28.5 ms | 2.63 ms |
| busy plus quiet | 200,000 + 50 | 80.4 ms | 4.69 ms | 40.8 ms | 644 µs |

Raw ns, in column order:

```
1,670 / 1,540 / 967.3 / 395.8
63,679 / 55,585 / 34,246 / 5,308
3,750,073 / 3,360,371 / 2,273,456 / 280,376
40,056,674 / 48,291,248 / 28,510,958 / 2,631,410
80,427,295 / 4,687,550 / 40,787,244 / 644,223
```

## Ratios, with both sides named

`sort.Slice` against `MergeByWalking`, same input:

| shape | ratio |
|---|---:|
| one screen, 20 + 20 | 4.2x |
| a day, 500 + 500 | 12.0x |
| the month view, 20,000 + 20,000 | 13.4x |
| an export, 200,000 + 200,000 | 15.2x |
| busy plus quiet, 200,000 + 50 | 124.8x |

`slices.SortFunc` against `MergeByWalking`, same input:

| shape | ratio |
|---|---:|
| one screen, 20 + 20 | 2.4x |
| a day, 500 + 500 | 6.5x |
| the month view, 20,000 + 20,000 | 8.1x |
| an export, 200,000 + 200,000 | 10.8x |
| busy plus quiet, 200,000 + 50 | 63.3x |

On the 200,000 + 50 shape only:

| comparison | ratio |
|---|---:|
| `sort.Slice` against `sort.SliceStable` | 17.2x |
| `sort.SliceStable` against `MergeByWalking` | 7.3x |
| `sort.Slice` on 200,050 events against the same call on 400,000 | 2.0x |

## Comparisons, counted rather than argued

`go test -run TestComparisonCounts -v`. The comparison function is wrapped in a
counter, so these are exact rather than estimated.

| shape | events | sort.Slice | sort.SliceStable | merge |
|---|---:|---:|---:|---:|
| one screen | 40 | 158 | 83 | 35 |
| a day | 1,000 | 9,366 | 2,569 | 975 |
| the month view | 40,000 | 585,847 | 99,939 | 39,982 |
| an export | 400,000 | 7,229,678 | 993,507 | 399,227 |
| busy plus quiet | 200,050 | 10,455,458 | 240,743 | 105 |

`slices.SortFunc` makes exactly the same number of comparisons as `sort.Slice`
on every shape - it is the same pdqsort. The time difference between those two
columns is per-comparison and per-swap overhead, not algorithm.

## Tie-breaking, measured

`go test -run 'Ties|NewEvent' -v`, on 200 + 200 events with about half of them
sharing a timestamp:

- `sort.Slice` places **271 of 400** events somewhere other than where the
  stable merge places them.
- Appending one event after everything else changes the position of **74 of
  the 400** events in front of it. The merge moves none of them.
