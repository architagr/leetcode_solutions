# Measured

```
Apple M1 Pro · go1.26.4 darwin/arm64
go test -bench=. -benchtime=20x      (three runs, median reported)
go test -run TestVisits -v
```

Reproduce with the commands above in this directory. `NetworksByReach` takes
over two seconds a call on the two big shapes, so a full run is a few minutes.

| shape | services | links | networks |
|---|---:|---:|---:|
| `isolated_5k` | 5,000 | 0 | 5,000 |
| `teams_5k` | 5,000 | 5,000 | 500 |
| `joined_5k` | 5,000 | 5,499 | 1 |
| `mesh_500` | 500 | 998 | 1 |
| `mesh_5k` | 5,000 | 9,998 | 1 |

`joined_5k` is `teams_5k` with one extra link from each team to the next: the
same services, the same clusters, and 499 more links.

## Services visited

`TestVisits`:

```
isolated_5k    5000 services,  5000 networks  |  services visited: reach per service        5000, one sweep   5000
teams_5k       5000 services,   500 networks  |  services visited: reach per service       50000, one sweep   5000
joined_5k      5000 services,     1 networks  |  services visited: reach per service    25000000, one sweep   5000
mesh_500        500 services,     1 networks  |  services visited: reach per service      250000, one sweep    500
mesh_5k        5000 services,     1 networks  |  services visited: reach per service    25000000, one sweep   5000
```

Asking every service for its reach visits k services k times in a network of
k, so the total is the sum of the squares of the network sizes. One sweep
visits every service once whatever the shape.

## Time

| shape | reach per service | one sweep | ratio |
|---|---:|---:|---:|
| `isolated_5k` | 0.61 ms | 11.7 µs | 51.7x |
| `teams_5k` | 2.82 ms | 22.7 µs | 124x |
| `joined_5k` | 1,715 ms | 23.1 µs | 74175x |
| `mesh_500` | 23.7 ms | 3.92 µs | 6051x |
| `mesh_5k` | 2,513 ms | 90.7 µs | 27696x |

Raw ns: 606,904 / 11,740 · 2,819,054 / 22,658 · 1,714,851,190 / 23,119 ·
23,724,279 / 3,921 · 2,513,095,419 / 90,740

The ratio column is reach per service against one sweep, same mesh.

`teams_5k` against `joined_5k`, same 5,000 services: reach per service goes
from 2.82 ms to 1.71 s, **608x**. One sweep goes from 22.7 µs to 23.1 µs,
**1.02x**.

## Allocated, same runs

| shape | reach per service | one sweep |
|---|---:|---:|
| `isolated_5k` | 328 KB, 5,046 allocs | 45.3 KB, 2 |
| `teams_5k` | 2.52 MB, 25,015 allocs | 45.3 KB, 2 |
| `joined_5k` | 1.97 GB, 295,131 allocs | 45.3 KB, 2 |
| `mesh_500` | 21.7 MB, 11,001 allocs | 4.50 KB, 2 |
| `mesh_5k` | 1.97 GB, 295,128 allocs | 45.3 KB, 2 |

Raw bytes: 335,928 / 46,336 · 2,637,599 / 46,336 · 2,120,614,900 / 46,336 ·
22,724,140 / 4,608 · 2,120,612,807 / 46,336

That is 1.97 GB allocated per call, not held: each `Reachable` returns its set
and the next call throws it away. The garbage collector is doing a good share
of the 1.71 seconds.
