# Measured

```
Apple M1 Pro · go1.26.4 darwin/arm64
go test -bench=ByLinkCheck -benchtime=5x                   (three runs, median)
go test -bench='ByCount|ByUnionFind' -benchtime=2000x      (three runs, median)
go test -run 'TestSearches|TestTwoParents' -v
```

Reproduce with the commands above in this directory. The link check takes
about two seconds a call at 5,000 components, so it runs at 5 iterations; the
other two are fast enough that 5 iterations is noise, so they run at 2,000.

| shape | components | links | a tree? |
|---|---:|---:|---|
| `tree_100` | 100 | 99 | yes |
| `tree_1k` | 1,000 | 999 | yes |
| `tree_5k` | 5,000 | 4,999 | yes |
| `loop_5k` | 5,000 | 5,000 | no: one extra link, on the last line |
| `swapped_5k` | 5,000 | 4,999 | no: right count, but the last link duplicates the first |

The links in every tree are shuffled, the way a config file grows by hand.

## Searches

`TestSearches`:

```
tree_100     n=   100 links=    99  valid=true   link check runs at least    198 full searches
tree_1k      n=  1000 links=   999  valid=true   link check runs at least   1998 full searches
tree_5k      n=  5000 links=  4999  valid=true   link check runs at least   9998 full searches
loop_5k      n=  5000 links=  5000  valid=false  link check runs at least   5000 full searches
swapped_5k   n=  5000 links=  4999  valid=false  link check runs at least   4999 full searches
```

Each of those searches rebuilds the adjacency list from the links it is given
before it walks anything.

## Time

| shape | link check | count, then one BFS | union-find |
|---|---:|---:|---:|
| `tree_100` | 0.99 ms | 7.97 µs | 620 ns |
| `tree_1k` | 68.2 ms | 49.8 µs | 6.51 µs |
| `tree_5k` | 1,955 ms | 279 µs | 42.2 µs |
| `loop_5k` | 561 ms | 2.42 ns | 41.6 µs |
| `swapped_5k` | 559 ms | 273 µs | 41.7 µs |

Raw ns, in column order:

```
988,292 / 7,971 / 619.9
68,213,100 / 49,836 / 6,506
1,954,961,917 / 279,190 / 42,218
560,890,017 / 2.417 / 41,635
559,299,667 / 272,895 / 41,651
```

`tree_1k` against `tree_5k`: 5x the components, and the link check takes
**28.7x** as long. Union-find takes **6.49x** as long.

Link check against union-find, same config: **1594x**, **10485x**, **46306x**,
**13472x**, **13428x**.

Link check against count-then-BFS on the valid trees: **124x**, **1369x**,
**7002x**; on `swapped_5k`, **2050x**.

Count-then-BFS against union-find: **12.9x**, **7.66x**, **6.61x** on the valid
trees, **6.55x** on `swapped_5k`. Union-find is the faster of the two everywhere
except `loop_5k`, where the count version rejects on `len(links) != n-1` in
2.42 ns and union-find reads the file: union-find is **17226x** slower there.

## Allocated, same runs

| shape | link check | count, then one BFS | union-find |
|---|---:|---:|---:|
| `tree_100` | 1.10 MB, 28,638 allocs | 6.34 KB, 186 | 896 B, 1 |
| `tree_1k` | 112 MB, 2,774,899 allocs | 60.5 KB, 1,818 | 8.00 KB, 1 |
| `tree_5k` | 2.74 GB, 69,097,617 allocs | 303 KB, 9,062 | 40.0 KB, 1 |
| `loop_5k` | 912 MB, 23,729,941 allocs | 0 B, 0 | 40.0 KB, 1 |
| `swapped_5k` | 912 MB, 23,734,178 allocs | 304 KB, 9,092 | 40.0 KB, 1 |

Raw bytes: 1,156,313 / 6,496 / 896 · 116,978,704 / 61,984 / 8,192 ·
2,946,262,009 / 310,452 / 40,960 · 956,174,646 / 0 / 40,960 ·
956,469,459 / 311,217 / 40,960

The count version's allocations are its adjacency list, one small slice per
component. Union-find's single allocation is its `root` slice.

## What none of them catch

`TestTwoParentsPassesTheUndirectedCheck`, a parent -> child config where
component 2 names both 0 and 1 as its parent:

```
count says tree=true, union-find says -1 (-1 is a tree)
```

Three components, two links, connected, no loop: a tree, undirected. As a
hierarchy it has two roots.
