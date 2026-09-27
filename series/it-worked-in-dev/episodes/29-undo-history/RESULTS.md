# Measured

```
Apple M1 Pro · go1.26.4 darwin/arm64
go test -bench=. -benchtime=20x      (three runs, median reported)
go test -run TestHeld -v
```

Reproduce with the commands above in this directory.

Each shape is a writing session on a starting document: mostly typing a word,
sometimes deleting a few characters, now and then pasting a paragraph. One
benchmark iteration plays the whole session and then undoes every edit.

`TestHeld`, the history each editor holds at the end of the session, before any
undo:

```
note_2k    document    8567 bytes after   500 edits  |  history held: snapshots     2662792 bytes, operations   19735 bytes
doc_50k    document   78321 bytes after  2000 edits  |  history held: snapshots   127702511 bytes, operations   80623 bytes
spec_200k  document  214155 bytes after  1000 edits  |  history held: snapshots   206876404 bytes, operations   40661 bytes
```

| shape | document at the end | snapshots held | operations held | ratio |
|---|---:|---:|---:|---:|
| `note_2k` | 8.37 KB | 2.54 MB | 19.3 KB | 135x |
| `doc_50k` | 76.5 KB | 122 MB | 78.7 KB | 1584x |
| `spec_200k` | 209 KB | 197 MB | 39.7 KB | 5088x |

## Time, a whole session and every undo

| shape | snapshots | operations, copying the document | operations, in place |
|---|---:|---:|---:|
| `note_2k` | 680 µs | 1.04 ms | 89.6 µs |
| `doc_50k` | 9.51 ms | 38.6 ms | 2.05 ms |
| `spec_200k` | 12.7 ms | 56.2 ms | 3.63 ms |

Raw ns, in column order:

```
680,400 / 1,041,915 / 89,588
9,505,319 / 38,550,019 / 2,053,015
12,729,919 / 56,247,425 / 3,633,115
```

Snapshots against operations in place, same session: **7.59x**, **4.63x**,
**3.5x**.

Operations that copy the document on every edit against snapshots: **1.53x**,
**4.06x**, **4.42x** slower than snapshots. An operation stack that still
copies the document does two copies per edit - once to make the edit, once to
undo it - where the snapshot editor's undo swaps a pointer.

## Allocated over the session

| shape | snapshots | operations, copying | operations, in place |
|---|---:|---:|---:|
| `note_2k` | 2.78 MB | 5.53 MB | 134 KB |
| `doc_50k` | 130 MB | 259 MB | 739 KB |
| `spec_200k` | 201 MB | 403 MB | 650 KB |

Raw bytes: 2,917,042 / 5,802,508 / 137,408 · 135,942,744 / 271,702,059 / 757,020 ·
211,161,170 / 422,065,937 / 665,345
