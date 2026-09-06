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

## Writing style — sound human, not AI

**This applies to every commit message too, not only the prose files — see step 3(n).**
No `Co-Authored-By: Claude ...` trailer, no "Generated with Claude Code" line, no 🤖, no
mention of Claude, Anthropic, or a model, in any commit this skill makes. A `commit-msg`
hook in this repo rejects such messages, so a slip here fails the commit outright.

Every prose file this skill writes (INTUITION.md, SOLUTION.md, POST_LINKEDIN_ARTICLE.md,
POST_LINKEDIN.md, POST_DISCORD.md) goes out under the user's own name on their own
channels. Default LLM prose has well-known tells that make it read as machine-written —
avoid them from the first draft, don't write generic and "clean up" after:

- **Cut the AI vocabulary.** Don't use: delve, dive into, navigate (figurative), underscore,
  bolster, foster, harness, leverage, unpack, shed light on, pave the way, pivotal,
  groundbreaking, cutting-edge, transformative, game-changing, innovative, robust,
  comprehensive, seamless, intricate, nuanced, vibrant, multifaceted, holistic, testament,
  landscape/realm (figurative), crucial, enhance, garner, showcase, tapestry, interplay,
  align with, enduring. Say the plain thing instead ("shows" not "showcases", "uses" not
  "leverages").
