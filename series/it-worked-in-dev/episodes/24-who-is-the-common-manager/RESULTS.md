# Measured

```
Apple M1 Pro · go1.26.4 darwin/arm64
go test -bench='ByPaths|ByCount|Number$' -benchtime=30x     (three runs, median)
go test -bench=ByNumbers -benchtime=20000x                  (three runs, median)
```

Reproduce with the commands above in this directory. The numbered version is
fast enough that 30 iterations is noise, so it runs at 20,000.

`company` is 50,000 employees, each manager with 3 to 8 direct reports, filled
level by level. The attendees are chosen at random with fixed seeds, except
`team_20`, which is 20 people from inside one manager's reporting line.

| shape | attendees |
|---|---:|
| `pair` | 2 |
| `team_20` | 20 |
| `allhands_200` | 200 |
| `allhands_2000` | 2,000 |

## Time

| shape | a breadcrumb per attendee | one bottom-up count | numbered chart |
|---|---:|---:|---:|
| `pair` | 0.22 ms | 0.32 ms | 22.8 ns |
| `team_20` | 2.62 ms | 0.38 ms | 253 ns |
| `allhands_200` | 17.2 ms | 1.00 ms | 1.87 µs |
| `allhands_2000` | 168 ms | 0.63 ms | 24.3 µs |

Raw ns, in column order:

```
216,267 / 316,738 / 22.77
2,621,826 / 378,729 / 252.9
17,215,733 / 1,004,361 / 1,869
167,868,197 / 633,664 / 24,283
```

Breadcrumbs against the bottom-up count, same attendees: **0.683x** for a pair
(the breadcrumbs are faster), then **6.92x**, **17.1x**, **265x**.

The bottom-up count against the numbered chart: **13910x**, **1498x**,
**537x**, **26.1x**.

Numbering the chart, once per change to it: **2.16 ms**, 2.26 MB. Against the
bottom-up count at `team_20`, that is paid back after six questions.

## Allocated, same runs

| shape | a breadcrumb per attendee | one bottom-up count | numbered chart |
|---|---:|---:|---:|
| `pair` | 648 B, 28 allocs | 0 B, 0 | 0 B, 0 |
| `team_20` | 5.43 KB, 254 allocs | 936 B, 5 | 0 B, 0 |
| `allhands_200` | 59.8 KB, 2,698 allocs | 9.13 KB, 11 | 0 B, 0 |
| `allhands_2000` | 612 KB, 27,404 allocs | 145 KB, 29 | 0 B, 0 |

Raw bytes: 648 / 0 / 0 · 5,560 / 936 / 0 · 61,272 / 9,352 / 0 ·
627,152 / 148,152 / 0

The count's allocations are its set of attendee IDs.
