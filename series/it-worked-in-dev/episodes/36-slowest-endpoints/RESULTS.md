# Measured

```
Apple M1 Pro · go1.26.4 darwin/arm64
go test -bench=. -benchtime=10x -count=3                        (median of three)
go test -bench='/(page|minute)' -benchtime=2000x -count=3       (median of three)
go test -run TestEntries -v
```

Reproduce with the commands above in this directory. `page` and `minute` are
fast enough that 10 iterations is noise; their rows use the 2,000-iteration
run. Every run asks for the ten slowest.

The logs are built in `slowest_test.go`. Latency is log-normal around 40 ms -
most requests fast, a long slow tail - over 300 routes. `page` is the last
1,000 requests; `minute` is 60,000, a minute at 1,000 a second; `hour` is
3,600,000. `backlog` is an hour in which a queue backs up, so each request waits
10 µs longer than the one before it on top of its own time. `ascending` is an
hour exported already ordered fastest first. `climbing` is the extreme: every
request strictly slower than the one before.

`TestEntries`, how many requests get past the comparison with the weakest of
the ten kept:

```
page           1000 requests  |  entered the top 10:        61 (6.1000%)
minute        60000 requests  |  entered the top 10:        95 (0.1583%)
hour        3600000 requests  |  entered the top 10:       142 (0.0039%)
backlog     3600000 requests  |  entered the top 10:      7751 (0.2153%)
ascending   3600000 requests  |  entered the top 10:   1135595 (31.5443%)
climbing    3600000 requests  |  entered the top 10:   3600000 (100.0000%)
```

`ascending` is under 100% because many requests share a latency, and a request
has to be strictly slower than the weakest kept to get in.

## Time

| shape | requests | sort a copy | heap of 10 | ratio |
|---|---:|---:|---:|---:|
| `page` | 1,000 | 52.5 µs | 1.93 µs | 27.2x |
| `minute` | 60,000 | 6.13 ms | 47.5 µs | 129x |
| `hour` | 3,600,000 | 417 ms | 2.72 ms | 154x |
| `backlog` | 3,600,000 | 387 ms | 2.86 ms | 135x |
| `ascending` | 3,600,000 | 52.0 ms | 22.7 ms | 2.30x |
| `climbing` | 3,600,000 | 20.8 ms | 87.5 ms | 0.24x |

Raw ns: 52,496 / 1,933 · 6,131,530 / 47,540 · 417,361,488 / 2,715,767 ·
386,748,750 / 2,857,508 · 52,020,071 / 22,663,079 · 20,812,492 / 87,467,108

On `climbing` the heap is **4.20x** slower than the sort: every request gets in
and costs a `heap.Fix`, while the standard library's pattern-defeating
quicksort recognises input that is already ordered. The sort itself is
**8.02x** faster on `ascending` than on `hour`, and **20.1x** faster on
`climbing`, for the same 3,600,000 requests.

## Allocated, same runs

| shape | sort a copy | heap of 10 |
|---|---:|---:|
| `page` | 24.6 KB | 504 B |
| `minute` | 1.44 MB | 504 B |
| every hour | 86.4 MB | 504 B |

Raw bytes: 24,576 · 1,441,792 · 86,401,024; the heap is 504 every time. The
sort copies the whole log, 24 bytes a request, because the log is shared. The
heap holds ten requests and the boxing `container/heap` does on each push.
