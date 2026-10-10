# Measured

```
Apple M1 Pro · go1.26.4 darwin/arm64
go test -bench=. -benchtime=100x -count=3              (median of three)
go test -bench=/page -benchtime=20000x -count=3        (median of three)
go test -run TestSteps -v
```

Reproduce with the commands above in this directory. `page` is fast enough
that 100 iterations is noise; its rows use the 20,000-iteration run.

The trees are built in `subtree_test.go`. `page` is a random 5,000-component
page, and the snippet is the last piece of it with 40 to 60 components, so the
walk has to get most of the way through before it finds it. `grid` is 2,000
product cards from one 60-component template, and the snippet is a new variant
of the card with its last component changed. `sections` is the same with 200
sections of 2,000 components. `thread` is one reply chain 5,000 comments deep,
each comment an avatar, a body and the next reply; the snippet is the new
render of the last 1,000 comments with the bottom one marked edited. Only
`page` finds its snippet.

`TestSteps`, how many component pairs the walk compares:

```
page     page    5000, snippet    59, anchors    447  |  nodes the walk compares:        513
grid     page  120001, snippet    60, anchors   2000  |  nodes the walk compares:     120000
sections page  400001, snippet  2000, anchors    200  |  nodes the walk compares:     400000
thread   page   15001, snippet  3001, anchors   5000  |  nodes the walk compares:   13503500
```

On `grid` and `sections` the walk compares each component once, because no card
is inside another card. On `thread` every comment is inside every comment above
it: 13.5 million comparisons for 15,001 components, 900 per component.

## Time

| shape | components | walk | text search | fingerprints | walk ÷ prints |
|---|---:|---:|---:|---:|---:|
| `page` | 5,000 | 37.6 µs | 155 µs | 71.3 µs | 0.53x |
| `grid` | 120,001 | 817 µs | 2.31 ms | 868 µs | 0.94x |
| `sections` | 400,001 | 5.55 ms | 8.28 ms | 3.72 ms | 1.49x |
| `thread` | 15,001 | 75.1 ms | 468 µs | 163 µs | 461x |

Raw ns, walk / text / prints: 37,622 / 154,517 / 71,297 ·
817,045 / 2,314,260 / 867,988 · 5,547,538 / 8,280,726 / 3,719,740 ·
75,113,604 / 468,268 / 162,946

On `page` the fingerprints are **1.90x** slower than the walk, which stops when
it finds the snippet and compares 513 pairs; the fingerprints read the whole
page every time. On `grid` the walk is **1.06x** faster. On `thread` the
fingerprints are **461x** faster than the walk and **2.87x** faster than writing
the page out; the walk is **160x** slower than the text search.

`thread` against `sections`, walk only: 15,001 components take **13.5x** as long
as 400,001.

## Allocated, same runs

The walk and the fingerprints allocate nothing. Writing out builds both trees
as strings:

| shape | text search |
|---|---:|
| `page` | 114 KB |
| `grid` | 3.24 MB |
| `sections` | 10.6 MB |
| `thread` | 628 KB |

Raw bytes: 114,416 · 3,243,760 · 10,637,561 · 628,208
