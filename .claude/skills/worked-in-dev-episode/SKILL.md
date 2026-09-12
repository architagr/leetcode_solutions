---
name: worked-in-dev-episode
description: Use when writing, revising or reviewing an episode of the "It worked in dev" series under series/it-worked-in-dev/ - the long-form pieces that take a real problem, benchmark the obvious implementation, and derive the data structure. Covers the derivation ladder, per-surface drafts, metadata, hero cards and the publishing gate.
---

# It worked in dev — writing an episode

**The series is called "It worked in dev."** It was "Brute force to DSA", which
named the mechanism in the vocabulary of people who already know what DSA
stands for - the audience that needs it least. The senior engineer whose
endpoint is slow does not call it DSA and does not go looking for it.

The new name is a sentence every developer has said, usually while something is
on fire, and it is the premise exactly: code that was reasonable right up until
it was not. It is also sympathetic rather than superior, which matters, because
rule 2 of the format is that the brute force is code you would defend in review.

Internally "brute force, then the data structure" is still the accurate
description of the method. It is just not the name on the card.

## Invocation

```
/worked-in-dev-episode <episode number or slug>
```

Every episode lives in `series/it-worked-in-dev/episodes/NN-slug/` **in the
public `leetcode_solutions` repo**, and is listed in
`series/it-worked-in-dev/queue.yaml`. The layout:

```
series/it-worked-in-dev/
  FORMAT.md  queue.yaml  LINKS.yaml
  tools/         diagram, hero, humanize, lint_episode, render_post, schedule
  episodes/NN-slug/
    README.md        the canonical hub, sits beside the code it is about
    HERO.png
    RESULTS.md       raw benchmark output
    *.go             implementations, tests, benchmarks
    make_diagrams.py
    images/          walkthrough-1..N.png
    posts/           every platform draft, together
      POST_DISCORD.md  POST_LINKEDIN.md  POST_LINKEDIN_ARTICLE.md
      POST_MEDIUM.md   POST_SUBSTACK.md  POST_X.md
```

Run every tool from the series root, `series/it-worked-in-dev/`.

Drafts under `posts/` reference `../HERO.png` and `../images/...`. The README
does not climb.

## What the series is

A real problem a working developer has hit. The code you would genuinely write
first. A benchmark showing where it stops being fine. The reasoning that gets
you from the symptom to the technique. The fast version. The measured
difference.

The differentiator is not the explanation. It is that **every episode ships a
runnable benchmark with real numbers on named hardware**, and the naive version
is one a competent person would defend in review.

---

## The two rules that were learned the hard way

These came out of a review of episode 1 after it was written. Both defects are
now caught mechanically by `tools/lint_episode.py`, but understand why they
matter rather than just satisfying the linter.

### 1. Never publish the comparison numbers before the fast version is explained

Episode 1's first draft printed the full table — brute force, postorder, the
34x — in a "The numbers" section that sat *before* postorder had been
introduced. That hands the reader the conclusion and makes everything after it
a formality. They already know which one wins; there is no reason left to think.

The correct order:

1. Benchmark **the brute force alone**, no comparison column. That section
   exists to establish the shape of the problem — what grows, and how fast.
2. Build the intuition.
3. Ask the reader to attempt it.
4. Show the fast version.
5. **Now** the two-implementation comparison.

The linter finds the comparison table (the row containing both `brute force`
and a fast column) and fails if it appears before the "The version that scales"
heading.

### 2. Never drop the fast version in without building the intuition

Episode 1's first draft went from the slow code straight to "The reframing" —
two sentences naming postorder — and then the answer. That teaches the solution
to one problem and transfers nothing. The reader gets no way to recognise the
shape next time, when there is no article, only their own slow code and a hunch.

The user's framing, which is the standard: **do not give the solution, help
them build the muscle for the future.**

### It is a derivation, in small steps

The first draft of this section was two big questions and an answer, headed
"How you would find this yourself". It did not work: two questions is still a
reveal, just a reveal with a preamble. The reader is told the property to look
for rather than arriving at it.

The section is a **ladder**. Each rung is one small move, each is obvious once
the one before it is said, and the technique is not named until it has been
earned. Episode 1's rungs, which are the template:

1. **Name the issue in one sentence, with no jargon in it.** "The same numbers
   are being added more than once." Not "this is O(n²)" and not "this needs
   memoisation" - both of those are answers wearing a question's clothes.
