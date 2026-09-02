---
name: leetcode-content
description: 'Generate teaching content (intuition/solution docs, commented code, companies, hero image, LinkedIn/Discord post drafts) for already-solved LeetCode questions from a problem-list URL, in resumable "365 Days of LeetCode Challenge" order. Trigger: /leetcode-content'
---

# LeetCode Content-Gen Agent

## Invocation

```
/leetcode-content <leetcode-problem-list-url> --filter difficulty=<easy|medium|hard> --list-name "<batch name>"
```

Design doc: `docs/superpowers/specs/2026-08-30-leetcode-content-gen-agent-design.md`

All `leetcodectl` calls below assume it's built once per session:

```bash
go build -o /tmp/leetcodectl ./challenge/cmd/leetcodectl
```

Every subcommand takes one JSON argument and prints one JSON result to stdout, e.g.:

```bash
/tmp/leetcodectl resolve '{"repoRoot":".","mapPath":"challenge/number_folder_map.yaml","number":965,"difficulty":"easy","slug":"univalued-binary-tree"}'
```

## Steps

1. **Refresh the git-history cache, and commit the refresh immediately, unconditionally —
   don't wait for the first per-question commit in step 3 to sweep it up.** If every
   question in this run turns out to be already-queued or not-yet-solved, the loop in
   step 3 never reaches a commit at all, and a rescanned `number_folder_map.yaml` would
   otherwise sit as an uncommitted working-tree change with nothing to notice it:
   ```bash
   /tmp/leetcodectl gitmap-update '{"repoRoot":".","mapPath":"challenge/number_folder_map.yaml"}'
   git add challenge/number_folder_map.yaml
   git diff --cached --quiet challenge/number_folder_map.yaml || git commit -m "chore(challenge): refresh git-history cache"
   ```

2. **Fetch the problem list:**
   ```bash
   /tmp/leetcodectl fetch-list '{"url":"<the URL>","difficulty":"<the difficulty value parsed out of --filter, e.g. --filter difficulty=easy -> \"easy\">"}'
   ```
   Sort the returned problems by `Number` ascending.

