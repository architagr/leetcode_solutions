# Measured

```
Apple M1 Pro · go1.26.4 darwin/arm64
go test -bench=. -benchtime=200x
go test -run TestPreviewsAgreeAndCountBuilds -v
```

Reproduce with the commands above in this directory.

| shape | comments | levels | visible lines |
|---|---:|---:|---:|
| `thread_1k` | 1,365 | 6 | 6 |
| `thread_5k` | 5,461 | 7 | 7 |
| `thread_56k` | 55,987 | 7 | 7 |
| `wide_50k` | 50,001 | 2 | 2 |
| `chain_2k` | 2,001 | 2,001 | 2,001 |

`chain_2k` is the shape where the answer is the whole thread, so nothing below
helps on it. It is in every table for that reason.

## Returning the visible IDs

| shape | every level, then the last | one slot per depth | newest reply first |
|---|---:|---:|---:|
| `thread_1k` | 24.7 µs | 4.18 µs | 3.69 µs |
| `thread_5k` | 92.5 µs | 16.8 µs | 16.9 µs |
| `thread_56k` | 816 µs | 173 µs | 171 µs |
| `wide_50k` | 429 µs | 138 µs | 112 µs |
| `chain_2k` | 99.7 µs | 37.6 µs | 37.4 µs |

Raw ns, in column order:

```
24,726 / 4,182 / 3,688
92,527 / 16,822 / 16,899
816,478 / 172,914 / 170,691
429,405 / 138,465 / 111,928
99,702 / 37,614 / 37,415
```

Building every level against walking newest first, same thread:

| shape | ratio |
|---|---:|
| `thread_1k` | 6.70x |
| `thread_5k` | 5.48x |
| `thread_56k` | 4.78x |
| `wide_50k` | 3.84x |
| `chain_2k` | 2.66x |

One slot per depth against walking newest first, same thread: **1.13x**,
**1.00x**, **1.01x**, **1.24x**, **1.01x**. The two are the same function with
the children iterated in opposite directions, and when the entry is a string
that already exists, reversing buys nothing.

Allocated, same runs:

| shape | every level, then the last | one slot per depth | newest reply first |
|---|---:|---:|---:|
| `thread_1k` | 65.9 KB, 26 allocs | 240 B, 4 | 240 B, 4 |
| `thread_5k` | 296 KB, 33 allocs | 240 B, 4 | 240 B, 4 |
| `thread_56k` | 2.87 MB, 41 allocs | 240 B, 4 | 240 B, 4 |
| `wide_50k` | 1.53 MB, 7 allocs | 48 B, 2 | 48 B, 2 |
| `chain_2k` | 233 KB, 4,015 allocs | 98.4 KB, 13 | 98.4 KB, 13 |

Raw bytes: 67,464 / 240 / 240 · 303,001 / 240 / 240 · 3,005,403 / 240 / 240 ·
1,605,764 / 48 / 48 · 238,440 / 100,720 / 100,720

Building every level against walking newest first, held: **281x**, **1263x**,
**12523x**, **33453x**, **2x**.

## When the visible entry costs something to build

`Preview` formats one line: the id, the reply count, and a little string
building. It is the same cost in both versions, so the only thing that differs
is how many times it is called.

| shape | comments | previews built, last writer wins | previews built, newest first |
|---|---:|---:|---:|
| `thread_1k` | 1,365 | 1,365 | 6 |
| `thread_5k` | 5,461 | 5,461 | 7 |
| `thread_56k` | 55,987 | 55,987 | 7 |
| `wide_50k` | 50,001 | 50,001 | 2 |
| `chain_2k` | 2,001 | 2,001 | 2,001 |

| shape | last writer wins | newest reply first | ratio |
|---|---:|---:|---:|
| `thread_1k` | 57.3 µs | 3.81 µs | 15.0x |
| `thread_5k` | 244 µs | 20.4 µs | 12.0x |
| `thread_56k` | 3.88 ms | 172 µs | 22.5x |
| `wide_50k` | 3.12 ms | 116 µs | 26.8x |
| `chain_2k` | 179 µs | 176 µs | 1.02x |

Raw ns: 57,310 / 3,810 · 244,381 / 20,427 · 3,879,454 / 172,022 ·
3,115,652 / 116,485 · 179,314 / 176,120

Allocated, same runs:

| shape | last writer wins | newest reply first | ratio |
|---|---:|---:|---:|
| `thread_1k` | 64.2 KB, 1,369 allocs | 528 B, 10 | 125x |
| `thread_5k` | 256 KB, 5,465 allocs | 576 B, 11 | 456x |
| `thread_56k` | 6.77 MB, 101,979 allocs | 1.12 KB, 17 | 6165x |
| `wide_50k` | 5.95 MB, 90,005 allocs | 245 B, 6 | 25470x |
| `chain_2k` | 370 KB, 3,915 allocs | 370 KB, 3,915 | 1x |

Raw bytes: 65,760 / 528 · 262,368 / 576 · 7,102,469 / 1,152 ·
6,240,210 / 245 · 379,264 / 379,264

`chain_2k` is identical in both columns, to the byte. Every level holds one
comment, so every comment is a first arrival and the reversal has nothing to
skip.

## The version that is wrong

`TestNewestBranchMissesDeeperSiblings`, on the four-comment thread in the
diagrams:

```
following the newest reply down gives [root new]
the collapsed view actually shows     [root new old-1 old-2]
```

Two lines where there should be four. It is not benchmarked, because a wrong
answer has no time worth reporting.
