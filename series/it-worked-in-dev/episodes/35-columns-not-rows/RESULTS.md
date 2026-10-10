# Measured

```
Apple M1 Pro · go1.26.4 darwin/arm64
go test -bench=. -benchtime=20x -count=3               (median of three)
go test -bench=/stump -benchtime=20000x -count=3       (median of three)
go test -run 'TestSortWork|TestUnstable' -v
```

Reproduce with the commands above in this directory. `stump` is fast enough
that 20 iterations is noise; its rows use the 20,000-iteration run.

The trees are built in `columns_test.go`. `stump` is a complete tree six levels
deep, 63 nodes, what the model viewer opens on. `model` is complete to depth 14,
16,383 nodes. `grown` is an unpruned tree grown by inserting 200,000 random keys,
lopsided. `deep` is complete to depth 20, 1,048,575 nodes.

The four versions:

- `unstable` - note each node's column and depth, sort by both. Wrong: two
  nodes in the same row and column can come out in either order.
- `stable` - the same with `sort.SliceStable`. Correct.
- `three keys` - `sort.Slice` with the walk's order as a third key. Correct.
- `by rows` - a breadth-first walk appending each node to its column. Correct.

`TestUnstableSortPassesSmallFailsLarge`: the `sort.Slice` version matches on the
7-node tree and is wrong on all 9 complete trees of depth 4 to 12.

`TestSortWork`, what the stable sort does:

```
stump        63 nodes  |  stable sort:        333 comparisons ( 5 per node),         302 swaps (  5 per node)
model     16383 nodes  |  stable sort:     158440 comparisons (10 per node),      451644 swaps ( 28 per node)
grown    200000 nodes  |  stable sort:    2479508 comparisons (12 per node),     9900572 swaps ( 50 per node)
deep    1048575 nodes  |  stable sort:   10426684 comparisons (10 per node),    50338018 swaps ( 48 per node)
```

## Time

| shape | nodes | unstable | stable | three keys | by rows | stable ÷ rows |
|---|---:|---:|---:|---:|---:|---:|
| `stump` | 63 | 3.85 µs | 6.32 µs | 4.72 µs | 2.14 µs | 2.95x |
| `model` | 16,383 | 977 µs | 6.11 ms | 2.77 ms | 488 µs | 12.5x |
| `grown` | 200,000 | 23.8 ms | 126 ms | 46.1 ms | 9.19 ms | 13.7x |
| `deep` | 1,048,575 | 71.5 ms | 580 ms | 222 ms | 30.7 ms | 18.9x |

Raw ns, unstable / stable / three keys / rows: 3,845 / 6,315 / 4,717 / 2,140 ·
976,840 / 6,110,050 / 2,774,715 / 488,096 ·
23,761,856 / 125,785,225 / 46,098,581 / 9,190,023 ·
71,508,258 / 580,111,310 / 222,271,392 / 30,690,740

Fixing the tie with `SliceStable` costs **1.64x** on `stump`, **6.25x** on
`model`, **5.29x** on `grown` and **8.11x** on `deep`, against the wrong
`sort.Slice` version. Fixing it with a third key instead costs **2.20x** to
**7.24x** against `by rows`; the stable sort is **2.61x** slower than the
three-key sort on `deep`. `by rows` is faster than even the wrong sort:
**2.33x** on `deep`.

## Allocated, same runs

| shape | unstable | stable | three keys | by rows |
|---|---:|---:|---:|---:|
| `stump` | 5.12 KB | 5.12 KB | 6.14 KB | 4.36 KB |
| `model` | 2.34 MB | 2.34 MB | 2.68 MB | 1.59 MB |
| `grown` | 35.1 MB | 35.1 MB | 39.7 MB | 24.8 MB |
| `deep` | 179 MB | 179 MB | 242 MB | 129 MB |

Raw bytes on `deep`: 179,065,296 / 179,065,296 / 241,871,944 / 129,355,296.
Every version allocates its result; the sorts add a record per node, and
`by rows` a queue entry per node.
