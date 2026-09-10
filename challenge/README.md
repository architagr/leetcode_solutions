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
  `HERO.png` by `leetcodectl hero-generate`. The intermediate HTML goes to a temp file and
  is deleted, so `HERO.png` is the only hero artifact in a solution folder.

  The card's background gradient rotates through `hero.Palettes` (eight dark themes) by
  day, so two posts in a row don't read as the same post sent twice. It's derived from the
  day rather than drawn at random: random can repeat back-to-back, and re-rendering a day
  would then produce a different image than the one already published. Everything that
  does the branding — the `#ffa116` accent, the logos, the type, the layout — is fixed
  across every palette, and `TestPalettesKeepBrandTextReadable` holds each one to WCAG AA
  contrast for the card's text, so a new palette can't quietly make a card unreadable.

## Commit messages

Commits here carry no AI attribution: no `Co-Authored-By: Claude ...` trailer, no
"Generated with Claude Code" line, no 🤖. The work publishes under Archit Agarwal's name on
his own channels, and the commit log is part of that.

`.githooks/commit-msg` enforces it — a message mentioning Claude, Anthropic or "generated
with" fails the commit. It's versioned rather than left in `.git/hooks` so it can be
reviewed and travels with the repo, but Git won't wire it up on a fresh clone by itself:

```bash
git config core.hooksPath .githooks    # once per clone
```

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
    x: null
    substack: null
```

Two of these post themselves and two don't, which is a property of the networks rather
than a choice: Discord and X have APIs that will accept a post from a cron, LinkedIn's
needs product approval and rotates tokens every ~60 days, and Substack has no publishing
API at all.

| Destination | How | Command |
|---|---|---|
| `discord` | automatic, daily cron | `post-discord` |
| `x` | automatic, daily cron | `post-x` |
| `linkedin_*` | manual, batch-prepared | `linkedin-batch` + `mark-posted` |
| `substack` | manual, batch-prepared | `substack-batch` + `mark-posted` |

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

### X — automatic

`.github/workflows/x-daily-post.yml` runs daily at 04:22 UTC (09:52 IST), an hour after
Discord's, and posts the oldest entry not yet on X with that day's `HERO.png` attached.
The message body is the day's `POST_X.md`, sent verbatim — that file is written to be
exactly what lands on the timeline, so it must stay under 280 characters, and `post-x`
refuses an over-length post rather than truncating one.

Everything runs through the official X API v2 with OAuth 1.0a user credentials. That is
worth stating plainly, because the "will this get my account limited" question has a
boring answer: X does not run an AI-text detector on posts, and there is nothing to
evade. What its rules actually police is *platform manipulation* — duplicate or
near-duplicate posts, bulk automated follows/likes/replies, and unauthorised automation
(browser drivers, scraped sessions). So the defences that matter here are structural, and
all four are already in place:

- one post a day, guarded by the same-day check in `PostX`, not by cron luck;
- a `concurrency` group, so two runs can never post seconds apart;
- `POST_X.md` written per day as its own copy, never a truncation of the Discord or
  LinkedIn text, so the same sentence isn't fanned out across four networks;
- no automated engagement of any kind — this posts, and does nothing else.

#### Setup

X Premium (the consumer subscription) does **not** include API access; the developer
project is separate and free at the tier this needs.

1. At `developer.x.com`, create a project and an app inside it.
2. In the app's **User authentication settings**, set App permissions to **Read and
   write**. This is the step people miss — tokens minted before this is set stay
   read-only, and posting fails with a 401 until they're regenerated.
3. From **Keys and tokens**, take the API Key and Secret (the consumer pair) and generate
   an Access Token and Secret (the user pair). All four are needed: posting is a
   user-context action, so an app-only bearer token cannot do it.
4. Store them as repository secrets: `X_API_KEY`, `X_API_SECRET`, `X_ACCESS_TOKEN`,
   `X_ACCESS_TOKEN_SECRET`. Set the handle as a repository *variable* `X_HANDLE` — it
   only builds the result URL, it authorises nothing.
5. Check them before trusting them. `verify-x` reads back the authenticated account and
   publishes nothing:

   ```bash
   set -a; source ~/path/to/your/x-keys.env; set +a   # never inside this repo
   go run ./challenge/cmd/leetcodectl verify-x
   ```

   A green `verify-x` and a 401 from `post-x` means one specific thing: the tokens were
   generated before App permissions were set to Read and write. Only a write reveals that,
   so regenerate the access token and secret and try again.
6. Trigger the workflow by hand once from the Actions tab before relying on the cron.

Only the OAuth 1.0a values are used. The Bearer Token is app-only auth and cannot post as
you; the OAuth 2.0 client credentials drive a redirect flow whose access tokens expire every
two hours, which is the wrong shape for an unattended daily job. OAuth 1.0a user tokens do
not expire.

Keep the keys out of this working tree. `.gitignore` covers `x.com`, `.env` and `*.secrets`,
but the only durable answer is to keep the file somewhere else entirely and source it.

There is no free tier. X retired it on 2026-02-06 and moved every account to pay-per-use
credits (Basic migrated 2026-06-01, Pro 2026-09-01), so a project with no credit balance
fails on its very first write with a 402 rather than after some allowance runs out. Add a
payment method and load credits in the portal before the first run.

The pricing has one sharp edge worth knowing before writing the posts rather than after:

| Request | Cost |
|---|---|
| Post, plain text | ~$0.015 |
| Post containing any URL | ~$0.20 |
| Media upload | bundled into the post |

A link costs 13x a plain post — X is deliberately taxing off-platform links, and its
ranking demotes them too. Every day's post carrying a link to `SOLUTION.md` is therefore
~$73/year and gets less reach than the same post without one; dropping the link is ~$5/year
and travels further, at the cost of the click-through. That is a content decision rather
than a technical one, so nothing here enforces either way — but it is the reason to make
the decision deliberately.

Run it locally:

```bash
X_API_KEY=... X_API_SECRET=... X_ACCESS_TOKEN=... X_ACCESS_TOKEN_SECRET=... \
  go run ./challenge/cmd/leetcodectl post-x \
  '{"repoRoot":".","queuePath":"challenge/queue.yaml","handle":"architagr"}'
