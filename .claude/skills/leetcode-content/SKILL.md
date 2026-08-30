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

1. **Refresh the git-history cache:**
   ```bash
   /tmp/leetcodectl gitmap-update '{"repoRoot":".","mapPath":"challenge/number_folder_map.yaml"}'
   ```

2. **Fetch the problem list:**
   ```bash
   /tmp/leetcodectl fetch-list '{"url":"<the URL>","difficulty":"<the --filter value>"}'
   ```
   Sort the returned problems by `Number` ascending.

3. **For each problem, in that sorted order:**

   a. Skip it if `queue-has` returns `{"has": true}`:
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
      - If `"status": "resolved"` but `canonical` is not `true` — note that the JSON encoder
        omits the `canonical` field entirely when it's `false`, so a non-canonical result
        looks like `{"status":"resolved","path":"...","source":"straggler"}` with no
        `canonical` key at all, rather than `"canonical":false`. Treat "canonical key absent"
        the same as "canonical false". This happens when `source` is `straggler`,
        `google_questions`, `linkedin_questions`, `gitmap`, or `fuzzy` — i.e. anything other
        than `source: "canonical"`. Reorganize it:
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

   e. Write `<folder>/INTUITION.md`: a plain-language walkthrough of the approach used in
      the existing code — the key insight, why this technique applies, time/space
      complexity. Written for a reader who's seen the problem but not the solution.

   f. Write `<folder>/SOLUTION.md`: a narrated walkthrough of the ACTUAL code in
      `main.go` — reference real function/variable names, explain each meaningful step
      in the order they appear in the code. This is not a generic solution write-up; it
      must describe this specific implementation.

   g. Edit `<folder>/main.go` in place to add inline `//` comments at non-obvious steps
      (loop invariants, why a particular data structure, edge cases handled). Do not
      change any logic, formatting style, or the test file.

   h. Look up companies:
      ```bash
      /tmp/leetcodectl companies-lookup '{"datasetPath":"challenge/companies_dataset.json","number":<number>}'
      ```
      If the returned `companies` array is non-empty, write `<folder>/COMPANIES.md` listing
      them. If empty, don't create the file at all.

   i. Render and screenshot the hero image. The `day` value is whatever the NEXT
      `queue-append` call in step (j) will assign — since `queue.yaml`'s `next_day` is
      known before appending, use that value here (it becomes the entry's `day`):
      ```bash
      /tmp/leetcodectl hero-render-html '{"templatePath":"challenge/hero_template.html","outPath":"<folder>/HERO.html","data":{"Day":<next_day>,"Total":365,"Topic":"<list-name>","Difficulty":"<difficulty>","Title":"<title>"}}'
      /tmp/leetcodectl hero-screenshot '{"htmlPath":"<folder>/HERO.html","outPath":"<folder>/HERO.png"}'
      ```
      If `hero-screenshot` fails (e.g. Playwright/Chromium isn't installed — see
      `challenge/README.md`), leave `HERO.html` in place, note the failure to the user at
      the end of the run, and continue with the rest of the pipeline; don't block the batch
      on it.

   j. Append to the queue (this assigns and returns the Day number — use it for the post
      drafts below):
      ```bash
      /tmp/leetcodectl queue-append '{"queuePath":"challenge/queue.yaml","entry":{"number":<number>,"title":"<title>","difficulty":"<difficulty>","folder":"<folder>","batch":"<list-name>"}}'
      ```

   k. Write `<folder>/POST_LINKEDIN.md`: header "365 Days of LeetCode Challenge — Day
      <day>/365", question title + LeetCode link, a 2-3 line intuition hook (not the full
      writeup — a teaser), a short code snippet or a link to the repo file, and the
      companies list if `COMPANIES.md` was written. Do NOT include a Discord-join call to
      action in the body — that's posted separately as a comment by the (future) poster
      subsystem.

   l. Write `<folder>/POST_DISCORD.md`: a shorter version of the same content, using
      Discord markdown (`**bold**`, `` ``` `` code fences), same Day N header.

   m. Commit everything for this question in one commit:
      ```bash
      git add <folder>
      git add challenge/queue.yaml challenge/number_folder_map.yaml
      git commit -m "Content(<number>) <title> - Day <day>"
      ```

4. **At the end of the run**, summarize for the user: how many questions were processed,
   how many were skipped as not-yet-solved, any folders reorganized, and any hero-image
   generation failures.

## Notes

- This skill only produces content and updates `challenge/queue.yaml` — it never posts
  anything to LinkedIn or Discord. Posting is a separate, not-yet-built subsystem.
- Re-running this skill with the same URL/list is always safe — already-queued questions
  are skipped, so partial batches (e.g. "10 of 30 done this week") resume cleanly.
