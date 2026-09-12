---
name: leetcode-content
description: 'Generate teaching content (intuition/solution docs, commented code, companies, hero image, LinkedIn/Discord/X/Substack post drafts) for already-solved LeetCode questions from a problem-list URL, in resumable "365 Days of LeetCode Challenge" order. Trigger: /leetcode-content'
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

## LeetCode session cookie

Step 3(b) asks LeetCode whether the user has solved a question that has no folder here,
and pulls their accepted submission down when they have. Both are user-scoped, so they
need the user's own login. `leetcodectl` reads it from `.leetcode_session` at the repo
root (gitignored), or from `$LEETCODE_SESSION` / `$LEETCODE_CSRF`:

```json
{ "leetcodeSession": "<LEETCODE_SESSION cookie value>", "csrfToken": "<csrftoken cookie value>" }
```

To get the values: log in at leetcode.com, open the browser's dev tools, and copy the
`LEETCODE_SESSION` and `csrftoken` cookies for `leetcode.com`. In Safari the cookie jar is
under Develop > Show Web Inspector > Storage > Cookies (the Develop menu has to be turned
on in Settings > Advanced first).

This file is a credential. It is gitignored, it must never be committed, echoed into a
post draft, or pasted into a commit message, and it expires every couple of weeks — an
auth error from `scaffold-from-submission` usually just means it needs refreshing.

Note that **Playwright cannot drive Safari**. It ships a bundled WebKit build, which is
not Safari.app and has none of Safari's cookies or logins. Nothing in this skill automates
a browser for LeetCode; the cookie is copied across by hand once, and every LeetCode call
is a plain authenticated GraphQL request.

## Writing style — sound human, not AI

**This applies to every commit message too, not only the prose files — see step 3(n).**
No `Co-Authored-By: Claude ...` trailer, no "Generated with Claude Code" line, no 🤖, no
mention of Claude, Anthropic, or a model, in any commit this skill makes. A `commit-msg`
hook in this repo rejects such messages, so a slip here fails the commit outright.

Every prose file this skill writes (INTUITION.md, SOLUTION.md, POST_LINKEDIN_ARTICLE.md,
POST_LINKEDIN.md, POST_DISCORD.md, POST_X.md, POST_SUBSTACK.md) goes out under the user's own name on their own
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
- [ ] `images/walkthrough-*.png` exists — every solution has them, 3-6 of them
- [ ] Each diagram shows state that actually changed, not the input redrawn
- [ ] Diagrams match the house style (1200px, STEP n / N, blue done / orange current / grey pending)
- [ ] The same diagrams appear in SOLUTION.md, POST_LINKEDIN_ARTICLE.md and POST_SUBSTACK.md
- [ ] No image references in POST_DISCORD.md
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

**Solved-but-missing problems**
- [ ] Every `unresolved` question was checked against LeetCode, not skipped silently
- [ ] End-of-run summary has a "Pulled from LeetCode" section (even if it says none)
- [ ] Any pulled submission was committed on its own, before the content commit
- [ ] No non-Go submission was committed or given generated content

**Queue and chaining**
- [ ] `builds_on` set, and every entry in it is on an EARLIER day (no forward references)
- [ ] "Builds on" section in INTUITION.md and the LinkedIn article uses full GitHub URLs
- [ ] Each link says what that earlier problem gives you, not just "related problem"
- [ ] Queue is not a long single-difficulty run (see "Queue ordering")
- [ ] `Day <n>/365` in every post matches this folder's queue entry
- [ ] No frozen (already-posted) entry had its day changed

**Content Files**
- [ ] No conflicting or incorrect information (no mixed right/wrong diagrams)
- [ ] All references use local image paths (no external URLs in SOLUTION.md, POST_LINKEDIN_ARTICLE.md)
- [ ] No AI tell vocabulary ("leverage", "harness", "dive into", etc.)
- [ ] No generic sentence patterns ("It's not just X, it's Y", etc.)

## Queue ordering — mix easy with medium and hard

**Never let the queue become a long run of one difficulty.** Publishing 28 consecutive
easies on a single topic loses the audience well before the mediums arrive, and dumping
the mediums at the end means nobody who followed the easy run ever sees them.

Target shape: **3-4 easy, then 1-2 medium or hard, repeating.** A reader finishing a month
should have cleared a real mix and feel ready for mediums on that topic, not just have
done thirty warm-ups.

Two rules constrain where a question can go:

