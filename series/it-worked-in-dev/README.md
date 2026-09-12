# It worked in dev

Long-form companion to the 365-day LeetCode challenge in this repo.

The daily posts teach a technique. These answer the question the daily posts do
not, which is *when does any of this matter to me*. Each episode takes a problem
a working developer has actually hit, writes the version you would genuinely
write first, benchmarks it until it stops being fine, and derives the data
structure from the symptom rather than announcing it.

Every episode ships a runnable benchmark with real numbers on named hardware,
and the naive version is one a competent person would defend in review. That is
the whole differentiator: not the explanation, the bill arriving.

## Episodes

| # | Episode | Technique | Built on |
|---|---|---|---|
| 1 | [Why `du` crawls on node_modules](episodes/08-directory-sizes/) | postorder traversal | days 1, 4, 9 |

## Why it lives in this repo

Every episode backlinks the days it is built on, so keeping them together makes
those links local. One clone gets you the solutions and the benchmarks the
article just told you to run, and `tools/schedule.py` reads the day queue two
directories up rather than trusting a separate checkout to be current.

## Layout

```
FORMAT.md     what the series is, and the rules an episode follows
episodes/     one folder each: code, tests, benchmarks, diagrams, the write-up
```

## Running an episode

```bash
cd episodes/08-directory-sizes
go test ./...                      # both implementations agree
go test -bench=. -benchtime=200x   # the numbers in RESULTS.md
```

`RESULTS.md` records the machine and the Go version, because a benchmark number
without them is decoration. If your numbers differ from the ones in the
write-up, yours are the true ones for your machine - the argument is about the
shape of the curve, not the absolute figures.

## What is not here

The social drafts, the hero cards, the publishing queue and the tooling that
builds them are kept in a private repo. Nothing a reader would follow a link to
lives there - the write-ups, the code, the tests and the diagrams are all here,
and that is deliberate: every episode ends by telling you to clone this and make
the tests pass.

`FORMAT.md` explains what the series is and the rules each episode follows.
