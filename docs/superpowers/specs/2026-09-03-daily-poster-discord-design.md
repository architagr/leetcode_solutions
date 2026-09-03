# Daily Poster (Discord) — Design

Date: 2026-09-03

## Context

Subsystem #2 of the "365 Days of LeetCode Challenge" mechanism, as decomposed in
`2026-08-30-leetcode-content-gen-agent-design.md`:

1. **Content-gen agent** — done. Turns already-solved questions into repo content + post
   drafts, in a fixed resumable order, tracked in `challenge/queue.yaml`.
2. **Daily poster (cron)** — this spec. Reads that queue and posts one entry per day.
3. **Distribution extras** — future. Auto-comment "join Discord" on LinkedIn, group
   cross-posting.

Subsystem #2 is itself two independent integrations. This spec covers **Discord only**.
LinkedIn gets its own spec, plan and build cycle afterwards, reusing the same queue
schema and CLI conventions established here.

## Goals

- One Discord message per day, posted automatically, in strict Day order.
- Never post the same day twice; never silently skip a day.
- Track posting state per destination, so LinkedIn (and anything after it) plugs into the
  same queue without a schema change.

## Queue schema change

`posted_at` becomes a map keyed by destination, replacing the single nullable timestamp:

```yaml
- day: 1
  number: 104
  title: Maximum Depth of Binary Tree
  difficulty: easy
  folder: easy_problems/101_200/maximum_depth_of_binary_tree
  batch: binary tree
  status: content_ready
  posted_at:
    discord: 2026-09-03T09:00:00Z
    linkedin_main_account: null
    linkedin_company_page: null
    linkedin_group: null
```

Rules:

- A destination key that is **absent or null** means "not yet posted there". A destination
  is never inferred from `status`.
- Keys are lowercase snake_case, matching every other field in this file.
- Adding a future destination means writing a new key. No migration, no code change in the
  `queue` package.
- `status` no longer tracks posting at all. It means only "content has been generated"
  (`content_ready`). Posting state lives entirely in `posted_at`. This removes the
  ambiguity of a single `status` field trying to describe N destinations.
- `queue.Load` accepts the old shape (`posted_at: null` or a bare timestamp string) and
  normalizes it to the new map form in memory, so existing entries keep working and get
  rewritten on the next save.

## Discord's constraints (these drive the content format)

Two hard limits shaped this design:

- **2000 characters** per message `content`. Several existing `POST_DISCORD.md` files
  already exceed this (2169–2511 chars), and one solution file alone
  (`kth_largest_element_in_a_stream/min.go`, 2819 chars) is over the entire budget before
  any prose.
- **SVG doesn't render.** Discord previews PNG/JPG/GIF/WEBP only. The walkthrough diagrams
  this project generates are SVG, so they cannot appear inline in a Discord message no
  matter how they're referenced.

So `POST_DISCORD.md` is redefined to be the **exact text of the Discord message**: short,
under 1900 characters, no SVG embeds, ending in a link to the full walkthrough on GitHub.
The poster posts the file verbatim — it does no markdown parsing, no section extraction,
no code re-derivation. One file, one message, no transformation in between.

The content-gen skill (`.claude/skills/leetcode-content/SKILL.md`, step m) is updated with
these constraints, and all 28 existing `POST_DISCORD.md` files are rewritten to match.

## Message format

```
**365 Days of LeetCode Challenge — Day <day>/365**
**<title>** (<difficulty>)
🔗 <leetcode url>

<2-3 line intuition>

```go
<the solution, or its core function if the whole file won't fit>
```

Full walkthrough with step-by-step diagrams: <github blob url to SOLUTION.md>
```

`HERO.png` is attached to the message as a file and referenced by the embed as
`attachment://HERO.png` — these images live in the repo, not on a public CDN, so there is
no URL to point an embed at.

When the full solution won't fit the budget, the code block carries only the core function
(dropping `TreeNode` boilerplate, imports and comments); if even that doesn't fit, the code
block is dropped entirely and the GitHub link carries the weight.

## Architecture