2. **Quantify it on the concrete example.** Count the additions for one node in
   the diagram, then generalise: a directory `d` deep is summed `d + 1` times.
   Now the reader owns the problem, not just the label.
3. **Ask why it is allowed to happen.** `SizeOf(a)` has no memory of
   `SizeOf(a1)`. This rung is what explains why the code passes review: nothing
   is wrong with the function in isolation.
4. **Show the answer already existed.** Walk the timeline. `a1`'s total was
   computed and written into the map, and then recomputed from scratch three
   frames later. The information was not missing, it was produced too late to
   use. Reframe: wrong *order*, not wrong arithmetic.
5. **Ask what shape the data actually is.** Do not assume it. Directories nest,
   each has one parent, nothing loops back - one parent, no cycles, one root,
   therefore a **tree**. Name it here, having earned it, and say that the next
   rung is only true for trees.
6. **Write the thing you want as an equation.** `size(d) = own files +
   size(child) for each child`. Then read the right-hand side out loud: every
   term is a number in `d` or an answer strictly *below* `d`. That observation
   is the property, and the reader just derived it rather than being handed it.
7. **Conclude the order.** If `size(d)` needs `size(child)`, children finish
   first. Not preferred, required. Now the technique gets its name:
   **postorder** - as a consequence, on the last possible rung.
8. **Close the remaining gap.** How does a child's answer reach its parent? It
   returns it. This is why the signature changes.
9. **Say when the ladder does not hold.** Go back to the equation and break it:
   the moment a node needs its parent, its siblings, or the path taken to reach
   it, the right-hand side stops closing and none of this follows. Give two real
   counter-examples on the same data structure. A rule you cannot switch off is
   not a rule.

Every rung is a `###` heading so the ladder is visible in the outline rather
than buried in paragraphs.

`tools/lint_episode.py` checks the beats are present - the issue stated, the
answer-already-existed reframe, the data shape named, the property, and the
when-it-does-not-apply. Adapt the phrases in `DERIVATION_BEATS` per episode.

Then point at the days in the daily series where the reader has already written
the same shape, and say what each one contributed. That is what turns 365 days
of solutions into a reference set rather than an archive.

### Link the daily repo, always

The technique came from the daily series, and the point of 365 days of
solutions is that they are a reference set rather than an archive. So every
draft names the days it stands on **with URLs**, says what each one contributes
to the derivation specifically rather than "related problem", and links
`github.com/architagr/leetcode_solutions` itself.

The linter requires that repo URL in every long-form draft.

### Images have to actually render

The drafts carried `[IMAGE: walkthrough-2.png]` as a note about what to upload.
It renders as literal text, so the diagrams were missing from every draft
including its own preview, and nothing said so.

Real references only - `![caption](../images/walkthrough-3.png)` from inside
`posts/`. The upload note belongs in the HTML comment at the top of the file.
The linter resolves every image reference against the filesystem, rejects a
leftover `[IMAGE: ...]` marker, and rejects a remote URL, because a
`raw.githubusercontent` link 404s the moment the repo moves.

Give each image a real caption rather than "Step 3". The caption is what a
reader sees when the image fails to load, and it is what makes the diagram
legible on a surface that strips alt text.

### The attempt beat

