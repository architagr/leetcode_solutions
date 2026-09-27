# Measured

```
Apple M1 Pro · go1.26.4 darwin/arm64
go test -bench=. -benchtime=20x                     (three runs, median)
go test -bench='/week_hours' -benchtime=20000x      (three runs, median)
go test -run TestSteps -v
```

Reproduce with the commands above in this directory. `week_hours` is fast
enough that 20 iterations is noise; its rows use the 20,000-iteration run.

`daily` series have a daily cycle, a weekly swing, a slow upward trend and
noise. `launch` series are the traffic after a launch: a big first few weeks
decaying to a steady base, with the same daily cycle and noise.

`TestSteps`, how far the forward scan has to look in total:

```
week_hours         168 periods,     13 never beaten  |  forward steps the scan takes:         1409
year_hours        8760 periods,     12 never beaten  |  forward steps the scan takes:        93709
year_minutes    525600 periods,     75 never beaten  |  forward steps the scan takes:     46381027
launch_hours      8760 periods,    143 never beaten  |  forward steps the scan takes:      1103549
launch_minutes   43200 periods,    603 never beaten  |  forward steps the scan takes:     17086892
```

`year_hours` and `launch_hours` are both 8,760 hours. The launch year takes the
scan 11.8 times as many steps.

## Time

| shape | periods | scan forward | monotonic stack | ratio |
|---|---:|---:|---:|---:|
| `week_hours` | 168 | 1.18 µs | 441 ns | 2.68x |
| `year_hours` | 8,760 | 95.9 µs | 26.2 µs | 3.66x |
| `year_minutes` | 525,600 | 18.0 ms | 3.98 ms | 4.53x |
| `launch_hours` | 8,760 | 386 µs | 34.9 µs | 11.1x |
| `launch_minutes` | 43,200 | 5.87 ms | 327 µs | 17.9x |

Raw ns: 1,180 / 440.7 · 95,933 / 26,242 · 18,049,004 / 3,982,619 ·
386,127 / 34,927 · 5,865,477 / 326,971

`year_hours` against `launch_hours`, same number of hours: the scan takes
**4.02x** as long on the launch year; the stack **1.33x**.

Per period: the scan costs 11 ns on `year_hours` and 136 ns on
`launch_minutes`; the stack costs 3 to 8 ns everywhere.

## Allocated, same runs

Both allocate the result: one int per period. The stack adds its own slice,
which grows past its initial 64 entries only on the minute-level series (a few
KB).

Raw bytes: 73,728 / 73,728 on the hourly years; 4,210,701 / 4,217,856 on
`year_minutes`; 352,256 / 359,424 on `launch_minutes`.
