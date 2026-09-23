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

Episodes are numbered in the order they unlock, which is the order the daily
challenge teaches what each one is built on. So episode 8 is the first one
written, not the first one numbered.

| # | Episode | Technique | Built on |
|---|---|---|---|
| 1 | [Why your CSV import crawls on a real customer's file](episodes/01-dedupe-user-list/) | set membership | — |
| 2 | [Your org chart is fine. Your comment thread is not.](episodes/02-org-chart-depth/) | level order traversal | day 1 |
| 3 | [Why adding pricing tiers made every request slower](episodes/03-pricing-tier-lookup/) | binary search over ranges | day 2 |
| 4 | [Why the breadcrumbs cost more than the file list](episodes/04-file-breadcrumbs/) | root-to-leaf path building | day 3 |
| 5 | [Why your nav health check slowed down as the menu grew](episodes/05-lopsided-categories/) | two facts from one pass | days 1, 5 |
| 6 | [The answer was two clicks away. We read all 5,464 pages.](episodes/06-sitemap-click-depth/) | stop at the first level that answers | days 1, 7 |
| 7 | [61 comments took longer to sort than 259](episodes/07-comment-thread/) | preorder traversal | days 3, 8 |
| 8 | [Why `du` crawls on node_modules](episodes/08-directory-sizes/) | postorder traversal | days 4, 9 |
| 9 | [601 menu items took 2.4 milliseconds to flatten](episodes/09-flatten-nav-config/) | writing into one destination | days 8, 10 |
| 10 | [Your fleet tripled and the check got 12x slower](episodes/10-two-furthest-services/) | diameter, from two child results at each node | days 11, 13 |
| 11 | [The cascade delete was fine until a reply got a reply](episodes/11-safe-delete-order/) | peeling leaves, one layer at a time | days 13, 16 |
| 12 | [The loop check allocated 4.7 MB to return a boolean](episodes/12-pipeline-never-finishes/) | two cursors at different speeds | days 19, 21 |
| 13 | [Your feed got slower when one source went quiet](episodes/13-merge-two-sorted-feeds/) | merging two sorted inputs | days 19, 22 |
| 14 | [Your crash report held the whole log to send 200 lines](episodes/14-last-n-in-one-pass/) | a fixed gap between two cursors | days 20, 24 |
| 15 | [Your org chart read 55,987 rows to draw the first 43](episodes/15-org-chart-by-level/) | level order with a queue | days 27, 29 |
| 16 | [The collapsed thread built 55,987 previews to show 7](episodes/16-collapsed-tree-view/) | first arrival wins | days 29, 31 |

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
