# Measured

```
Apple M1 Pro · go1.26.4 darwin/arm64
go test -bench=. -benchtime=50x -count=3              (median of three)
go test -run TestSteps -v
```

Reproduce with the commands above in this directory. Both versions prune in
place, so every iteration prunes a fresh copy; the copy is made with the timer
stopped.

The configs are built in `prune_test.go`. `service` is one service's config,
5,000 entries, at most eight children per section and six levels, with 30% of
the keys still read. `platform` is every service's config in one tree, 500,000
entries and eight levels, also 30% read. `stale` is the same tree after years:
0.2% of keys still read. `rules` is a pricing rule chain 5,000 rules deep: each
rule has a condition, a price and an `else` holding the next rule, and only the
fallback at the bottom is still read.

`TestSteps`, how many entries the "is anything under here read?" questions read
in total:

```
service      5000 entries,    2075 kept  |  entries the questions read:       8211
platform   500000 entries,  201512 kept  |  entries the questions read:     767765
stale      500000 entries,    3467 kept  |  entries the questions read:     769488
rules       15004 entries,    5004 kept  |  entries the questions read:   37527506
```

On the three random configs the questions read each entry about 1.5 times. On
`rules` they read each entry 2,501 times.

## Time

| shape | entries | ask each section | bottom-up | ratio |
|---|---:|---:|---:|---:|
| `service` | 5,000 | 46.7 µs | 39.6 µs | 1.18x |
| `platform` | 500,000 | 5.35 ms | 5.42 ms | 0.99x |
| `stale` | 500,000 | 3.59 ms | 3.77 ms | 0.95x |
| `rules` | 15,004 | 104 ms | 92.4 µs | 1,130x |

Raw ns: 46,738 / 39,596 · 5,354,858 / 5,422,041 · 3,593,892 / 3,766,257 ·
104,376,445 / 92,358

On the stale platform config, asking is **1.05x** faster than bottom-up: almost
nothing is read, so the first question at the root reads everything once and
every dead section is dropped without being visited again. On `rules`,
bottom-up is **1,130x** faster.

`rules` against `platform`, asking each section: 15,004 entries take **19.5x**
as long as 500,000.

## Allocated, same runs

Neither allocates. Both reuse each section's own `Kids` slice for the
survivors.
