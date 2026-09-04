# challenge/ — 365 Days of LeetCode Challenge tooling

Supports the `/leetcode-content` Claude Code skill (`.claude/skills/leetcode-content/SKILL.md`).
Design doc: `docs/superpowers/specs/2026-08-30-leetcode-content-gen-agent-design.md`.

## Layout

- `cmd/leetcodectl/` — CLI binary: deterministic logic (matching, reorg, queue, hero
  rendering, LeetCode fetch) behind JSON-in/JSON-out subcommands.
- `internal/` — the packages behind that CLI, each independently tested
  (`slugutil`, `queue`, `gitmap`, `resolver`, `reorg`, `leetcode`, `companies`, `hero`, `cli`).
- `queue.yaml` — the ordered, resumable record of which questions have content generated
  and/or have been posted. Read/written by `leetcodectl`; don't hand-edit unless fixing a
  mistake.
- `number_folder_map.yaml` — incremental cache mapping LeetCode question numbers to repo
  folders, derived from commit history. Safe to delete; it'll rebuild (slowly, via a full
  history scan) on the next `gitmap-update`.
- `companies_dataset.json` / `companies_dataset.md` — vendored, manually-refreshed
  company-tag data (see `companies_dataset.md` for how to refresh it).
- `hero_template.html` — the HTML template rendered + screenshotted into each post's
  `HERO.png`.

## One-time setup

```bash
go build -o /tmp/leetcodectl ./challenge/cmd/leetcodectl   # or wherever you keep local tools
npx playwright install chromium                             # needed for hero-image screenshots
```

The `/leetcode-content` skill also shells out to `curl` to download question images —
it's preinstalled on macOS and virtually all Linux distros, so there's normally nothing
to do here, but it's worth knowing about if you're on a minimal environment that lacks it.

## Running the tests

```bash
go test ./challenge/...
```

## Posting

Content generation (the `/leetcode-content` skill) fills `queue.yaml`; posting reads it.
Each entry tracks posting per destination under `posted_at`, so a day can be on Discord
but not yet on LinkedIn:

```yaml
posted_at:
    discord: "2026-09-03T09:00:00Z"
    linkedin_main_account: null
    linkedin_company_page: null
    linkedin_group: null
```

`status` says nothing about posting — it only means content was generated. Adding a new
place to post means adding a key to `queue.KnownDestinations`; no migration needed.

### Discord — automatic

`.github/workflows/discord-daily-post.yml` runs daily at 03:22 UTC (08:52 IST) and posts
the oldest entry not yet on Discord, attaching that day's `HERO.png`. The message body is
the day's `POST_DISCORD.md`, sent verbatim — that file is written to be exactly what lands
in the channel, so it must stay under Discord's 2000-character limit and must not
reference SVGs (Discord can't preview them).

Setup: create an incoming webhook on the target channel (Edit Channel → Integrations →
Webhooks → New Webhook → Copy Webhook URL) and store it as the repository secret
`DISCORD_WEBHOOK_URL`. Trigger the workflow by hand once from the Actions tab before
relying on the cron.

The queue is only written after Discord accepts the message, so a failed run leaves the
day unclaimed and the next run retries it instead of skipping ahead.

GitHub's scheduled runs are best-effort — the 03:30 UTC slot once fired at 08:09 UTC — so
the workflow carries four hourly crons on an odd minute (03:22, 04:22, 05:22, 06:22 UTC)
rather than one. Doubling up is prevented in code, not by luck: a run that finds a day
already posted to Discord on the same UTC calendar day exits without posting, so whichever
cron lands first sends the day and the rest are no-ops. To post a second day on purpose
(catching up after a missed day), run the workflow by hand with the `force` input, or pass
`"force":true` locally.

Run it locally against a real or fake webhook:

```bash
DISCORD_WEBHOOK_URL=... go run ./challenge/cmd/leetcodectl post-discord \
  '{"repoRoot":".","queuePath":"challenge/queue.yaml"}'
```

### LinkedIn — manual, batch-prepared

LinkedIn's API requires product approval, rotates tokens every ~60 days, and has no
endpoint for publishing newsletter articles, so these are scheduled by hand in LinkedIn's
own composer. The tooling just prepares the content and tracks what went out.

```bash
# Write the next 7 unposted days to one paste-ready file
go run ./challenge/cmd/leetcodectl linkedin-batch \
  '{"repoRoot":".","queuePath":"challenge/queue.yaml","count":7}'

# After scheduling them, mark only the ones that actually went out
go run ./challenge/cmd/leetcodectl mark-posted \
  '{"queuePath":"challenge/queue.yaml","destination":"linkedin_main_account","numbers":[104,108]}'
```

`linkedin-batch` never marks anything itself — scheduling by hand is a partial process, and
claiming days that were never scheduled would skip them permanently. It prints the exact
`mark-posted` command to run afterwards. Valid destinations are `linkedin_main_account`,
`linkedin_company_page` and `linkedin_group`.
