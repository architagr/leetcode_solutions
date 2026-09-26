# Measured

```
Apple M1 Pro · go1.26.4 darwin/arm64
go test -bench=. -benchtime=50x      (three runs, median reported)
go test -run TestShapes -v
```

Reproduce with the commands above in this directory.

The four `origins_*` shapes are the same 10,000-node gossip cluster (each node
peers with one earlier node, plus 10,000 random peerings), with the origins
spread evenly across it. `ring_2k_20` is 2,000 nodes in a ring with 20 origins.

`TestShapes`:

```
origins_1       10000 nodes,     1 origins,   9 rounds  |  nodes visited: search each      10000, one wave  10000
origins_10      10000 nodes,    10 origins,   9 rounds  |  nodes visited: search each     100000, one wave  10000
origins_100     10000 nodes,   100 origins,   7 rounds  |  nodes visited: search each    1000000, one wave  10000
origins_1000    10000 nodes,  1000 origins,   6 rounds  |  nodes visited: search each   10000000, one wave  10000
ring_2k_20       2000 nodes,    20 origins,  50 rounds  |  nodes visited: search each      40000, one wave   2000
```

More origins, fewer rounds: the invalidation itself gets faster as the origins
go up.

## Time

| shape | origins | rounds | search from each | one wave | ratio |
|---|---:|---:|---:|---:|---:|
| `origins_1` | 1 | 9 | 0.50 ms | 196 µs | 2.54x |
| `origins_10` | 10 | 9 | 2.81 ms | 197 µs | 14.2x |
| `origins_100` | 100 | 7 | 26.7 ms | 198 µs | 135x |
| `origins_1000` | 1,000 | 6 | 269 ms | 200 µs | 1350x |
| `ring_2k_20` | 20 | 50 | 0.36 ms | 8.76 µs | 40.6x |

Raw ns: 499,208 / 196,275 · 2,806,441 / 197,028 · 26,743,015 / 197,672 ·
269,397,452 / 199,597 · 355,043 / 8,755

The ratio column is search from each origin against one wave, same cluster and
same origins.

`origins_1` against `origins_1000`, same cluster: search from each goes from
0.50 ms to 269 ms, **540x**. One wave goes from 196 µs to 200 µs, **1.02x**.

With a single origin the two do the same search, and search from each is still
**2.54x** slower: `HopsFrom` grows its queue from empty (20 allocations against
2), and the minimum-then-maximum pass over the result is a second walk over
every node.

## Allocated, same runs

| shape | search from each | one wave |
|---|---:|---:|
| `origins_1` | 509 KB, 20 allocs | 160 KB, 2 |
| `origins_10` | 4.27 MB, 191 allocs | 160 KB, 2 |
| `origins_100` | 42.0 MB, 1,902 allocs | 160 KB, 2 |
| `origins_1000` | 419 MB, 19,010 allocs | 160 KB, 2 |
| `ring_2k_20` | 1.47 MB, 281 allocs | 32.0 KB, 2 |

Raw bytes: 521,456 / 163,840 · 4,477,404 / 163,840 · 44,035,763 / 163,840 ·
439,619,008 / 163,840 · 1,544,390 / 32,768
