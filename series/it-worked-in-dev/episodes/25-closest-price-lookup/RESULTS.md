# Measured

```
Apple M1 Pro · go1.26.4 darwin/arm64
go test -bench='Sorting|SortedFlat|Walk' -benchtime=30x     (three runs, median)
go test -bench=Descent -benchtime=200000x                   (three runs, median)
```

Reproduce with the commands above in this directory.

`small` is 1,000 products and `big` 200,000, with distinct prices spaced 1 to
50 paise apart, inserted in random order. Each shape asks for the price
nearest a target 3 paise above an existing price, in the cheap end, the middle
or the dear end of the catalogue.

The descent is fast enough that 30 iterations is noise, so it runs at 200,000.
It asks the same question every iteration, so its path through the tree is in
cache: in a 20-iteration run, with the path cold, it measured 60 to 148 ns.

## Time

| shape | sort, then search | flat sort, then search | walk until past | descent |
|---|---:|---:|---:|---:|
| `1k_middle` | 55.2 µs | 9.78 µs | 2.78 µs | 12.4 ns |
| `200k_cheap` | 36,355 µs | 5,893 µs | 20.2 µs | 16.2 ns |
| `200k_middle` | 32,785 µs | 6,924 µs | 1,290 µs | 17.8 ns |
| `200k_dear` | 32,761 µs | 6,022 µs | 2,651 µs | 19.6 ns |

Raw ns, in column order:

```
55,179 / 9,781 / 2,781 / 12.44
36,354,804 / 5,892,599 / 20,172 / 16.2
32,784,507 / 6,924,496 / 1,290,333 / 17.84
32,761,185 / 6,021,900 / 2,650,786 / 19.59
```

Sort-then-search against the flat sort, same query: **5.64x**, **6.17x**,
**4.73x**, **5.44x**. That is the cost of building the list by concatenation.

Sort-then-search against the descent: **4436x**, **2244124x**, **1837697x**,
**1672342x**.

Flat sort against the descent: **786x**, **363741x**, **388144x**, **307397x**.

The walk against the descent: **224x**, **1245x**, **72328x**, **135313x**. The
walk stops once it passes the target, so it is cheap for a cheap target and
costs a walk of most of the catalogue for a dear one.

## Allocated, same runs

| shape | sort, then search | flat sort, then search | walk until past | descent |
|---|---:|---:|---:|---:|
| `1k_middle` | 134 KB, 1,430 allocs | 24.6 KB, 12 | 0 B, 0 | 0 B, 0 |
| 200k shapes | 51.5 MB, 290,046 allocs | 7.98 MB, 31 | 0 B, 0 | 0 B, 0 |

Raw bytes: 136,976 / 25,208 / 0 / 0 · 53,953,265 / 8,369,419 / 0 / 0
(the 200k shapes are within a few hundred bytes of each other).
