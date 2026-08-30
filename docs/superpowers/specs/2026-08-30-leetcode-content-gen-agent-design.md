# LeetCode Content-Gen Agent — Design

Date: 2026-08-30

## Context

Goal is a full "365 Days of LeetCode Challenge" engagement mechanism: generate teaching
content from already-solved solutions, then post it daily to LinkedIn + Discord. That
larger goal decomposes into three independent subsystems, each with its own spec/plan:

1. **Content-gen agent** (this spec) — turns a batch of already-solved LeetCode questions
   into repo content + post drafts, in a fixed, resumable order.
2. **Daily poster (cron)** — future spec. Reads the queue this subsystem produces, posts
   one entry/day to LinkedIn + Discord on a GitHub Actions schedule.
3. **Distribution extras** — future spec. Auto-comment "join Discord" on the LinkedIn post,
   cross-post to a list of LinkedIn groups.

#2 depends on #1's queue file existing; #3 depends on #2's posting mechanism. This spec
covers #1 only.

## Invocation

```
/leetcode-content <leetcode-problem-list-url> --filter difficulty=easy --list-name "binary tree easy problems"
```

A Claude Code skill, run manually by the user once or twice a week against a topic+difficulty
problem-list URL. Designed to be re-run on the same URL/list repeatedly (e.g. this week 10 of
30 questions in the list are solved, agent processes those 10; next week the remaining 20 are
solved, a re-run with the same command picks up exactly those 20 and skips the 10 already done).

## Pipeline (per run)

1. **Fetch + filter.** Pull the problem list from the URL via LeetCode's public GraphQL API
   (no auth required for list + question statement). Filter to the requested difficulty.
2. **Sort.** Sort the filtered list by LeetCode question number ascending — gives a
   deterministic, natural "solve order" within a batch.
3. **Skip already-queued.** Drop any question whose number already has an entry in
   `challenge/queue.yaml` — this is what makes partial-batch resumption work.
4. **Resolve solved-status + location**, for each remaining question, checking candidate
   locations in this order:
   1. Canonical: `{difficulty}_problems/{100-range}/{slug}` (e.g. `easy_problems/901_1000/univalued_binary_tree`)
   2. Root straggler: `{difficulty}_problems/{slug}` (older solutions saved directly under
      the difficulty root, not in a number-range subfolder)
   3. `google_questions/*/{slug}` (old topic-organized folder)
   4. `linkedin_questions/*/{slug}` (old topic-organized folder)
   5. Git-log number map: commit messages largely follow `Type(Number) Title` (e.g.
      `Easy(965) Univalued Binary Tree`). Build/maintain an incremental
      number→folder cache at `challenge/number_folder_map.yaml`, rescanning only commits
      newer than the last scan.
   6. Slug-fuzzy fallback: convert the LeetCode slug (dashes→underscores) and search all
      difficulty roots for a matching folder name.
   7. Still unresolved → ask the user inline during the session which folder (if any)
      corresponds to this question. Not solved yet → skip silently, it'll be picked up on
      a future re-run once solved.
5. **Reorg if non-canonical.** If the question was found via 4.2–4.4 (root straggler, or in
   `google_questions`/`linkedin_questions`), `git mv` it into
   `{difficulty}_problems/{100-range}/{slug}`, normalizing the folder name to the current
   LeetCode slug. The 100-range bucket is computed from the question number
   (`bucket_start = floor((number-1)/100)*100 + 1`, e.g. 965 → `901_1000`). This preserves
   git history via `mv` and fixes existing inconsistencies opportunistically (e.g. the
   dash-named `medium_problems/1-100` vs. the underscore convention everywhere else) only
   when that folder happens to be touched by a batch — no separate bulk migration.
6. **Generate content** for each resolved, newly-solved question (see below).
7. **Append to queue** (`challenge/queue.yaml`), assigning the next global Day N.
8. **Commit** — one commit per question: `Content(<number>) <title> - Day <N>`, directly to
   `main` (matches existing repo convention of pushing solutions straight to main).