```

As with Discord, the queue is only written after X accepts the post, so a failed run
leaves the day unclaimed and the next run retries it instead of skipping ahead.

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

The batch file lifts each article's YAML front matter into its own "publish settings"
block — meta title and meta description — because LinkedIn asks for those in separate
fields at publish time rather than reading them off the article. Days whose
metadata is missing or too long still prepare; the gaps come back in `warnings` and are
listed at the bottom of the batch file. Articles written before the front matter existed
have none, and that is fine.

### Substack — manual, batch-prepared

Substack has no public publishing API — only inbound RSS and email import — so there is
nothing to automate against even in principle. The tooling prepares the post and tracks
what went out; the paste is yours.

```bash
go run ./challenge/cmd/leetcodectl substack-batch \
  '{"repoRoot":".","queuePath":"challenge/queue.yaml","count":4}'
```

Images are referenced by local path (`images/<file>`), the same as in `SOLUTION.md` and the
LinkedIn article — they get attached by hand when the post is pasted in, so the path is a
pointer for the author rather than something Substack resolves.

The document carries each day's `POST_SUBSTACK.md` with its publish settings lifted out,
the same way `linkedin-batch` does — including the post's `tags`, which are Substack's
discovery mechanism and are set in the publish dialog rather than written into the body.
Substack does not index inline hashtags, so `POST_SUBSTACK.md` carries none. Mark what actually went out:

```bash
go run ./challenge/cmd/leetcodectl mark-posted \
  '{"queuePath":"challenge/queue.yaml","destination":"substack","numbers":[104,108]}'
```

There is no canonical URL field, because neither platform lets a publisher set one:
LinkedIn's composer has none, and Substack treats its own domain as canonical while
exposing only an SEO title and subtitle. The piece does go out in both places, so the only
lever available is the one both posts use — a prominent link back to the repo in the body.