1. **A question must come after everything in its `builds_on` list.** That is the whole
   point of the chain — the medium lands while the technique from its prerequisites is
   still fresh. Verify this before writing the queue; a medium scheduled before its
   prerequisite is a worse bug than a bad ratio, because the write-up will link forward
   to a day the reader has not reached.
2. **An entry whose `posted_at` has any non-null timestamp is frozen.** Its day number is
   already public on LinkedIn and Discord. Renumber around it, never through it.

When the ratio and the prerequisites conflict, prerequisites win. Expect the tail of a
batch to drift medium-heavy as the easies run out, which is fine and arguably right.

**Reordering an existing queue** (as opposed to appending to a well-shaped one) is a
bigger operation than it looks, because the day number is copied into published-facing
artifacts. Renumbering day N to day M means all of:

- rewrite `Day <N>/365` and `![Day <N>](HERO.png)` in `POST_LINKEDIN.md`,
  `POST_DISCORD.md`, `POST_LINKEDIN_ARTICLE.md`, `POST_X.md` and `POST_SUBSTACK.md`
- re-render `HERO.png` with the new day (the card has the number printed on it)
- fix any `[Day <N>: <title>](...)` cross-reference in *other* questions' files that
  points at this folder

Audit that last one across the whole repo rather than per-folder, since the stale link
lives in the referring question, not the moved one. After reordering, re-check that every
`Day <n>/365` in a folder matches that folder's queue entry, and that every cross-ref day
matches the queue entry of the folder it links to.

## Problem chaining — `builds_on`

Each queue entry carries `builds_on`, a list of earlier LeetCode numbers whose technique
this question reuses. It turns the series into a chain: a reader arriving at a medium gets
pointed back at the easies that taught the pieces.

Pick prerequisites by **technique actually reused**, not by superficial topic overlap. Two
or three is usually right, and zero is a perfectly good answer for a question that
introduces something new. Some worked examples from the binary-tree batch:

| Question | `builds_on` | Why |
|---|---|---|
| 222 Count Complete Tree Nodes | 104, 110 | height recursion, plus bottom-up returning more than one fact |
| 102 Level Order Traversal | 637 | 637 is the gentle introduction to BFS by level |
| 103 Zigzag Level Order | 102 | same traversal, one twist on top |
| 2265 Count Nodes Equal to Average | 543, 563 | both are bottom-up "return a pair up the tree" |
| 124 Max Path Sum (hard) | 543 | diameter is the same split-at-a-node shape |
| 236 LCA of a Binary Tree | 235 | the BST version first, then drop the ordering guarantee |
| 450 Delete Node in a BST | 700, 98 | search to find the node, validity to know what may move |

Two things to keep honest:

- **Only list problems already in the queue on an earlier day.** A forward reference is
  broken for every reader following along in order.
- **Say what the earlier problem gives you.** "Related problem" is filler; "bottom-up
  returning more than one fact" is the sentence that makes the chain worth reading.

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

   As of 2026-09, LeetCode's GraphQL rejects the difficulty filter outright:
   `Variable "$filter" got invalid value {"difficulty": "MEDIUM"}. In field "difficulty":
   Unknown field.` Until `fetch-list` is updated for the new schema, omit `difficulty`
   from the payload and filter the returned array yourself on its `Difficulty` field.