Extends the existing `leetcodectl` CLI rather than introducing a second tool or a
throwaway script, so this reuses the `queue` package and follows the same JSON-in/JSON-out
subcommand pattern, with the same test discipline, as everything else under `challenge/`.

- **`challenge/internal/discordpost`** (new package)
  - `SelectNext(q *queue.Queue, destination string) (queue.Entry, bool)` — the oldest entry
    with no timestamp for that destination.
  - `BuildPayload(entry queue.Entry, messageText string) Payload` — the Discord webhook
    JSON body.
  - `Post(ctx, webhookURL string, payload Payload, heroPNGPath string) error` — multipart
    POST (JSON payload + attached PNG). `BaseURL`-style injection so tests can point it at
    an `httptest.Server`.
- **`challenge/internal/queue`** (extended)
  - `Entry.PostedAt` becomes `map[string]*string`.
  - `MarkPosted(number int, destination string, at time.Time)`.
  - `Load` normalizes the legacy shape.
- **`leetcodectl post-discord`** (new subcommand) — JSON arg
  `{"repoRoot":".","queuePath":"challenge/queue.yaml","repoSlug":"architagr/leetcode_solutions","branch":"main"}`.
  The webhook URL is read from the `DISCORD_WEBHOOK_URL` environment variable, never from
  the JSON argument, so it can't end up in a logged command line.
- **`.github/workflows/discord-daily-post.yml`** — daily cron, builds and runs the
  subcommand, then commits the updated `queue.yaml` back to the repo.

## Data flow

1. Load `queue.yaml`.
2. Select the oldest entry whose `posted_at.discord` is absent or null.
3. If there is none, log "nothing to post" and exit 0. An idle day is not a failure.
4. Read `<folder>/POST_DISCORD.md` verbatim as the message text. Read `<folder>/HERO.png`
   as the attachment.
5. POST to the webhook as multipart form data.
6. On 2xx: set `posted_at.discord` to the current UTC time (RFC3339), save `queue.yaml`,
   exit 0.
7. The workflow commits the changed `queue.yaml` (`chore(challenge): mark Day N posted to
   Discord`) and pushes.

## Error handling

- **Webhook returns non-2xx, or the request fails** — return the error, exit non-zero,
  leave `queue.yaml` untouched. The Actions run goes red and the same day is retried on the
  next run. No double-post, no skipped day.
- **Message exceeds 2000 chars** — fail before sending, naming the file and its length.
  This is a content bug to fix in `POST_DISCORD.md`, not something to paper over by
  truncating mid-sentence at post time.
- **`POST_DISCORD.md` or `HERO.png` missing for the selected entry** — hard error, exit
  non-zero. Do not skip ahead to the next entry: that would silently reorder the Day
  sequence, which is exactly the failure mode the queue exists to prevent.
- **Rate limited (429)** — surface Discord's `retry_after` in the error message. A daily
  single-message cadence should never hit this; if it does, something is wrong and it
  should be loud.

## Testing

- `SelectNext` — table-driven: empty queue, all posted, mixed destinations, ordering by day
  rather than file order, legacy `posted_at: null` entries.
- `queue.MarkPosted` and legacy-shape normalization — save/load round trips.
- Message-length validation — at, just under, and just over the limit.
- `Post` — `httptest.Server` asserting the multipart body carries both the JSON payload and
  the PNG, plus non-2xx and 429 handling.
- The GitHub Actions workflow itself is verified by a manual `workflow_dispatch` run before
  the cron is relied on.

## Secrets

One GitHub Actions repository secret: `DISCORD_WEBHOOK_URL`, generated from the target
channel's own settings (Server Settings → Integrations → Webhooks). A webhook URL is a
bearer credential for that one channel — anyone holding it can post there — so it is never
committed, echoed, or passed as a CLI argument.

## Out of scope

- LinkedIn posting (its own spec, next).
- The "join Discord" auto-comment and LinkedIn group cross-posting (subsystem #3).
- Editing or deleting already-posted messages.
- Backfilling Discord posts for Days 1–28 in bulk. The cron picks them up one per day, in
  order, which is the intended cadence.
