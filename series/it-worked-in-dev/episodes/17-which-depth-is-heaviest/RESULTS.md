# Measured

```
Apple M1 Pro · go1.26.4 darwin/arm64
go test -bench=. -benchtime=2000x      (three runs, median reported)
go test -run 'TestMap|TestShapes' -v
```

Reproduce with the commands above in this directory.

| shape | categories | depths | heaviest depth |
|---|---:|---:|---:|
| `store_400` | 400 | 4 | 3 |
| `store_7k` | 7,381 | 5 | 4 |
| `market_56k` | 55,987 | 7 | 6 |
| `flat_50k` | 50,001 | 2 | 1 |
| `chain_2k` | 2,001 | 2,001 | 6 |

## Does it give the same answer twice

`TestMapAnswerChangesBetweenRunsOnATie`, on the seven-category catalogue in the
diagrams, where depth 1 and depth 2 both hold 12 SKUs:

```
map version over 1,000 calls on the same catalogue: map[1:883 2:117]
```

The split moves from run to run; the fact that there is one does not. The two
ordered versions return depth 1 every time.

`TestMapOnTheChain`, where each depth holds between 1 and 7 SKUs and ties are
everywhere:

```
chain_2k, correct answer depth 6; map version over 200 calls gave 145 different depths
```

## Time

| shape | map, as written | map, explicit tie-break | slice by depth | level walk |
|---|---:|---:|---:|---:|
| `store_400` | 3.76 µs | 3.91 µs | 1.19 µs | 3.11 µs |
| `store_7k` | 58.5 µs | 59.2 µs | 22.0 µs | 57.5 µs |
| `market_56k` | 406 µs | 408 µs | 174 µs | 684 µs |
| `flat_50k` | 347 µs | 363 µs | 138 µs | 200 µs |
| `chain_2k` | 119 µs | 125 µs | 27.5 µs | 14.2 µs |

Raw ns, in column order:

```
3,761 / 3,907 / 1,188 / 3,108
58,460 / 59,154 / 22,044 / 57,535
406,249 / 408,321 / 174,243 / 684,082
346,958 / 363,279 / 138,350 / 200,308
118,754 / 125,497 / 27,549 / 14,193
```

Map as written against slice by depth, same catalogue: **3.17x**, **2.65x**,
**2.33x**, **2.51x**, **4.31x**.

Map with the explicit tie-break against slice by depth, same catalogue:
**3.29x**, **2.68x**, **2.34x**, **2.63x**, **4.56x**. Fixing the tie-break
costs the map nothing measurable; it is the map that costs.

Level walk against slice by depth, same catalogue: the level walk is slower by
**2.62x**, **2.61x**, **3.93x** and **1.45x** on the four bushy shapes, and
faster by **1.94x** on `chain_2k`, where the slice version recurses 2,001 frames
deep and the level walk does not recurse at all.

Map as written against the level walk on `chain_2k`: **8.37x**.

## Allocated, same runs

| shape | map, as written | map, explicit tie-break | slice by depth | level walk |
|---|---:|---:|---:|---:|
| `store_400` | 0 B, 0 allocs | 0 B, 0 | 56 B, 3 | 10.0 KB, 11 |
| `store_7k` | 0 B, 0 allocs | 0 B, 0 | 120 B, 4 | 195 KB, 21 |
| `market_56k` | 0 B, 0 allocs | 0 B, 0 | 120 B, 4 | 2.11 MB, 37 |
| `flat_50k` | 0 B, 0 allocs | 0 B, 0 | 24 B, 2 | 392 KB, 1 |
| `chain_2k` | 145 KB, 29 allocs | 145 KB, 29 | 58.6 KB, 14 | 8 B, 1 |

Raw bytes: 0 / 0 / 56 / 10,240 · 0 / 0 / 120 / 199,392 · 0 / 0 / 120 / 2,213,029 ·
0 / 0 / 24 / 401,408 · 148,152 / 148,152 / 60,024 / 8

The map versions allocate nothing on the bushy shapes because a map of eight
or fewer small keys fits in the space Go reserves on the stack for it. Its cost
there is hashing, not memory. The level walk holds a whole level of pointers at
a time, which on `market_56k` is the 46,656 leaves.