3. **For each problem, in that sorted order:**

   If any `leetcodectl` call in this loop (other than the two failure modes explicitly
   handled below — unresolved `resolve`, failed `hero-generate`) exits non-zero or
   returns unparseable output: stop processing THIS problem only, note it in the
   end-of-run summary with the error message, and continue to the next problem in the
   list. Never let one problem's failure abort the whole batch, and never silently
   proceed past a failed call as if it had succeeded (e.g. don't fabricate a folder path
   or day number when a call that was supposed to produce one failed).

   **A failure after `queue-append` (step i) leaves the day claimed but unwritten.**
   There is no "un-append" operation, so a failure in any of steps (j) through (n) —
   hero rendering (other than the already-handled `hero-generate` case, which is
   non-fatal by design), writing the post drafts, or the final `git commit` — leaves a
   queue entry whose number is taken and whose content is incomplete.

   That state is recoverable rather than permanent, because the entry is still
   `pending_content`: `queue-lookup` reports `"needsContent": true` for it, so the next
   run picks it up and finishes it instead of skipping it. Only `mark-content-ready`,
   at the very end of a successful iteration, declares a question done.

   Still report it. Call it out prominently and separately in the end-of-run summary
   (e.g. "Day N / question NUMBER was queued but its commit failed — will be retried on
   the next run: <error>"), because a half-written folder sitting in the working tree is
   something the user should see rather than discover later.

   a. Look the question up in the queue and branch on what comes back:
      ```bash
      /tmp/leetcodectl queue-lookup '{"queuePath":"challenge/queue.yaml","number":<number>}'
      ```
      Three outcomes, and conflating the first two strands a day forever:

      - `{"has": false}` — not queued. Generate content and **append** it, which assigns
        the day number (step i).
      - `{"has": true, "needsContent": true}` — **queued but not yet written.** Its day
        number is already reserved, and posting skips it until content exists. Generate
        content for it exactly as normal, but use the `day` the lookup returned rather
        than appending, and finish with `mark-content-ready` instead of `queue-append`
        (see step i-bis). Do **not** skip these.
      - `{"has": true, "needsContent": false}` — already written. Skip it silently.

      The reserved-day case is how a batch's topic and difficulty shape gets decided
      before any of it is written: the whole ordering is laid out first, then the
      write-ups are filled in arc by arc. Treating "queued" as "done" — which this step
      used to do — leaves every reserved day permanently unwritten, because posting
      skips a day with no content and nothing ever comes back to fill it in.

   b. Resolve its location:
      ```bash
      /tmp/leetcodectl resolve '{"repoRoot":".","mapPath":"challenge/number_folder_map.yaml","number":<number>,"difficulty":"<difficulty>","slug":"<slug>"}'
      ```
      - If `"status": "unresolved"` — there is no folder for this question in the repo.
        That has two very different causes, and they must not be collapsed: the user may
        genuinely not have solved it, or they may have solved it **on LeetCode** and never
        committed the code here. Ask LeetCode which it is, and pull the code down if so:
        ```bash
        /tmp/leetcodectl scaffold-from-submission '{"repoRoot":".","sessionPath":".leetcode_session","number":<number>,"difficulty":"<difficulty>","slug":"<slug>","title":"<title>"}'
        ```
        This needs an authenticated session — see "LeetCode session cookie" below. Act on
        the returned `status`:
        - `"not-solved"` — genuinely unsolved. Skip it silently and move on; a future
          re-run picks it up once solved. Do NOT ask the user about these.
        - `"scaffolded"` — solved on LeetCode, missing here. The canonical folder now
          exists with their own accepted submission in it, at the returned `file`. Commit
          that on its own, before any content work, so the pulled code is one reviewable
          commit separate from the generated prose:
          ```bash
          git add <result.path>
          git commit -m "<Difficulty>(<number>) <Title>"
          ```
          Then continue this problem from step (c) as normal — the folder is now a
          solved folder like any other. Record it for the end-of-run summary as
          "pulled from LeetCode (submitted <result.submittedAt>)".
        - `"scaffolded"` with `"nonGo": true` — their accepted submission was not Go, so
          the code landed under its own extension (e.g. `main.python3`) and the repo's
          Go-only convention is broken. Do NOT commit it and do NOT generate content.
          Leave the file in the working tree, stop processing this problem, and flag it
          prominently in the end-of-run summary as needing a hand-written Go port.
        - `"exists"` — a solution file is already there but `resolve` could not see it.
          That is a tooling miss, not a missing solution. Do not overwrite anything; use
          the returned `path` as the folder and continue from step (c), and note the
          resolver gap in the end-of-run summary.
        - The command exiting non-zero because no session cookie is configured is NOT a
          per-problem failure. Stop the whole run and tell the user to set the cookie up
          (see below) — otherwise every unsolved-looking question in the batch gets
          misreported as unsolved.

      **The end-of-run summary must always have a "Pulled from LeetCode" section** listing
      every problem that was solved on LeetCode but missing from the repo, even when the
      count is zero. This is the gap the user asked to see; a silent skip hides it.
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

      **Chain it to the earlier problems it builds on.** The queue entry carries a
      `builds_on` list of problem numbers (see the "Problem chaining" section below).
      For each one, look up its queue entry to get its day number, title and folder, and
      add a short section near the top of INTUITION.md:

      ```markdown
      ## Builds on

      - [Day 1: Maximum Depth of Binary Tree](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/maximum_depth_of_binary_tree/) — the height recursion this reuses
      - [Day 5: Balanced Binary Tree](https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/balanced_binary_tree/) — bottom-up returning more than one fact
      ```

      Write a real reason after the em dash saying what that earlier problem gives you
      here, not a generic "related problem" label. The point is that a reader following
      the series in order recognises the technique instead of meeting it cold. If
      `builds_on` is empty, skip the section entirely rather than inventing links.

      Use full `https://github.com/architagr/leetcode_solutions/blob/main/<folder>/` URLs,
      never relative paths — this section is copied into the LinkedIn article (step k),
      where relative links are dead.

   f. Write `<folder>/SOLUTION.md`: a narrated walkthrough of the ACTUAL code in
      `main.go` — reference real function/variable names, explain each meaningful step
      in the order they appear in the code. This is not a generic solution write-up; it
      must describe this specific implementation. Follow the "Writing style" section
      above.

      **Every solution gets walkthrough diagrams. This is not optional.** A solution
      folder with no `images/walkthrough-*.png` is incomplete, in the same way one with no
      SOLUTION.md would be. The whole series is a teaching series, and a reader following
      a recursion through four levels needs to see the state at each level, not be asked
      to hold it in their head.

      Aim for 3-6 of them, saved to `<folder>/images/walkthrough-<n>.png` (1-based, in
      step order). Each one is a single state snapshot, and a reader should see at a
      glance what changed since the last.

      What they must show is **algorithm progression** — the state that actually moves as
      the code runs. The accumulator filling up, the queue draining, the pointers
      relinking, which node is current and which are already done. A picture of the input
      tree, redrawn five times with nothing changing, is not a walkthrough diagram.

      - **Never use ASCII art, in any file, for anything.** Not in SOLUTION.md, not in a
        post, and not rendered into a PNG and called a diagram. If a tree needs drawing,
        draw it as an SVG and export it.
      - If step (d) downloaded example images to `<folder>/images/`, use them for the
        *input* rather than redrawing it — reference the local `images/<n>.<ext>` path.
        Those illustrate the problem; the walkthrough diagrams illustrate the solution.
        Both belong in the file, doing different jobs.

      **House style, so the series looks like one series.** Match the existing diagrams
      (see `easy_problems/101_200/binary_tree_preorder_traversal/images/`): 1200px wide,
      white background, `STEP n / N` in grey monospace at the top left, and one line of
      monospace caption along the bottom naming the state after this step (`arr = [1,2,3]`).
      Nodes are circles, radius 45, dark `#2c3e50` stroke: blue `#3178c6` filled for done,
      a thick orange `#e8890c` ring on white for the node being visited now, grey `#d3d3d3`
      for not yet reached. Edges are `#8a9a9a`, 4px.

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

      **The same example images and walkthrough PNGs go into the LinkedIn article (k) and
      the Substack post (o), at the same points in the narrative.** All three files carry
      the full walkthrough, so all three are visual for the same reason — SOLUTION.md is
      not the only one a reader learns from, and an article that describes in prose what
      the repo shows in pictures is the weaker copy of the two.

      Reference them by their local `images/walkthrough-<n>.png` path in all three. The
      images are attached by hand when a post is pasted into LinkedIn's or Substack's
      editor, so the path is a pointer for the author rather than something the platform
      resolves, and keeping all three files on the same paths means a renamed diagram is a
      one-place edit.

      **Not the Discord post.** Step (m) is explicit that POST_DISCORD.md carries no image
      references at all — the poster attaches HERO.png itself and markdown image syntax
      just lands as raw noise in the channel. The walkthrough link at the bottom of that
      post is how Discord readers reach the diagrams.

   g. Edit `<folder>/main.go` in place to add inline `//` comments at non-obvious steps
      (loop invariants, why a particular data structure, edge cases handled). Do not
      change any logic, formatting style, or the test file.

   h. Look up companies:
      ```bash
      /tmp/leetcodectl companies-lookup '{"datasetPath":"challenge/companies_dataset.json","number":<number>}'
      ```
      If the returned `companies` array is non-empty, write `<folder>/COMPANIES.md` listing
      them. If empty, don't create the file at all.

   i. **If step (a) reported `"needsContent": true`, do not append.** The day number is
      already reserved and `queue-append` would assign a second one to a question that
      already has a day. Use the `day` the lookup returned for steps (j)-(n), and close
      the question out with `mark-content-ready` after its files are written and
      committed:
      ```bash
      /tmp/leetcodectl mark-content-ready '{"queuePath":"challenge/queue.yaml","numbers":[<number>]}'
      ```
      The entry already carries `title`, `difficulty`, `folder`, `batch` and `builds_on`
      from when the ordering was laid out, so none of those need setting again — flipping
      the status is the whole of it. Everything else in this iteration is unchanged.
      Then skip to step (j).

      Otherwise, append to the queue FIRST — this assigns and returns the authoritative
      Day number that every remaining step in this iteration needs (hero image, both post
      drafts):
      ```bash
      /tmp/leetcodectl queue-append '{"queuePath":"challenge/queue.yaml","entry":{"number":<number>,"title":"<title>","difficulty":"<difficulty>","folder":"<folder>","batch":"<list-name>","builds_on":[<earlier problem numbers>]}}'
      ```
      Set `builds_on` per the "Problem chaining" section below. It drives the "Builds on"
      section in step (e) and the LinkedIn article, so fill it in at append time rather
      than backfilling later.

      **A plain append puts this question at the end of the queue, which is only right
      when the queue is already in mixed order.** If you are adding a medium or hard to a
      queue whose tail is a run of easies, read the "Queue ordering" section below first —
      appending a batch of mediums onto the end is exactly the failure mode that section
      exists to prevent.
      Use the returned `day` for steps (j)-(n) below. Do not try to predict or read
      `next_day` yourself before calling this — `queue-append` is the only source of
      truth for which day number a question gets, and calling it exactly once per
      question, before rendering anything that embeds the day number, is what keeps a
      multi-question batch's day numbers correct. (This ordering matters specifically
      because there is no read-only way to peek `next_day` without mutating it — the
      only two things that touch it are `queue-lookup`, which doesn't return it, and
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
      LinkedIn Newsletter/Article edition.

      **Open the file with a YAML front matter block** carrying the SEO metadata LinkedIn
      asks for in its own fields at publish time. (`tags:` is Substack's field — LinkedIn
      has no equivalent, so leave it out here.) `leetcodectl linkedin-batch` lifts this
      out into a "publish settings" block and warns when a field is missing or too long,
      so it is worth getting right here rather than in the composer:
      ```
      ---
      meta_title: "<=60 chars"
      meta_description: "<=155 chars"
      ---
      ```
      - `meta_title` is NOT the article's H1. The H1 is read in context, with the hero
        image above it; the meta title has to work alone in a search result or a shared
        card. Name the technique and the payoff, skip the "Day N/365" prefix — it means
        nothing to someone arriving from search.
      - `meta_description` is the blurb under that title. One sentence, the specific
        insight, not "In this article we explore...". Left empty, the platform excerpts
        the opening line, which is rarely the line you'd choose.
      - There is deliberately no canonical URL field. Neither platform lets a publisher
        set one — LinkedIn's composer has no such field, and Substack treats its own domain
        as canonical, exposing only an SEO title and subtitle. The same piece does go out
        in both places, so the only lever available is the one both posts already use:
        link back to the repo prominently in the body.
      - Count characters, not bytes — an em dash is one character of the budget.
      - **Always wrap both values in double quotes.** A good meta title very often contains
        a colon ("Postorder traversal: why the node comes last"), and an unquoted colon is
        a YAML mapping separator — the file then fails to parse and takes the whole batch
        down, not just that day.

      After the front matter, the article proper: header "365 Days of LeetCode Challenge — Day
      <day>/365", question title + LeetCode link, the FULL intuition write-up (based on
      INTUITION.md, expanded for a public audience who hasn't seen INTUITION.md itself),
      the FULL solution walkthrough with the real code (based on SOLUTION.md), the
      "Builds on" links from step (e) (full GitHub URLs, since relative paths are dead on
      LinkedIn), and the companies list if `COMPANIES.md` was written. This is the actual content — the
      short post below just points people at it. Follow the "Writing style" section
      above.

      **End the body with the walkthrough link, in exactly this form:**
      ```
      Full code and the step-by-step walkthrough:
      [<folder basename>](https://github.com/architagr/leetcode_solutions/blob/main/<folder>/SOLUTION.md)
      ```
      A bare backticked repo path — `Full code: easy_problems/101_200/path_sum/ in the
      repo.` — is what this used to say, and it is dead text on LinkedIn: nothing there
      turns a path into a link, so a reader who wants the code has to go and search for it.
      The link goes to `SOLUTION.md` rather than the folder because that is the page with
      the walkthrough diagrams on it. Same block, same wording, as `POST_SUBSTACK.md`.

      Then a line of 5-8 hashtags (see hashtags note below step m), and the AI-disclosure
      line (see below) as the very last thing in the file.

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

   n. Write `<folder>/POST_X.md`. **This file IS the exact post** — `leetcodectl post-x`
      sends its bytes verbatim (trailing newline trimmed) with `HERO.png` attached.

      - **Hard cap: 280 characters**, counted as characters, not bytes. `post-x` refuses
        to send an over-length post rather than truncating it, which fails the day's cron.
        Check with `python3 -c "print(len(open('<folder>/POST_X.md').read().strip()))"`.
        X collapses any URL to 23 characters regardless of real length, so a long GitHub
        link costs less than it looks — but the tooling counts it in full, so write to the
        stricter number and you will always fit.
      - **Write it fresh. Do not truncate POST_DISCORD.md or POST_LINKEDIN.md.** This is
        the one rule here that is about the account rather than the prose: near-identical
        text fanned out across several networks is the duplicate-content pattern X's
        platform manipulation policy actually polices, and it is also just worse writing.
        Same insight, different sentence.
      - **End with one or two hashtags, on their own line at the bottom.** They do help
        reach on X — but only in ones and twos. The five-to-eight line the LinkedIn posts
        carry is a LinkedIn convention, and the same line on X reads as spam. Pick the
        ones a person actually searches: the language and the topic (`#golang #leetcode`,
        `#golang #binarytree`), not the generic career tags.
      - Budget them deliberately. A link costs 23 characters on X however long it is, so
        `leetcodectl` counts it that way too — but the tags still come out of the same 280,
        and trimming a sentence to pay for them is the right trade, not padding the post
        to the cap.
      - No image markdown — the poster attaches `HERO.png` itself.

      Structure, roughly (adapt it; a template applied 365 times stops reading as a
      person):
      ```
      Day <day>/365 · <title> (<difficulty>)

      <one or two lines: the actual insight, not a summary of the problem>

      <complexity, when it's the interesting part>

      <github blob url to that folder's SOLUTION.md>
      ```

   o. Write `<folder>/POST_SUBSTACK.md`: the long-form piece again, for the CodeStreak
      Daily Substack. Substack has no publishing API, so `leetcodectl substack-batch`
      prepares this for pasting by hand.

      - **Open with the same front matter block as step (k)** — `meta_title` and
        `meta_description` — and the same rules. `substack-batch` surfaces them as publish
        settings and warns on gaps. Substack's SEO Options map onto them directly: its SEO
        Title and SEO Subtitle are these two fields.
      - The body may reuse the article's structure and code, but rewrite the opening and
        closing paragraphs. Substack readers subscribed to a newsletter, not a feed —
        the piece can open slower and go deeper, and an opening line that reads as a
        LinkedIn hook lands badly there.
      - **Reference images by their local path, `images/<file>`, exactly as SOLUTION.md
        and the LinkedIn article do.** Not a `raw.githubusercontent.com` URL. The images
        get attached by hand when the post is pasted into Substack's editor, the same way
        they are for a LinkedIn article, so the path in the file is a pointer for the
        author rather than something the platform resolves. Keeping all three files on the
        same paths means a diagram can be renamed in one place.

        This applies to image references only. The LeetCode link, the "Builds on" links
        and the walkthrough link at the bottom are real external links and stay absolute.
      - **Tags go in the front matter, not in the body.** Substack has no inline-hashtag
        discovery — a `#golang` line in the prose is just text there. What it actually
        indexes is the post's tags, set in the publish dialog. So add up to five to the
        front matter block and `substack-batch` will surface them as a publish setting:
        ```
        tags: [golang, binary-tree, recursion, dsa]
        ```
        Pick tags that describe this piece, not the series. Past five they stop describing
        the post and start reading as keyword stuffing to someone who can see them all.
      - Carry the same AI-disclosure line as the LinkedIn article, as the last thing in
        the file.

   p. Commit everything for this question in one commit. The message body is
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
  anything itself. Posting is a separate subsystem, and it is split: Discord and X go out
  automatically on their own GitHub Actions crons (`post-discord`, `post-x`), while
  LinkedIn and Substack are prepared as paste-ready batches (`linkedin-batch`,
  `substack-batch`) and marked by hand with `mark-posted` after they actually go out.
  See `challenge/README.md` for the whole picture.
- Re-running this skill with the same URL/list is always safe — questions whose content
  is already written
  are skipped, so partial batches (e.g. "10 of 30 done this week") resume cleanly.
