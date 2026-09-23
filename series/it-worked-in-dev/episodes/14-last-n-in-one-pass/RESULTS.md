# Measured

```
Apple M1 Pro · go1.26.4 darwin/arm64
go test -bench=. -benchtime=20x
go test -run TestRetainedHeap -v
```

Reproduce with the commands above in this directory. `n` is 200 everywhere -
the last 200 lines of the log.

A log line here is 68 bytes, the shape of a real one: a timestamp, a level, a
request id, a method, a path, a status and a duration.

## Live heap with the answer still held

This is the number the episode is about. `TestRetainedHeap` forces a GC,
reads `HeapAlloc`, runs the function, forces another GC and reads it again with
the result still referenced, so what it reports is what the collector cannot
take back. Three runs, median.

| shape | lines | collecting | collecting and copying | ring |
|---|---:|---:|---:|---:|
| the log in front of you | 200 | 21.8 KB | 20.5 KB | 20.5 KB |
| one chatty request | 10,000 | 959 KB | 20.5 KB | 20.5 KB |
| a pod since its last restart | 1,000,000 | 93.2 MB | 20.6 KB | 21.0 KB |
| a rotation period | 2,000,000 | 186 MB | 20.8 KB | 22.3 KB |

Raw bytes: 22,368 / 20,960 / 20,960 · 981,728 / 20,960 / 20,960 ·
97,753,664 / 21,088 / 21,520 · 194,702,944 / 21,296 / 22,832

Collecting against the ring, same input:

| shape | lines | ratio |
|---|---:|---:|
| the log in front of you | 200 | 1.1x |
| one chatty request | 10,000 | 46.8x |
| a pod since its last restart | 1,000,000 | 4542x |
| a rotation period | 2,000,000 | 8528x |

The 200-line row is 1.1x, which is to say the same. At that size the collected
slice *is* the answer.

## Why collecting retains that much after returning 200 lines

`all[len(all)-n:]` is a window onto the array `all` grew into, and a slice
keeps its whole backing array alive. `TestCollectedTailPinsTheWholeLog` reads
the capacity off the returned value rather than arguing about it:

| lines | len of the answer | cap of the answer |
|---:|---:|---:|
| 100,000 | 200 | 17,960 |
| 1,000,000 | 200 | 109,704 |
| 2,000,000 | 200 | 169,032 |

The capacity is what is left of the array past the window. Everything before
the window is in the same allocation, so it is scanned and kept too.

## Time

| shape | lines | collecting | copying | counting twice | ring |
|---|---:|---:|---:|---:|---:|
| the log in front of you | 200 | 34.0 µs | 33.6 µs | 61.3 µs | 30.7 µs |
| one chatty request | 10,000 | 2.20 ms | 1.78 ms | 3.36 ms | 1.65 ms |
| a pod since its last restart | 1,000,000 | 182 ms | 177 ms | 331 ms | 166 ms |
| a rotation period | 2,000,000 | 356 ms | 359 ms | 658 ms | 329 ms |

Raw ns, in column order:

```
33,954 / 33,552 / 61,292 / 30,656
2,195,108 / 1,780,335 / 3,355,638 / 1,651,600
181,902,317 / 177,168,646 / 331,457,629 / 166,050,977
355,502,315 / 359,271,575 / 658,353,156 / 328,733,175
```

Collecting against the ring, same input: **1.11x**, **1.33x**, **1.10x**,
**1.08x**. This is not a speed fix, and the episode says so.

Counting twice against the ring, same input: **2.00x**, **2.03x**, **2.00x**,
**2.00x**. Two passes cost two passes, at every size.

## Allocated against retained

`B/op` from the same benchmark, which counts every byte allocated rather than
every byte kept:

| shape | lines | collecting, B/op | ring, B/op | ratio |
|---|---:|---:|---:|---:|
| a pod since its last restart | 1,000,000 | 168 MB | 84.0 MB | 1.91x |
| a rotation period | 2,000,000 | 333 MB | 168 MB | 1.98x |

Raw bytes: 176,021,006 / 88,035,778 · 348,676,597 / 176,065,965

Both versions allocate every line, because the source produces every line
either way. Allocation differs by **1.98x**. Retention differs by **8528x**.
They are not the same measurement and only one of them is what fills a
container.

## The same question on a chain already in memory

`NthFromEndByCounting` against `NthFromEndByGap`, n=200, zero allocations for
both:

| entries | counting the chain first | a fixed gap | ratio |
|---:|---:|---:|---:|
| 1,000 | 1.77 µs | 1.05 µs | 1.69x |
| 100,000 | 246 µs | 183 µs | 1.34x |
| 1,000,000 | 3.77 ms | 1.66 ms | 2.27x |

Raw ns: 1,769 / 1,046 · 245,600 / 183,048 · 3,774,704 / 1,662,010

The counting version walks the chain once to measure it and then walks most of
it again. The gap version walks it once.