- **Cut the AI sentence patterns.** No "It's not just X — it's Y", "Not only X, but Y",
  "This isn't about X. It's about Y.", "No X. No Y. Just Z." False contrasts like these are
  one of the most recognizable AI tells. Also drop throat-clearing openers ("In today's
  fast-paced world...", "It's important to note that...", "When it comes to...") and weak
  transitions ("This is where X comes in", "Let's break it down", "Let's dive in").
- **Don't force things into threes.** LLMs default to rule-of-three lists ("fast, reliable,
  and scalable") to sound thorough. Use however many points actually apply — one, two,
  four, whatever's true.
- **Skip the "Bold term: explanation" list format.** It's the single most recognizable AI
  formatting tell (`- **Performance:** Performance has been improved...`). Write it as
  prose, or as a plain list without the bolded-header-plus-colon pattern.
- **One em dash, maybe, not five.** AI prose overuses em dashes as a punchy connector.
  Prefer a period, a comma, or a parenthetical. If a paragraph already has an em dash,
  don't add a second.
- **No emoji decoration on headings or bullets**, no curly/smart quotes (use straight `"`
  `'`), sentence case in headings (not Title Case Everywhere).
- **Vary sentence length on purpose.** A run of same-length, same-shape sentences is a
  giveaway. Follow a long sentence with a short one sometimes.
- **Write like the author has an opinion.** This is the user's own solution and their own
  walkthrough of it — it's fine, even good, for INTUITION.md or a post to say what's
  actually satisfying, surprising, or fiddly about the approach, in first person, rather
  than neutrally narrating every step with no point of view. "This one's a nice one, the
  parent has to spot the left leaf because the leaf can't see itself" beats "This solution
  demonstrates an elegant technique."
- **Don't summarize with a generic uplifting close.** No "This represents a great learning
  opportunity" or "Exciting times ahead" wrap-ups. End when the point is made.
- **Never leave in chatbot artifacts** — no "I hope this helps!", "Let's explore...",
  "Great question!", or any other trace of this having been a conversation with an
  assistant. The reader should never be able to tell an LLM was involved in writing it.

Run a quick self-check on every prose file before considering it done: read it back and
ask whether it sounds like a specific person who solved this problem talking about it, or
like a generic summary that could've been written about any problem. If it's the latter,
rewrite it.

## Quality Validation Checklist

Before marking a problem complete, verify:

**Solution Walkthrough (step f)**
- [ ] Walkthrough trace is accurate end-to-end (manually verify)
- [ ] Each step's output is correct
- [ ] No ASCII art trees in SOLUTION.md (use LeetCode example image + walkthrough diagrams only)
- [ ] All image references use local paths: `images/filename.png` (not GitHub URLs)
- [ ] Code in main.go correctly implements the described algorithm
- [ ] Variable names in code match walkthrough description

**Images & Diagrams**
- [ ] Example images downloaded from LeetCode (if present in problem statement)
- [ ] Walkthrough diagrams show algorithm progression at key steps
- [ ] All `.svg` converted to `.png` and SVG files deleted (PNG only in repo)
- [ ] Image paths consistent across SOLUTION.md and POST_LINKEDIN_ARTICLE.md (both use `images/`)
- [ ] No `?raw=true` or GitHub blob URLs in markdown files

**Complexity Analysis (step e/f)**
- [ ] Time complexity is accurate for this specific algorithm
- [ ] Space complexity is accurate (includes recursion stack, temp data structures)
- [ ] Complexity matches code, not theoretical best-case

**Code Comments (step g)**
- [ ] Comments added only at non-obvious steps
- [ ] No logic changes to main.go
- [ ] Comments explain WHY, not WHAT

**Content Files**
- [ ] No conflicting or incorrect information (no mixed right/wrong diagrams)
- [ ] All references use local image paths (no external URLs in SOLUTION.md, POST_LINKEDIN_ARTICLE.md)
- [ ] No AI tell vocabulary ("leverage", "harness", "dive into", etc.)
- [ ] No generic sentence patterns ("It's not just X, it's Y", etc.)

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
   handled below — unresolved `resolve`, failed `hero-generate`) exits non-zero or
   returns unparseable output: stop processing THIS problem only, note it in the
   end-of-run summary with the error message, and continue to the next problem in the
   list. Never let one problem's failure abort the whole batch, and never silently
   proceed past a failed call as if it had succeeded (e.g. don't fabricate a folder path
   or day number when a call that was supposed to produce one failed).

   **Post-`queue-append` failures need special handling, since `queue-append` (step i)
   is the point of no return.** Once it succeeds, `queue-has` will report this number as
   already queued on every future run — there is no "un-append" operation. So a failure
   in any of steps (j) through (n) — hero rendering (other than the already-handled
   `hero-generate` case, which is non-fatal by design), writing the post drafts, or the
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
      Follow the "Writing style" section above.

   f. Write `<folder>/SOLUTION.md`: a narrated walkthrough of the ACTUAL code in
      `main.go` — reference real function/variable names, explain each meaningful step
      in the order they appear in the code. This is not a generic solution write-up; it
      must describe this specific implementation. Follow the "Writing style" section
      above.

      Make the walkthrough visual wherever a picture beats a paragraph:
      - If step (d) downloaded example images (tree/graph/matrix diagrams) to
        `<folder>/images/`, reuse them here instead of re-describing the example input in
        prose — reference the local `images/<n>.<ext>` path. Don't regenerate an example
        that LeetCode's own statement already illustrates well.
      - **Never use ASCII art trees or diagrams.** Use visual images only (example images + walkthrough diagrams).
      - **Do NOT create walkthrough diagrams that are just ASCII art trees rendered as images** — that wastes tokens and space. Create diagrams ONLY if they show actual algorithm progression (state changes, data structure contents evolving, etc.) that cannot be explained clearly in prose. For tree problems, reference the LeetCode example image instead.
      - For tree/graph traversal algorithms where the visual state change is essential (e.g., level-by-level order, stack/queue contents changing), create proper walkthrough diagrams showing algorithm progression, saved to `<folder>/images/walkthrough-<n>.png` (1-based, in step order). Keep each diagram small and focused (one state snapshot) — reader should see at a glance what changed since previous step.

      **Every walkthrough SVG must be exported to PNG**, same basename, into the folder's
      `images/` directory — and the SVG itself must not be left there. Draw the SVGs in a
      temp directory, convert, then delete them:
      ```bash
      rsvg-convert -w 1200 --keep-aspect-ratio -b white \
        -o <folder>/images/walkthrough-<n>.png /tmp/walkthrough-<n>.svg
      ```
      **Reference the `.png` everywhere — SOLUTION.md, POST_LINKEDIN_ARTICLE.md, all of
      it. Never embed the `.svg`.** LinkedIn's article editor rejects SVG uploads outright,
      and an SVG that's subtly malformed fails silently as a broken image rather than
      erroring, which is how a batch of them once shipped broken. Since nothing references
      the SVG, keeping it in the solution folder only adds a second copy of every diagram
      for readers to trip over — the PNG is the artifact. (The trade: a diagram that needs
      changing later gets redrawn rather than edited.)

      After exporting, confirm nothing is cut off at the edges — text that overflows the
      `viewBox` gets clipped mid-sentence in both formats:
      ```bash
      python3 -c "
      from PIL import Image; import numpy as np, sys
      a = np.array(Image.open(sys.argv[1]).convert('L')) < 240
      edges = a[:,0].sum() + a[:,-1].sum() + a[0,:].sum() + a[-1,:].sum()
      print('CLIPPED - widen the viewBox' if edges else 'ok')" <folder>/images/walkthrough-<n>.png
      ```

      **Never write `--` inside an SVG comment.** A double hyphen is illegal in XML
      comments, and it silently breaks the whole file: the SVG won't render on GitHub or
      convert to PNG. Use a single hyphen or an en dash for a parenthetical instead. After
      writing the SVGs, verify every one parses before moving on:
      ```bash
      python3 -c "import xml.etree.ElementTree as ET,glob,sys; [ET.parse(f) for f in glob.glob('/tmp/walkthrough-*.svg')]" && echo "all SVGs valid"
      ```

      Reuse the same example images and walkthrough PNGs in the LinkedIn article (k) and
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
      /tmp/leetcodectl hero-generate '{"templatePath":"challenge/hero_template.html","outPath":"<folder>/HERO.png","data":{"Day":<day>,"Total":365,"Topic":"<list-name>","Difficulty":"<difficulty>","Title":"<title>"}}'
      ```
      `hero-generate` renders the HTML to a temp file, screenshots it, and deletes it.
      **`HERO.png` is the only hero artifact that belongs in a solution folder** — never
      write `HERO.html` there. If the command fails (e.g. Playwright/Chromium isn't
      installed — see `challenge/README.md`), it keeps the rendered HTML in the temp dir
      and names that path in the error; report the failure to the user at the end of the
      run and continue with the rest of the pipeline, don't block the batch on it.

      The card's background colour comes from `hero.PaletteFor(day)`, which rotates through
      eight dark gradients so consecutive days don't look like the same post re-sent. It's
      derived from the day, not random, so re-rendering a day reproduces its image. Every
      other element — the orange accent, the logos, the type, the layout — is fixed, and
      that's the branding; don't vary it per day. Adding a palette means adding it to
      `hero.Palettes`, where a test enforces WCAG AA contrast for the card's text.

      `challenge/hero_template.html` already has the CodeStreak Daily logo baked in (as a
      base64 `<img>`, sourced originally from `/Users/architagarwal/CodeStreak Daily/codestreak_logo.png`)
      alongside the LeetCode badge and author photo — every hero rendered from this template
      is automatically branded, nothing extra to do per-question. This content series
      publishes under the CodeStreak Daily newsletter specifically; the weekly golang
      journal is a separate, unrelated newsletter and its logo should NOT be added here.

   k. Write `<folder>/POST_LINKEDIN_ARTICLE.md`: the long-form piece, styled as a
      LinkedIn Newsletter/Article edition. Header "365 Days of LeetCode Challenge — Day
      <day>/365", question title + LeetCode link, the FULL intuition write-up (based on
      INTUITION.md, expanded for a public audience who hasn't seen INTUITION.md itself),
      the FULL solution walkthrough with the real code (based on SOLUTION.md), and the
      companies list if `COMPANIES.md` was written. This is the actual content — the
      short post below just points people at it. Follow the "Writing style" section
      above. End with a line of 5-8 hashtags (see hashtags note below step m), followed by
      the AI-disclosure line (see below) as the very last thing in the file.

      **AI disclosure (POST_LINKEDIN_ARTICLE.md only, after the hashtags):**
      ```
      ---

      *Solution and code by Archit Agarwal. Write-up drafted with AI assistance from the
      code and problem statement.*
      ```
      This is deliberate, not a hedge against the "Writing style" work above — the two
      aren't in tension. The prose should still read like a person wrote it; this line is
      just an honest, low-key credit line, not a disclaimer plastered over the top. It
      goes ONLY in the long-form article, not in POST_LINKEDIN.md (the short teaser) or
      POST_DISCORD.md.

   l. Write `<folder>/POST_LINKEDIN.md`: a SHORT teaser/share post for the article above.
      Header "365 Days of LeetCode Challenge — Day <day>/365", question title + LeetCode
      link, a 2-3 line intuition summary (a hook, not the full write-up), and a line
      pointing readers to the full article (e.g. "Full breakdown in today's newsletter
      article ⬇" — the actual link/attachment mechanics are the future posting
      subsystem's job, not this skill's). Do NOT include a Discord-join call to action in
      the body — that's posted separately as a comment by the (future) poster subsystem.
      Follow the "Writing style" section above. End with a line of 5-8 hashtags (see
      hashtags note below).

   m. Write `<folder>/POST_DISCORD.md`. **This file IS the exact Discord message** — the
      poster subsystem posts its bytes verbatim, with no parsing, trimming or
      transformation. Whatever is in this file is what lands in the channel, so it has to
      satisfy Discord's own limits:

      - **Hard cap: keep the whole file under 1900 characters.** Discord rejects a message
        body over 2000, and the poster fails loudly rather than truncating mid-sentence.
        Check the actual byte count before you finish (`wc -c <folder>/POST_DISCORD.md`).
      - **No SVG embeds, and no image references at all.** Discord only previews
        PNG/JPG/GIF/WEBP, so the `walkthrough-<n>.svg` diagrams cannot render there. Don't
        reference them — the poster attaches `HERO.png` itself, and the walkthrough link
        (below) is how people reach the diagrams. Markdown image syntax in this file just
        shows up as raw noise in the channel.
      - **The code block usually won't be the whole file.** Include the core function(s)
        only — drop the `TreeNode` struct boilerplate, imports, and inline comments, which
        the reader can get from the repo. If even the core function would bust the budget
        (some solutions are 2000+ characters on their own), drop the code block entirely
        and let the walkthrough link carry it. Never ship a file over the cap "because the
        code needed it".

      Structure, in order:
      ```
      **365 Days of LeetCode Challenge — Day <day>/365**
      **<title>** (<difficulty>)
      🔗 <leetcode url>

      <2-3 line intuition, per the Writing style section>

      ```go
      <core function(s), or omit this block if it doesn't fit>
      ```

      Full walkthrough with step-by-step diagrams: <github blob url to that folder's SOLUTION.md>
      ```

      The GitHub link is what replaces the old "full content inline" approach — it points
      at `https://github.com/architagr/leetcode_solutions/blob/main/<folder>/SOLUTION.md`,
      where the SVG walkthrough diagrams do render. Follow the "Writing style" section
      above. No hashtag line here — hashtags are a LinkedIn convention, not a Discord one.

      **Hashtags (POST_LINKEDIN_ARTICLE.md and POST_LINKEDIN.md only):** a single line of
      5-8 hashtags at the very end, mixing a few general ones (from e.g. `#DSA #LeetCode
      #100DaysOfCode #SoftwareEngineering #CodingInterview #TechCareer #Programming
      #Algorithms`) with 1-3 that name this problem's actual topic/technique (e.g.
      `#BinaryTree #BFS #DFS #Recursion #BinarySearchTree #Golang` — pick whichever
      genuinely apply, don't reuse the same topic tags for every problem regardless of fit).

   n. Commit everything for this question in one commit. The message body is
      user-facing product content like everything else this skill writes — it must never
      mention Claude, an AI assistant, or that the content was generated by a model (no
      "Generated with Claude", no `Co-Authored-By: Claude ...` trailer, nothing like it).
      The repo's `commit-msg` hook enforces this; a message that trips it fails the commit
      and the queue entry stays unclaimed until the message is fixed:
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