3. **For each problem, in that sorted order:**

   If any `leetcodectl` call in this loop (other than the two failure modes explicitly
   handled below — unresolved `resolve`, failed `hero-screenshot`) exits non-zero or
   returns unparseable output: stop processing THIS problem only, note it in the
   end-of-run summary with the error message, and continue to the next problem in the
   list. Never let one problem's failure abort the whole batch, and never silently
   proceed past a failed call as if it had succeeded (e.g. don't fabricate a folder path
   or day number when a call that was supposed to produce one failed).

   **Post-`queue-append` failures need special handling, since `queue-append` (step i)
   is the point of no return.** Once it succeeds, `queue-has` will report this number as
   already queued on every future run — there is no "un-append" operation. So a failure
   in any of steps (j) through (n) — hero rendering (other than the already-handled
   `hero-screenshot` case, which is non-fatal by design), writing the post drafts, or the
   final `git commit` — leaves a permanently "claimed" queue entry with incomplete or
   missing content, which a future re-run will silently skip forever rather than retry.
   If this happens: do NOT treat it as a normal per-problem skip. Call it out prominently
   and separately in the end-of-run summary (e.g. "Day N / question NUMBER was queued but
   its commit failed — needs manual follow-up: <error>"), so the user knows to
   investigate and finish that entry by hand rather than assuming a clean re-run will
   pick it up.

   a. Skip it if `queue-has` returns `{"has": true}` (the subcommand's JSON result, not a bare `true`):
      ```bash
      /tmp/leetcodectl queue-has '{"queuePath":"challenge/queue.yaml","number":<number>}'
      ```

   b. Resolve its location:
      ```bash
      /tmp/leetcodectl resolve '{"repoRoot":".","mapPath":"challenge/number_folder_map.yaml","number":<number>,"difficulty":"<difficulty>","slug":"<slug>"}'
      ```
      - If `"status": "unresolved"` — this question isn't solved yet (or genuinely can't be
        found). Skip it silently and move to the next problem; it'll be picked up on a future
        re-run once solved. Do NOT ask the user about every unresolved question — only ask if
        you have a specific, well-founded suspicion it IS solved but the tooling missed it
        (e.g. you can see a matching file via other means).
      - If `"source"` is anything other than `"canonical"` — the question was found in a
        non-canonical location (a root straggler, or under `google_questions`/`linkedin_questions`,
        or via the gitmap/fuzzy fallback). Note: `ResolveResult.Canonical` is tagged
        `json:"canonical,omitempty"`, so a non-canonical result never actually contains a
        `"canonical":false` key — the key is simply absent. Check `source`, not the
        presence/value of `canonical`. Reorganize it:
        ```bash
        /tmp/leetcodectl reorg '{"repoRoot":".","fromPath":"<result.path>","difficulty":"<difficulty>","number":<number>,"slug":"<slug>"}'
        ```
        Use the returned `path` as the folder for the rest of this step.

   c. Read the existing solution at `<folder>/main.go` (and `main_test.go` if present) —
      this is the source of truth for INTUITION.md, SOLUTION.md, and the inline comments.
      Never rewrite the algorithm; only add commentary.

   d. If `<folder>/README.md` is missing, fetch the statement and write it:
      ```bash
      /tmp/leetcodectl fetch-question '{"slug":"<slug>"}'
      ```
      Write `<folder>/README.md` from the returned `Title`/`Difficulty`/`ContentHTML`,
      converting the HTML statement to clean Markdown (headings, code blocks for
      examples, a Constraints list) — follow the style of existing README.md files
      elsewhere in the repo (e.g. `easy_problems/1_100/climbing_stairs/README.md`).

      `ContentHTML` may contain `<img>` tags (tree/graph diagrams, matrix illustrations,
      etc. — common on LeetCode statements). For each one: download it to
      `<folder>/images/<n>.<ext>` (`<ext>` from the URL or `Content-Type`, `<n>` a
      1-based index in the order the images appear) with
      `curl -sL '<image-url>' -o '<folder>/images/<n>.<ext>'`, and rewrite that image's
      markdown reference to the local relative path (`images/<n>.<ext>`) instead of the
      remote URL. Do not leave a remote LeetCode CDN URL in the committed README — this
      repo should be self-contained and not depend on LeetCode's CDN staying up. If a
      download fails, skip that one image (note it in the end-of-run summary) rather than
      failing the whole README.

   e. Write `<folder>/INTUITION.md`: a plain-language walkthrough of the approach used in
      the existing code — the key insight, why this technique applies, time/space
      complexity. Written for a reader who's seen the problem but not the solution.

   f. Write `<folder>/SOLUTION.md`: a narrated walkthrough of the ACTUAL code in
      `main.go` — reference real function/variable names, explain each meaningful step
      in the order they appear in the code. This is not a generic solution write-up; it
      must describe this specific implementation.

      Make the walkthrough visual wherever a picture beats a paragraph:
      - If step (d) downloaded example images (tree/graph/matrix diagrams) to
        `<folder>/images/`, reuse them here instead of re-describing the example input in
        prose — reference the local `images/<n>.<ext>` path. Don't regenerate an example
        that LeetCode's own statement already illustrates well.
      - For the parts of the walkthrough LeetCode doesn't illustrate — the data structure's
        state as the algorithm progresses (stack/queue contents, tree pointers, DP table
        fill order, etc.) — draw it yourself as one SVG per meaningful step, saved to
        `<folder>/images/walkthrough-<n>.svg` (1-based, in step order), and embed each
        inline at the point in SOLUTION.md that narrates that step. Keep each SVG small and
        focused (one state snapshot, not the whole trace) — a reader should be able to see
        at a glance what changed since the previous step.

      Reuse the same example images and walkthrough SVGs in the LinkedIn article (k) and
      Discord post (m) too — both carry the full solution walkthrough, so both should be
      visual for the same reasons, not just SOLUTION.md.

   g. Edit `<folder>/main.go` in place to add inline `//` comments at non-obvious steps
      (loop invariants, why a particular data structure, edge cases handled). Do not
      change any logic, formatting style, or the test file.

   h. Look up companies:
      ```bash
      /tmp/leetcodectl companies-lookup '{"datasetPath":"challenge/companies_dataset.json","number":<number>}'
      ```
      If the returned `companies` array is non-empty, write `<folder>/COMPANIES.md` listing
      them. If empty, don't create the file at all.

   i. Append to the queue FIRST — this assigns and returns the authoritative Day number
      that every remaining step in this iteration needs (hero image, both post drafts):
      ```bash
      /tmp/leetcodectl queue-append '{"queuePath":"challenge/queue.yaml","entry":{"number":<number>,"title":"<title>","difficulty":"<difficulty>","folder":"<folder>","batch":"<list-name>"}}'
      ```
      Use the returned `day` for steps (j)-(n) below. Do not try to predict or read
      `next_day` yourself before calling this — `queue-append` is the only source of
      truth for which day number a question gets, and calling it exactly once per
      question, before rendering anything that embeds the day number, is what keeps a
      multi-question batch's day numbers correct. (This ordering matters specifically
      because there is no read-only way to peek `next_day` without mutating it — the
      only two things that touch it are `queue-has`, which doesn't return it, and
      `queue-append`, which advances it. Calling `queue-append` first removes any need
      to predict its value.)

   j. Render and screenshot the hero image, using the real `day` from step (i):
      ```bash
      /tmp/leetcodectl hero-render-html '{"templatePath":"challenge/hero_template.html","outPath":"<folder>/HERO.html","data":{"Day":<day>,"Total":365,"Topic":"<list-name>","Difficulty":"<difficulty>","Title":"<title>"}}'
      /tmp/leetcodectl hero-screenshot '{"htmlPath":"<folder>/HERO.html","outPath":"<folder>/HERO.png"}'
      ```
      If `hero-screenshot` fails (e.g. Playwright/Chromium isn't installed — see
      `challenge/README.md`), leave `HERO.html` in place, note the failure to the user at
      the end of the run, and continue with the rest of the pipeline; don't block the batch
      on it.

   k. Write `<folder>/POST_LINKEDIN_ARTICLE.md`: the long-form piece, styled as a
      LinkedIn Newsletter/Article edition. Header "365 Days of LeetCode Challenge — Day
      <day>/365", question title + LeetCode link, the FULL intuition write-up (based on
      INTUITION.md, expanded for a public audience who hasn't seen INTUITION.md itself),
      the FULL solution walkthrough with the real code (based on SOLUTION.md), and the
      companies list if `COMPANIES.md` was written. This is the actual content — the
      short post below just points people at it.

   l. Write `<folder>/POST_LINKEDIN.md`: a SHORT teaser/share post for the article above.
      Header "365 Days of LeetCode Challenge — Day <day>/365", question title + LeetCode
      link, a 2-3 line intuition summary (a hook, not the full write-up), and a line
      pointing readers to the full article (e.g. "Full breakdown in today's newsletter
      article ⬇" — the actual link/attachment mechanics are the future posting
      subsystem's job, not this skill's). Do NOT include a Discord-join call to action in
      the body — that's posted separately as a comment by the (future) poster subsystem.

   m. Write `<folder>/POST_DISCORD.md`: header "365 Days of LeetCode Challenge — Day
      <day>/365", a 2-3 line intuition summary, then the ENTIRE solution — the real code
      in a fenced code block plus its walkthrough — using Discord markdown (`**bold**`,
      `` ``` `` code fences for the code block). Unlike the LinkedIn split, Discord gets
      the full content directly in one post; there's no separate "article" for Discord.

   n. Commit everything for this question in one commit:
      ```bash
      git add <folder>
      git add challenge/queue.yaml challenge/number_folder_map.yaml
      git commit -m "Content(<number>) <title> - Day <day>"
      ```

4. **At the end of the run**, summarize for the user: how many questions were processed,
   how many were skipped as not-yet-solved, any folders reorganized, any hero-image
   generation failures, any per-image download failures in README generation, and any
   questions skipped mid-pipeline due to an unexpected `leetcodectl` failure. Give any
   post-`queue-append` failure (steps j-n, per the note above) its OWN separate, clearly
   flagged line — e.g. "⚠ Day N / question NUMBER was queued but incomplete — needs
   manual follow-up: <error>" — do not fold it into the general skipped-questions bullet,
   since it cannot be silently retried on a future run.

## Notes

- This skill only produces content and updates `challenge/queue.yaml` — it never posts
  anything to LinkedIn or Discord. Posting is a separate, not-yet-built subsystem.
- Re-running this skill with the same URL/list is always safe — already-queued questions
  are skipped, so partial batches (e.g. "10 of 30 done this week") resume cleanly.
