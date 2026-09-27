# Measured

```
Apple M1 Pro · go1.26.4 darwin/arm64
go test -bench='ByList|ExportAll|ByStack/(page_5000|last_page)' -benchtime=30x      (three runs, median)
go test -bench='PageAfter|ByStack/(page_1$|page_100$)' -benchtime=20000x            (three runs, median)
go test -run TestStackStaysShallow -v
```

Reproduce with the commands above in this directory. Anything under about 20
µs runs at 20,000 iterations, because at 30 it is noise.

`big` is 200,000 products, prices 10 to 2,000,000 in steps of 10, inserted in
random order. Pages are 20 products. There are 10,000 pages.

`TestStackStaysShallow`:

```
200000 prices walked in order; the stack never held more than 31 nodes
```

## Time, one page

| shape | build the list, slice it | stack iterator, skip to page | stack iterator from a cursor |
|---|---:|---:|---:|
| `page_1` | 6.52 ms | 377 ns | 397 ns |
| `page_100` | 5.98 ms | 12.1 µs | 419 ns |
| `page_5000` | 6.12 ms | 1.29 ms | 723 ns |
| `last_page` | 6.12 ms | 2.94 ms | 254 ns |

Raw ns, in column order:

```
6,523,229 / 376.9 / 396.7
5,980,599 / 12,088 / 419.4
6,120,528 / 1,293,378 / 723.2
6,123,276 / 2,937,347 / 254.4
```

Building the list against the cursor, same page: **16444x**, **14260x**,
**8463x**, **24069x**.

Skipping to the page against the cursor, same page: **0.95x** on page 1 (the
same work, and within noise), then **28.8x**, **1788x**, **11546x**.

Building the list against skipping to the page: **17308x**, **495x**, **4.73x**,
**2.08x**. Skipping still steps through every product before the page, so on
the last page it is only twice as fast as building the list.

## Time, every product

The export: all 200,000 products, in order.

| | build the list, then read it | stack iterator |
|---|---:|---:|
| export | 6.26 ms | 2.60 ms |

Raw ns: 6,260,760 / 2,601,946. Building the list is **2.41x** slower even when
every product is wanted.

## Allocated, same runs

| | build the list | stack iterator, either way |
|---|---:|---:|
| one page | 7.98 MB, 33 allocs | 216 to 664 B, 4 to 7 allocs |
| export | 7.98 MB, 32 allocs | 504 B, 6 allocs |

Raw bytes: 8,369,605 / 408 on `page_1`; 8,369,443 / 504 on the export. The
stack iterator's allocations are its stack and the page slice.
