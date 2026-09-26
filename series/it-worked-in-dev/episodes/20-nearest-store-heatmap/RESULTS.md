# Measured

```
Apple M1 Pro · go1.26.4 darwin/arm64
go test -bench=. -benchtime=20x      (three runs, median reported)
go test -run TestSweepsBreakWhenAWalkMustDoubleBack -v
```

Reproduce with the commands above in this directory.

Every shape is the same 500 x 500 city, 250,000 blocks, with stores scattered
at random (fixed seeds, so the same count is always the same stores).

## Time

| shape | stores | scan every store | multi-source BFS | two sweeps |
|---|---:|---:|---:|---:|
| `stores_1` | 1 | 1.03 ms | 2.25 ms | 1.22 ms |
| `stores_10` | 10 | 2.92 ms | 2.22 ms | 1.19 ms |
| `stores_100` | 100 | 21.8 ms | 2.26 ms | 1.22 ms |
| `stores_1000` | 1,000 | 219 ms | 2.80 ms | 1.48 ms |

Raw ns, in column order:

```
1,025,958 / 2,247,356 / 1,220,246
2,924,729 / 2,217,275 / 1,191,550
21,773,417 / 2,262,615 / 1,216,802
218,805,908 / 2,800,675 / 1,482,408
```

`stores_1` against `stores_1000`, same city: scanning every store goes from
1.03 ms to 219 ms, **213x**. BFS goes up **1.25x**, the two sweeps **1.21x**.

Scan against two sweeps, same stores: **0.841x** at one store (the scan is the
faster of the two), **2.45x** at ten, **17.9x** at a hundred, **148x** at a
thousand.

BFS against two sweeps, same stores: **1.84x**, **1.86x**, **1.86x**, **1.89x**.

Scan against BFS at one store: **0.457x**; the scan takes less than half the
time.

## Allocated, same runs

| shape | scan every store | multi-source BFS | two sweeps |
|---|---:|---:|---:|
| every shape | 1.96 MB, 501 allocs | 5.79 MB, 502 allocs | 1.96 MB, 501 allocs |

Raw bytes: 2,060,288 / 6,066,176 / 2,060,288 (within a few hundred bytes
across shapes).

1.96 MB is the heatmap itself: 500 rows of 500 ints. The BFS's extra 3.82 MB is
its queue, which holds every block once, 16 bytes a `Point`.

## When there are walls

`TestSweepsBreakWhenAWalkMustDoubleBack`, on a 5 x 5 grid with two rivers and
one bridge each at opposite ends:

```
bottom-right block: BFS 16, two sweeps -1, straight-line scan 8
```

BFS walks the zigzag and gets 16. The sweeps never reach the block. The scan
ignores the rivers and says 8.
