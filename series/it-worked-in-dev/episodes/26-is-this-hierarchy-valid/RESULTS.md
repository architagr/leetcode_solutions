# Measured

```
Apple M1 Pro · go1.26.4 darwin/arm64
go test -bench=. -benchtime=20x                                        (three runs, median)
go test -bench='(Children|Bounds)/corrupt_early' -benchtime=100000x    (three runs, median)
go test -run TestTheChildrenCheckAcceptsACorruptTree -v
```

Reproduce with the commands above in this directory. `corrupt_early` is decided
in a few nanoseconds by the two top-down checks, so for those two it runs at
100,000 iterations; the rows below use that figure.

| shape | products | the tree |
|---|---:|---|
| `valid_32k` | 32,767 | balanced, valid |
| `valid_262k` | 262,143 | balanced, valid |
| `corrupt_deep_262k` | 262,143 | the dearest price left of the root patched to above the root |
| `corrupt_early_262k` | 262,143 | the root's left child patched to above the root |
| `chain_5k` | 5,000 | valid, one long right spine, as a sorted import builds it |

## The check that is wrong

`TestTheChildrenCheckAcceptsACorruptTree`, on a 1,024-product tree with the
same patch as `corrupt_deep`:

```
1024 products, one price on the wrong side of the root: children check says valid
```

The three other checks reject it.

## Time

| shape | children only | every subtree | in-order, sorted | bounds |
|---|---:|---:|---:|---:|
| `valid_32k` | 126 µs | 716 µs | 414 µs | 112 µs |
| `valid_262k` | 1.01 ms | 7.21 ms | 3.14 ms | 0.90 ms |
| `corrupt_deep_262k` | 1.02 ms, says valid | 208 µs | 2.74 ms | 465 µs |
| `corrupt_early_262k` | 2.24 ns | 211 µs | 3.23 ms | 4.22 ns |
| `chain_5k` | 57.5 µs | 91.9 ms | 79.6 µs | 57.2 µs |

Raw ns, in column order:

```
125,860 / 715,519 / 414,252 / 111,669
1,008,773 / 7,210,410 / 3,142,694 / 896,912
1,018,196 / 208,342 / 2,735,052 / 465,421
2.238 / 210,796 / 3,231,531 / 4.221
57,510 / 91,927,300 / 79,606 / 57,183
```

Every subtree against bounds, same tree: **6.41x**, **8.04x**, **0.448x**,
**49940x**, **1608x**. On `corrupt_deep` the subtree check is faster: its first
act at the root is to scan the left subtree for its maximum, which finds the
patched price; bounds walks the left subtree in order and meets that price last.

`chain_5k` against `valid_32k` for the subtree check: 91.9 ms against 716 µs,
**128x**, for a seventh of the products.

In-order-and-sorted against bounds: **3.71x**, **3.5x**, **5.88x**,
**765584x**, **1.39x**. It always reads the whole tree before it can say no.

The children check against bounds: **1.13x**, **1.12x**, **2.19x**, **0.53x**,
**1.01x**. It is no faster than the correct check, and on `corrupt_deep` it
reads the whole tree and says valid.

## Allocated, same runs

Only the in-order check allocates: its list of prices. 1.11 MB at 32,767
products, 10.1 MB at 262,143, 125 KB at 5,000.

Raw bytes: 1,160,441 · 10,564,858 · 128,248