## Content generated per question

| File | Content | Notes |
|---|---|---|
| `README.md` | Problem statement, examples, constraints | Only written if missing. Any `<img>` in the statement (diagrams, matrix illustrations) is downloaded into `<folder>/images/` and the reference rewritten to the local path — the repo doesn't hotlink LeetCode's CDN |
| `INTUITION.md` | Plain-language approach + why + complexity | New file |
| `SOLUTION.md` | Narrated walkthrough of the *existing* solution code | Explains the user's actual code, doesn't rewrite it |
| `main.go` | Same logic, inline `//` comments added | Comments only, no logic changes |
| `COMPANIES.md` | Companies known to have asked this question | Only written if the vendored dataset has a hit; silently skipped otherwise |
| `HERO.png` | Branded image: "Day N / 365", topic, difficulty, question title, LeetCode logo | Rendered from `challenge/hero_template.html` via headless-browser screenshot |
| `POST_LINKEDIN.md` | Day N header, title+link, 2-3 line intuition hook, code snippet/link, companies if present | Discord-join CTA is NOT baked in here — it's added as a comment by the poster (subsystem #2), not part of the post body |
| `POST_DISCORD.md` | Shorter, Discord-markdown formatted version | Same Day N |

### Companies dataset

Vendored as a static snapshot at `challenge/companies_dataset.json`, sourced once from a
public community dataset (LeetCode company-tag data is Premium-only; no live public API
exists for it). Refreshed manually by the user occasionally — not re-fetched per run.

### Hero image

Template lives at `challenge/hero_template.html` (HTML/CSS with placeholders for day
counter, topic, difficulty, title). Rendered to PNG via a headless-browser screenshot
invoked by the agent. LeetCode logo asset vendored once at
`challenge/assets/leetcode_logo.png`.

## Data model — `challenge/queue.yaml`

Shared state file; this subsystem writes it, the future poster subsystem (#2) reads and
advances it. This is the sole handoff contract between the two.

```yaml
next_day: 24
entries:
  - day: 23
    number: 965
    title: "Univalued Binary Tree"
    difficulty: easy
    folder: easy_problems/901_1000/univalued_binary_tree
    batch: "binary tree easy problems"
    status: content_ready   # content_ready | posted
    posted_at: null
```

Entries are appended in the order resolved (ascending question number within a batch), and
batches are appended in the order the user fed them to the agent — this is what gives the
global Day N its "365 days" narrative continuity across topics.

## Data model — `challenge/number_folder_map.yaml`

Incremental cache, rebuilt by scanning new git commits since the last recorded commit SHA:

```yaml
last_scanned_commit: 8d3b1eb57b8e80ef791f1aa34afa2de040537e3a
map:
  965: easy_problems/901_1000/univalued_binary_tree
  1022: easy_problems/1001_1100/sum_of_root_to_leaf_binary_numbers
```

## Out of scope for this spec

- Daily poster / GitHub Actions cron (subsystem #2)
- LinkedIn / Discord posting APIs, auth, rate limits (subsystem #2)
- Discord-join auto-comment on LinkedIn posts, LinkedIn group cross-posting (subsystem #3)
- Bulk one-time migration of `google_questions/` / `linkedin_questions/` (explicitly deferred
  to lazy, batch-triggered reorg instead)

## Open risks / notes for implementation

- LeetCode's public GraphQL endpoint is unofficial and undocumented; schema or rate limits
  could change without notice. No auth/session cookie assumed for v1 (premium fields like
  official company tags are unavailable this way, hence the vendored dataset instead).
- Headless-browser screenshot rendering adds a non-Go runtime dependency to an otherwise
  pure-Go repo; acceptable since it only runs during the manual content-gen skill invocation,
  not in any hot path.
- Git-log commit message parsing is best-effort (some historical commits omit the question
  number entirely) — the fuzzy/manual fallbacks exist specifically to cover that gap.
