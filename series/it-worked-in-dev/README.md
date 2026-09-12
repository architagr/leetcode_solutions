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
| 1 | [Why `du` crawls on node_modules](episodes/01-directory-sizes/) | postorder traversal | days 1, 4, 9 |

## Why it lives in this repo

Every episode backlinks the days it is built on, so keeping them together makes
those links local. One clone gets you the solutions and the benchmarks the
article just told you to run, and `tools/schedule.py` reads the day queue two
directories up rather than trusting a separate checkout to be current.

## Layout

```
FORMAT.md     what the series is, and the rules an episode follows
queue.yaml    the publishing queue, and which days each episode needs
LINKS.yaml    canonical URL per surface
tools/        diagram, hero, humanize, lint_episode, render_post, schedule
episodes/     one folder each: code, tests, benchmarks, diagrams, drafts
```

## Working on an episode

Run everything from this directory.

```bash
cd episodes/01-directory-sizes
go test ./...                      # both implementations agree
go test -bench=. -benchtime=200x   # the numbers in RESULTS.md

cd ../..
python3 tools/lint_episode.py            # structure, metadata, units, images
python3 tools/checklinks.py              # every link accounted for
python3 tools/hero.py --check-palettes   # backgrounds distinct and readable
python3 tools/schedule.py                # what is publishable, and what is blocked
python3 tools/render_post.py episodes/01-directory-sizes   # paste-ready copy
```

An episode may only publish once **every day it cites has actually posted**. A
reader following a backlink has to land on something they can read, so
`schedule.py` checks that and refuses to call an episode publishable until then.

`FORMAT.md` has the rules. `.claude/skills/worked-in-dev-episode/SKILL.md` at
the repo root has the full working guide.
