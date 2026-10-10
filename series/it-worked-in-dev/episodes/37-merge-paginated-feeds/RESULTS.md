# Measured

```
Apple M1 Pro · go1.26.4 darwin/arm64
go test -bench='/(heavy|export)' -benchtime=5x -count=3       (median of three)
go test -bench='/(new|typical)' -benchtime=200x -count=3      (median of three)
go test -run TestWork -v
```

Reproduce with the commands above in this directory. The fold takes three
seconds an iteration on `heavy`, so the two big shapes run five iterations.

The feeds are built in `merge_test.go`. Every followed account returns its
latest 100 posts, newest first. Accounts post at very different rates - a few
every few minutes, most a few times a week - so the newest posts come from a
handful of feeds. `new` follows 10 accounts, `typical` 300 and `heavy` 5,000;
each asks for the first page of 50. `export` follows 300 and asks for all
30,000, the whole merge.

The three versions:

- `fold` - merge the feeds in two at a time: the first two, then the result
  with the third, and so on.
- `sort` - append every feed to one slice, sort it, take the page.
- `heap` - a heap of the k fronts, newest at the root; take the root, move that
  feed along, stop when the page is full.

`TestWork`, how many posts the fold copies in total:

```
new         10 feeds,    1000 posts, page   50  |  fold copies        5500 posts
typical    300 feeds,   30000 posts, page   50  |  fold copies     4515000 posts
heavy     5000 feeds,  500000 posts, page   50  |  fold copies  1250250000 posts
export     300 feeds,   30000 posts, page 30000  |  fold copies     4515000 posts
```

On `heavy` the fold copies 1.25 billion posts to merge 500,000: the first feed
is copied 5,000 times.

## Time

| shape | feeds | page | fold | sort | heap | fold ÷ heap |
|---|---:|---:|---:|---:|---:|---:|
| `new` | 10 | 50 | 19.2 µs | 42.5 µs | 1.44 µs | 13.3x |
| `typical` | 300 | 50 | 11.5 ms | 2.50 ms | 6.94 µs | 1,658x |
| `heavy` | 5,000 | 50 | 3.03 s | 55.1 ms | 130 µs | 23,272x |
| `export` | 300 | 30,000 | 11.8 ms | 2.55 ms | 3.05 ms | 3.87x |

Raw ns, fold / sort / heap: 19,223 / 42,491 / 1,444 ·
11,513,077 / 2,504,159 / 6,944 · 3,034,069,558 / 55,116,533 / 130,375 ·
11,797,625 / 2,546,700 / 3,045,533

On `new` the fold is **2.21x** faster than the sort. From 300 feeds it is slower:
**4.60x** on `typical`, **55.0x** on `heavy`. The heap against the sort:
**29.4x** on `new`, **361x** on `typical`, **423x** on `heavy` - and on `export`,
where the whole merge is wanted, the sort is **1.20x** faster than the heap.

## Allocated, same runs

| shape | fold | sort | heap |
|---|---:|---:|---:|
| `new` | 92.9 KB | 48.1 KB | 1.10 KB |
| `typical` | 73.4 MB | 2.66 MB | 5.81 KB |
| `heavy` | 20.0 GB | 44.7 MB | 82.9 KB |
| `export` | 73.4 MB | 2.66 MB | 493 KB |

Raw bytes: 92,928 / 48,128 / 1,104 · 73,411,918 / 2,657,281 / 5,808 ·
20,024,273,936 / 44,731,440 / 82,864 · 73,411,884 / 2,657,280 / 493,040

The fold allocates a new list per merge, each as long as everything merged so
far. The heap allocates one cursor per feed and the page.