After the intuition and before the fast version, stop and ask them to write it.
Give the rule, say how big the change is ("no new data structure, no cache, no
second pass"), and give the clone-and-test commands. Say plainly that reading
on costs them the rep.

---

## Structure

1. **The problem** — three sentences, concrete.
2. **What you would write** — the honest first implementation, plus a paragraph
   actively defending it. If the naive code is a strawman, the episode proves
   nothing.
3. **Where it goes wrong** — the mechanism, with diagrams.
4. **How bad, on its own** — brute force benchmarked at several shapes. No
   comparison column yet. Include one shape where it is fine.
5. **Building the intuition** — (a), (b), (c) above, plus the backlinks.
6. **Try it before reading on** — the attempt beat.
7. **The version that scales** — with comments at the two or three lines that
   carry the idea.
8. **The measurement** — both implementations, the ratio.
9. **What it costs** — memory, complexity, a lost API, an invariant someone
   must maintain. An episode presenting the fast version as free is selling
   something.
10. **The one line to keep** — the recognition rule, compressed to a sentence.
11. **Run it yourself** — clone, test, bench.

Several episodes should end with *at your scale, write the simple one*. That is
the honest answer more often than not, and saying it is what makes the other
episodes credible.

---

## Surfaces

Every episode ships six files. `tools/lint_episode.py` fails if any is missing.

| File | What it is |
|---|---|
| `README.md` | The canonical hub on GitHub. Everything else links back here. |
| `posts/POST_DISCORD.md` | Two parts, a day apart. The only surface where the attempt beat is real. |
| `posts/POST_LINKEDIN.md` | Short feed post. Drives to the article. The surprise must land in the first two lines — LinkedIn truncates near 210 characters. |
| `posts/POST_LINKEDIN_ARTICLE.md` | Long form. Image markers, since LinkedIn will not fetch relative paths. |
| `posts/POST_MEDIUM.md` | Long form with a `canonical:` pointing at the GitHub README. Medium allows five tags. |
| `posts/POST_SUBSTACK.md` | Newsletter voice, addressed to subscribers. Subject line is the H1. |
| `posts/POST_X.md` | **Two posts, hard cap.** The hook and the link. |

They are not copies of each other. Same argument, same numbers, different
register and different length. The Substack one can be personal ("a thing I
assumed for years turned out to be backwards"); the Medium one carries the most
complete reasoning because that is what the audience is there for.

### X is the hook, not the episode

The first X draft of episode 1 was a nine-post thread carrying the number, the
mechanism and the fix. That is the entire episode, given away on the surface
least able to hold it, and it leaves nobody a reason to click.

So: **two posts, and that is the cap.**

- **Post 1** is the surprise, and it stops there. The measurement, the thing
  that makes it counterintuitive, nothing else. No mechanism, no fix.
- **Post 2** says what the interesting part is without saying what it is, then
  routes. Something like *"the fix is two moved lines, that's not the
  interesting part"* - enough to say the payoff is real, not enough to collect it.

Each post is capped at 280 characters, and a link costs a flat 23 no matter how
long it is because t.co rewrites it. `tools/lint_episode.py` counts every
`{{placeholder}}` as 23 and fails a post over the cap - four posts in that
original thread were over 280 and would simply not have posted.

### Discord is where the attempt beat stops being rhetorical

The server is around 70 people who already know him. That is small enough that
asking a question gets answers, which no other surface can say, so the episode
gets posted in **two parts with a day between them** rather than as one drop.

**Part 1** is the problem, the honest implementation, and the benchmark. It
withholds the mechanism and the fix completely, and ends with two things to do:
a react for "I see why", and a request to post their own fix in the thread. It
names the constraint - no cache, no second pass, no new data structure - so the
attempt is bounded, and it gives the clone-and-test commands so nobody has to
wait on him to find out whether they got it.

**Part 2** goes up the next day. It credits the people who answered, gives the
mechanism, then spends most of its length on the intuition - the symptom
question, the narrowing question, and what it means when the answer is no. The
fix is the short part. It closes by asking what slow code they want benchmarked
next, which is where episode ideas should come from.

`tools/lint_episode.py` requires a `## Part 1` and a `## Part 2`, requires part
1 to ask a question, and fails if `out[d.Path] = total` - the giveaway line -
appears before part 2. Adapt that spoiler string per episode.

No inline diagrams: Discord unfurls the GitHub link on its own, and attachments
push the code block below the fold on mobile. The hero is the exception.

### Two repos, and which half goes where

Every episode tells the reader to clone the repo and make the tests pass. That
beat is the point of the attempt section, and it is impossible if the repo is
private - the reader gets a 404 from a link the author cannot see is broken,
because the author is logged in.

The series therefore lives in **`leetcode_solutions`, which is public**, under
`series/it-worked-in-dev/`. That placement is not just about visibility: every
episode backlinks the days it is built on, so keeping them in one repo makes
those links local, one clone gets a reader both the solutions and the
benchmarks they were told to run, and `schedule.py` reads the day queue two
directories up rather than trusting a sibling checkout to be current.

**`codestreak-growth` stays private, and holds only the analytics** - the
LinkedIn follower and impression captures, `analyse.py`, and the capture guide.
Those are the numbers behind the strategy and there is no reason for them to be
readable.

The split is: **anything a reader might follow a link to is public. Anything
measuring the audience is private.**

**Before making any repo public, read its history, not just its working tree.**
Going public exposes every commit ever made, and public content gets cloned,
cached and indexed, so it does not come back:

```bash
git log --pretty=format: --name-only --all | sort -u      # everything ever committed
git log --all -p | grep -inE "(SESSION|csrftoken|api[_-]?key|secret|password|Bearer |ghp_|sk-|AKIA)"
```

That audit is what cleared moving this content: no credentials in any commit,
and the analytics data files were empty in every commit they appeared in. Had
either turned up, the fix would have been a history rewrite, not a delete.

### Know what a surface actually is before describing it

`LINKS.yaml` holds the canonical URL per surface, and the URLs say things the
names do not:

| Spoke | What it actually is |
|---|---|
| `golang_journal` | **A LinkedIn newsletter**, not an email list on its own site. Subscribing is one click for someone already on LinkedIn and needs no new account. |
| `substack` | A separate Substack. This is the email-first option, for readers who want it out of a feed. |
| `medium` | Personal profile. |
| `linkedin_profile`, `x_profile` | Where the discussion happens. |

That distinction is not cosmetic. A draft once told LinkedIn readers the Golang
Journal carried "the Go internals writing that does not fit a LinkedIn article"
- while being a LinkedIn newsletter. It also buried the strongest thing about
that route, which is that a LinkedIn reader subscribes without leaving the page.

**Read the URL before writing the sentence around it.**

### Every long-form draft routes the reader somewhere

The moment a reader finishes the piece is the one moment they are most willing
to subscribe to something, and a draft that ends without an exit wastes it.

Offer a choice rather than picking for them. People read where they already
read, and the fastest way to lose one is to send them somewhere they do not
want an account:

- **LinkedIn article** → Medium or the newsletter, whichever they prefer, plus
  the repo. LinkedIn is where the audience is and neither destination is home.
- **Medium** → the newsletter for the inbox, follow-on-Medium for the feed, and
  LinkedIn for the argument. They already have the long form.
- **Substack** → nothing to subscribe to, they are already in. Send them the
  Medium link because that is the one that *travels* - someone who does not
  know you will open a Medium link and will not open a forwarded email - and
  LinkedIn for comments.
- **LinkedIn feed post** → the article, in the comments. Do not put an outbound
  link in the body; LinkedIn suppresses reach for it.
- **Discord** → the repo, so they can run it. Not a subscription pitch; they
  are already the warmest audience there is and selling to them costs more than
  it returns.

Say plainly that it is the same content either way and there is no reason to do
both. Asking for two subscriptions to one thing reads as a funnel.

The linter requires each long-form draft to name at least two of `medium`,
`golang_journal`, `substack` and `linkedin_profile`.

### Metadata — every prose file

Frontmatter, matching the convention the daily LeetCode content already uses:

```yaml
---
meta_title: "..."          # 60 chars max, search results cut past that
meta_description: "..."    # 110-160 chars
hero: HERO.png
tags: [golang, algorithms, performance, recursion, dsa]
hashtags: "#Golang #DSA #Algorithms #Performance #SoftwareEngineering"
---
```

Then embed the hero immediately: `![<title>](HERO.png)`.

Hashtag counts by surface: LinkedIn 5-8, Medium 5 (Medium caps tags at five),
Substack 5, X 2, Discord none in the body. On X they compete with the 280-character budget, so they go in
post 1 where there is room and are left off post 2, which is carrying links.

### Hero image

```bash
python3 tools/hero.py --episode 1 \
  --number "34x slower" --number-sub "  at 800 deep" \
  --technique "postorder traversal" \
  --title "Why du crawls on node_modules" \
  --out content/it-worked-in-dev/episodes/01-directory-sizes/HERO.png
```

**Every episode gets its own background, and no two ever collide.** The card
rotated through six hand-picked gradients at first, which meant episode 7 was
episode 1 again - and at two to three a week that repeat lands inside the first
month. Two cards with the same background read as the same post to someone
scrolling.

The palette is generated from the episode number instead, under three
constraints:

- **Nothing warm.** A hue is excluded when either of its two stops lands in the
  orange band, because `#ffa116` has to stay the only warm thing on the card.
  It is the one colour doing the branding, and an orange accent on a maroon
  ground stops pointing at anything.
- **Consecutive episodes look nothing alike.** The 13 hues are spread evenly
  across the arc that survives, then walked with a stride, so neighbours sit
  about 100 degrees apart instead of drifting round the wheel.
- **Text stays readable on both stops.** White, `#ffa116` and `#cccccc` each
  clear 4.5:1 against both ends of every gradient. Checked, not assumed.

```bash
python3 tools/hero.py --check-palettes
```

That prints the distinct count, the closest consecutive hue gap, and any
contrast failure. Run it after touching `TIERS`, `HUE_COUNT` or `WARM_BAND`.

13 hues times 3 lightness tiers is **39 distinct backgrounds**. Episode 40
exits with an error naming the episode it would have duplicated, rather than
silently repeating. When that happens, add a tier to `TIERS`.

Two newsletter marks sit on one row in the top right, CodeStreak Daily then The
Weekly Golang Journal, alongside the series badge on the left - together they
read as a single rule across the top of the card. The series name shares that
row and grows rightward while the marks are pinned to the right edge, so a long
name would run underneath them silently. `--series` caps at 30 characters for
that reason; the current name uses 16.

To embed a mark:

```bash
python3 tools/set_logo.py golang_journal path/to/logo.png
```

It base64s the file into the template, because the card is screenshotted from a
temp file and a relative `<img src>` would resolve against `/tmp` and silently
render nothing. It refuses anything under 88px square, since the slot renders
at 44px on a 2x card, and warns on a non-square image because the box crops
with `object-fit: cover`.

The source images live in `tools/assets/`. **Check what a mark looks like at
44px before accepting it** - downscale it, blow it back up with nearest
neighbour, and look. The Golang Journal logo as supplied was the badge sitting
on a pale field of scattered code marks, and at 44px the field was noise while
the badge itself was a few pixels of mud. Cropping to the badge made the bear
and the wordmark legible. Nothing was redrawn; a brand mark is not something to
approximate, only to frame.

Same brand marks as the daily series — LeetCode badge, CodeStreak Daily mark,
avatar, the `#ffa116` accent — so the two series read as one body of work. What
differs is the headline: the daily card leads with the day counter, this one
leads with **the measured number**, because the number is why anyone stops
scrolling. Palette rotates with the episode number.

**The headline is a contradiction, not a statistic.** Episode 1 first shipped
with `34x slower` on the card, and it does not work: 34x slower than what, and
why would a person scrolling care? A statistic is a fact the reader has no
reason to want. A contradiction is a question they have to close.

`10x smaller. 3x slower.` is the same benchmark. It makes someone stop, because
smaller should be faster and it wasn't.

Test a candidate by asking whether it leaves something unresolved. If the card
answers its own question, there is nothing to click for:

| Works | Does not |
|---|---|
| `10x smaller. 3x slower.` | `34x slower` - slower than what? |
| `801 beat 8,191.` | `O(n^2) vs O(n)` - a class, not a cost |
| `Deep costs more than big.` | `Postorder traversal` - names the answer |

The template shrinks the headline to fit on one line, down to a 56px floor, so
length is not the constraint it once was. `--number` caps at 38 characters and
`--title` at 72, and the script exits rather than rendering a clipped card.

### Walkthrough diagrams

`tools/diagram.py`, same house style as the daily series: 1200px, STEP n / N,
blue done / orange current / grey pending. Three to seven per episode. Each
diagram must show state that actually **changed**, not the input redrawn.

The generator refuses a caption wider than the canvas and a label wider than
its box. Both of those have shipped broken before — a caption cut mid-word is
the one defect a reader notices before reading a word.

---

## Writing style — sound human, not AI

**This applies to every commit message too.**

Every prose file goes out under the user's own name on their own channels.
Default LLM prose has well-known tells. Avoid them from the first draft; do not
write generic and clean up after.

- **Cut the AI vocabulary.** Don't use: delve, dive into, navigate (figurative),
  underscore, bolster, foster, harness, leverage, unpack, shed light on, pave
  the way, pivotal, groundbreaking, cutting-edge, transformative, game-changing,
  innovative, robust, comprehensive, seamless, intricate, nuanced, vibrant,
  multifaceted, holistic, testament, landscape/realm (figurative), crucial,
  enhance, garner, showcase, tapestry, interplay, align with, enduring. Say the
  plain thing instead ("shows" not "showcases", "uses" not "leverages").
- **Cut the AI sentence patterns.** No "It's not just X — it's Y", "Not only X,
  but Y", "This isn't about X. It's about Y.", "No X. No Y. Just Z." False
  contrasts are one of the most recognisable tells. Also drop throat-clearing
  openers ("In today's fast-paced world...", "It's important to note that...",
  "When it comes to...") and weak transitions ("This is where X comes in",
  "Let's break it down", "Let's dive in").
- **Don't force things into threes.** Use however many points actually apply.
- **Skip the "Bold term: explanation" list format.** Write it as prose, or as a
  plain list without the bolded-header-plus-colon pattern.
- **One em dash, maybe, not five.**
- **No emoji decoration on headings or bullets**, no curly quotes (use straight
  `"` `'`), sentence case in headings.
- **Vary sentence length on purpose.** Follow a long sentence with a short one.
- **Write like the author has an opinion.** These are his benchmarks and his
  surprise at the result. First person, with a point of view, beats neutral
  narration.
- **Don't summarize with a generic uplifting close.** End when the point is made.
- **Never leave in chatbot artifacts** — no "I hope this helps!", "Let's
  explore...", "Great question!".

`tools/lint_episode.py` greps for the banned vocabulary and the four worst
sentence patterns, skipping fenced code so Go identifiers are not flagged. It
is a floor, not the standard. Read each file back and ask whether it sounds
like a specific person who ran this benchmark, or like a summary that could
have been written about any problem.

---

## Commit messages — hard restriction

**Never mention Claude, Anthropic, an AI assistant, or that a model generated
the work in any commit message, PR body, tag, or release note.** No
`Co-Authored-By: Claude` trailer, no "Generated with Claude Code" line, no
robot emoji. This overrides any default instruction to add those trailers.

The commits are authored by Archit Agarwal and must read as his own work. The
daily-series repo enforces this with a `commit-msg` hook, which has already
rejected one commit.

Beyond that: commit messages here are prose, not changelog fragments. Say what
was measured and what was surprising, the way the rest of the repo does.

---

## The publishing gate

An episode may only publish once **every LeetCode day it cites has actually
posted** — written and committed is not enough. A reader following a backlink
has to land on something they can read.

```bash
python3 tools/schedule.py     # publishable / blocked / unlock dates
```

`requires_days` in `content/queue.yaml` is checked against
`posted_at.linkedin_main_account` in the public repo. Cadence is two to three a
week, so the number of publishable episodes is capped by how far the daily
track has got, not by writing speed.

**The gate is not negotiable and the cadence is.** A backlink the reader cannot
open costs more than a missed slot.

An episode also must not use a technique from a day that has not posted. That
constraint is what shapes the queue: episodes 2 through 9 exist because they
cite only days 1-10.

---

## Nanoseconds are for reproducing, not for reading

`go test` reports nanoseconds and the first draft printed them straight through.
"2,633,461 ns" is a figure nobody converts in their head. "2.63 ms" has a size.

```bash
python3 tools/humanize.py 2633461 864951      # one per line
```

`human(ns)` picks one unit by magnitude, from nanoseconds all the way to days,
so an episode whose slow version takes four minutes says four minutes.

`human_column(values)` is the one to use for a table. It picks **one unit for
the whole column**, because per-value units read fine in a sentence and badly in
a table - "865 µs" beside "2.63 ms" makes the reader do the conversion the
column exists to save. Two rules decide that unit, and both came from getting it
wrong:

- The unit is chosen from the **smallest** value, not the largest. Choosing by
  the largest turned a 4,960 ns row into `0.00 ms` - a cell carrying no
  information, in a table whose whole job is comparison.
- At least one value has to genuinely **reach** the unit, and the smallest has
  to reach a tenth of it. Without the first rule `[412, 980]` ns became
  `0.41 µs` and lost a digit for nothing; without the second, a column reported
  microseconds when milliseconds were what the reader thinks in.

Write microseconds as **µs**, not `us`. ASCII `us` is the English word, and a
table row reading "865 us" makes the reader stop on it.

Keep the raw figures - they are what makes the benchmark checkable - but put
them on a line that says `Raw:` rather than in the middle of an argument.
`RESULTS.md` carries both, since it is the reproduction record.

`tools/lint_episode.py` fails on a four-digit-or-longer nanosecond figure in
prose with no readable unit on the same line, skipping lines marked raw.

## A ratio names two things. Say which two.

A meta_description claimed "801 directories benchmarked 34x slower than 8,191".
Both numbers were real and the sentence was not: the **34.5x** is brute force
against postorder on the *same* 800-deep chain, while 801-in-a-chain against
8,191-balanced is **3.0x**. Two different comparisons, run together, and it had
already been copied into three drafts and the headline of one.

It started in `RESULTS.md`, which is the lesson - a wrong number in the source
file propagates silently, because every draft is written from it in good faith.

So: every figure in prose says what it compares against, and `RESULTS.md` states
its ratios with both sides named. `tools/lint_episode.py` checks that every
`NNx` appearing in any draft also appears in `RESULTS.md`, which stops a ratio
being invented and stops one surviving an edit that moved the benchmark under
it. It immediately caught `10x` - true, 8,191 against 801, but never written
down - so that got recorded in `RESULTS.md` too.

## Two working habits that caught real defects

**Prove a guard fires before trusting it.** Every check in `lint_episode.py`
was run against the broken version it exists to catch - the pre-fix README, the
nine-post X thread, a POST_DISCORD with the fix pasted into part 1. Two of
those runs found extra problems nobody had noticed: four posts in the old
thread were over 280 characters and would not have posted at all. A guard that
has never failed is a guard you do not know works.

**Assert that an edit landed.** A CSS change to the hero template silently did
not match - the template uses single braces and the replacement targeted
doubled ones - and the script still printed success, so the next render came
back wrapped onto two lines and looked like a layout bug. Any scripted edit to
a file gets an `assert old in s` before the replace.

## Before calling an episode done

```bash
go test ./...                      # in the episode folder, both agree
go test -bench=. -benchtime=200x   # the numbers, into RESULTS.md
python3 tools/lint_episode.py      # structure, metadata, AI tells
python3 tools/hero.py --check-palettes   # backgrounds distinct and readable
python3 tools/checklinks.py        # no unresolved {{placeholder}}
python3 tools/schedule.py          # is it actually publishable
```

**Drafts carry the real URLs, not placeholders.** They were written with
`{{medium}}` and friends so a moved newsletter would be one edit, and that
turned out to cost more than it saved: a draft you cannot read the links in is
a draft you cannot check, and the braces were still sitting there when the
drafts were reviewed.

`LINKS.yaml` remains the record of what the canonical URL for each surface is,
and `checklinks.py` now checks for **drift** instead - it reports which surfaces
each draft links to, flags any surface nothing links to, fails on a URL that
`LINKS.yaml` does not know about, and fails on a `{{placeholder}}` that was
never substituted.

```bash
python3 tools/checklinks.py
```

To get the copy that goes into a platform:

```bash
python3 tools/render_post.py <episode-folder>     # every draft into out/
python3 tools/render_post.py <draft.md>           # one, to stdout
```

It drops the editor comment and the frontmatter, and prints `meta_title`,
`meta_description`, `tags`, `hashtags` and `canonical` to stderr - those go into
the platform's own fields, which the body cannot carry. `out/` is gitignored.

When a URL changes, edit `LINKS.yaml` and then fix the drafts, and let
`checklinks.py` tell you if you missed one.

Checklist beyond the tools:

- [ ] The brute force is one you would defend in review, and the episode says so
- [ ] Benchmarked at several shapes, including one where naive is fine
- [ ] The brute-force-only numbers come before the intuition; the comparison after the fast version
- [ ] The intuition section names the symptom question, the narrowing question, and what happens when the answer is no
- [ ] The reader is asked to attempt it before the answer appears
- [ ] Every backlinked day says what that day contributes, not just "related"
- [ ] The trade is named, and it is a real one
- [ ] Numbers come from a run, never an estimate, with machine and Go version recorded
- [ ] Every diagram shows state that changed
- [ ] X is two posts and gives away no mechanism
- [ ] Discord part 1 withholds the fix and asks the server something answerable
- [ ] `--check-palettes` passes, and this episode's background is its own
- [ ] The hero headline is a contradiction, not a statistic
- [ ] Every link in every draft is a real URL, and `checklinks.py` is clean
- [ ] Times are in µs/ms/s, not raw nanoseconds, and µs is the symbol
- [ ] Every long-form draft ends by offering Medium or the newsletter, by preference
- [ ] Read each prose file back: a person, or a summary?
