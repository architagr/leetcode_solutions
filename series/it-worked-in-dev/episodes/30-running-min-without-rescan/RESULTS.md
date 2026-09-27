# Measured

```
Apple M1 Pro · go1.26.4 darwin/arm64
go test -bench='/(month_100|year_2k)' -benchtime=500x      (three runs, median)
go test -bench='/(decade_20k|log_100k)' -benchtime=5x      (three runs, median)
go test -run 'TestSessionSizes|TestOneMinIsWrongAfterUndo' -v
```

Reproduce with the commands above in this directory.

Each shape is a session of changes to a budget plan - mostly adding a
transaction, now and then undoing a few - with the lowest balance asked for
after every change, as the planner's overdraft warning does. One benchmark
iteration plays the whole session.

`TestSessionSizes`:

```
month_100       109 changes:     89 added,    20 undone, plan peaks at     71 transactions
year_2k        2108 changes:   1907 added,   201 undone, plan peaks at   1706 transactions
decade_20k    21143 changes:  19066 added,  2077 undone, plan peaks at  16989 transactions
log_100k     106103 changes:  94950 added, 11153 undone, plan peaks at  83797 transactions
```

`log_100k` is a small business's plan: its change log, replayed when the
planner loads, with the warning recomputed at every step for the history chart.

## The version that is wrong

`TestOneMinIsWrongAfterUndo`, with an opening balance of 1,000:

```
undo the -900: one tracked minimum says -100, the plan's lowest is now 800
```

## Time, a whole session

| shape | rescan on every change | lowest carried per entry | ratio |
|---|---:|---:|---:|
| `month_100` | 3.85 µs | 1.45 µs | 2.65x |
| `year_2k` | 1.26 ms | 25.7 µs | 49x |
| `decade_20k` | 134 ms | 261 µs | 512x |
| `log_100k` | 3.04 s | 1.37 ms | 2224x |

Raw ns: 3,848 / 1,451 · 1,257,552 / 25,684 · 133,745,058 / 261,233 ·
3,035,645,133 / 1,365,242

Per change on `log_100k`: 28.6 µs for the rescan, 13 ns for the entry.

## Allocated, same runs

| shape | rescan on every change | lowest carried per entry |
|---|---:|---:|
| `month_100` | 2.02 KB | 6.01 KB |
| `year_2k` | 38.6 KB | 124 KB |
| `decade_20k` | 645 KB | 1.86 MB |
| `log_100k` | 3.07 MB | 10.6 MB |

Raw bytes: 2,072 / 6,152 · 39,576 / 126,995 · 660,784 / 1,945,608 ·
3,216,700 / 11,128,888

Carrying two more numbers per transaction triples what the plan holds: 24
bytes an entry instead of 8.
