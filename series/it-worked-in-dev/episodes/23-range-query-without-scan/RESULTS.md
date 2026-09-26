# Measured

```
Apple M1 Pro · go1.26.4 darwin/arm64
go test -bench=. -benchtime=300x      (three runs, median reported)
go test -run TestVisits -v
```

Reproduce with the commands above in this directory.

`cat` is 200,000 products with distinct prices 0, 10, 20 ... 1,999,990 paise,
inserted in random order: a binary search tree 44 levels deep. `imported` is
20,000 products inserted in increasing price order, the way a bulk import from
a sorted CSV builds it: a chain 20,000 deep.

`TestVisits`:

```
narrow_20        depth     44,     20 in range  |  nodes visited: scan  200000, pruning      43
band_2k          depth     44,   2000 in range  |  nodes visited: scan  200000, pruning    2021
wide_100k        depth     44, 100000 in range  |  nodes visited: scan  200000, pruning  100015
everything       depth     44, 200000 in range  |  nodes visited: scan  200000, pruning  200000
import_low_20    depth  20000,     20 in range  |  nodes visited: scan   20000, pruning      20
import_high_20   depth  20000,     20 in range  |  nodes visited: scan   20000, pruning   20000
```

## Time

| shape | in range | scan and filter | pruned walk | ratio |
|---|---:|---:|---:|---:|
| `narrow_20` | 20 | 2.33 ms | 357 ns | 6517x |
| `band_2k` | 2,000 | 2.40 ms | 32.4 µs | 74x |
| `wide_100k` | 100,000 | 4.09 ms | 2.90 ms | 1.41x |
| `everything` | 200,000 | 5.87 ms | 6.27 ms | 0.937x |
| `import_low_20` | 20 | 0.20 ms | 286 ns | 710x |
| `import_high_20` | 20 | 0.20 ms | 167 µs | 1.21x |

Raw ns: 2,327,739 / 357.2 · 2,397,919 / 32,393 · 4,088,227 / 2,900,003 ·
5,872,954 / 6,270,529 · 203,243 / 286.2 · 201,053 / 166,565

The ratio is scan against pruned walk, same tree and range. Below 1, the scan
is faster: on `everything`, where every node is in range, pruning only adds two
comparisons per node.

## Allocated, same runs

Both versions allocate the same: the result slice growing. 504 B for 20
results, 58.6 KB for 2,000, 3.91 MB for 100,000, 7.98 MB for 200,000.

Raw bytes: 504 · 60,024 · 4,101,373 · 8,369,420 (within a few bytes between
the two versions).
