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
