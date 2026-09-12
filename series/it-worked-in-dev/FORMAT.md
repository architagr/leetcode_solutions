# It worked in dev — the format

A real problem a working developer has actually hit. The code you would genuinely
write first. A benchmark showing where it stops being fine. The reframing. The fast
version. The measured difference.

## Why this exists

The daily LeetCode series proves consistency and teaches technique. It does not answer
the question most working developers actually have, which is *when does any of this
matter to me*.

Almost every DSA explainer answers that with complexity classes. O(n²) versus O(n log n)
is true and it is abstract, and abstract does not change behaviour. A number does.

So the differentiator here is not the explanation. It is that **every episode ships a
benchmark with real numbers on real hardware**, and the naive version is one a competent
person would defend in review.

## The rules

**1. The problem is one a developer has hit.**
Not "given an array of integers". A slow endpoint, a build that will not resolve, a
report that times out. If it needs a story to sound real, it is not real enough.

**2. The brute force is genuinely reasonable.**
It must be the version a good engineer writes first, not a strawman. If the naive code
is obviously stupid, the episode proves nothing. The point is that reasonable code has a
scale at which it stops working.

**3. The benchmark is real and reproducible.**
Go's `testing.B`, committed, runnable. Numbers come from a run, never from an estimate.
Every episode records the machine and the Go version, because a number without them is
decoration.

**4. Show where it breaks, not just that it does.**
Benchmark at several sizes. The interesting fact is rarely "the fast one wins" — it is
that they are indistinguishable at n=100 and 2000x apart at n=1,000,000. The crossover
is the lesson.

**5. Name the trade.**
The fast version costs something: memory, complexity, a dependency, an invariant someone
must maintain. An episode that presents it as free is selling something.

**6. Link to the day that teaches the technique.**
Every episode points at the LeetCode write-up covering the underlying method. That is
what makes 81 days of published work a reference library rather than an archive.

**7. Build the intuition before showing the fast version.**
Going from the slow code straight to the fast code teaches the answer to one problem and
transfers nothing. The reader has to leave able to recognise the shape next time, when
there is no article - only their own slow code and a hunch. So every episode reasons from
the symptom to the technique, gives the one question that narrows it, says what happens
when the answer to that question is no, and then asks the reader to attempt it before the
answer appears.

Not the solution. The muscle.

**8. The comparison numbers come after the fast version, never before.**
Benchmark the brute force on its own first - that establishes the shape of the problem.
Publishing the two-implementation table early hands over the conclusion and makes
everything after it a formality.

## Structure

1. **The problem** - three sentences, concrete.
2. **What you would write** - the honest first implementation, plus a paragraph
   defending it.
3. **Where it goes wrong** - the mechanism, with diagrams.
4. **How bad, on its own** - the brute force benchmarked at several shapes. No
   comparison column yet.
5. **Building the intuition** - symptom, the question that narrows it, and what it
   means when the answer is no. Plus the days where the reader already wrote this shape.
6. **Try it before reading on** - the attempt beat.
7. **The version that scales** - with the trade stated.
8. **The measurement** - both implementations, the ratio.
9. **What it costs**.
10. **The one line to keep** - the recognition rule in a sentence.

## Surfaces

Six files per episode: `README.md` (the canonical hub), `POST_LINKEDIN.md`,
`POST_LINKEDIN_ARTICLE.md`, `POST_MEDIUM.md`, `POST_SUBSTACK.md`, `POST_X.md`.
Same argument and same numbers in each, different register and length.

Every prose file carries frontmatter with `meta_title`, `meta_description`, `hero`,
`tags` and `hashtags`, and embeds `HERO.png` - the same convention the daily LeetCode
content uses. A platform left to invent its own preview text never picks the right
sentence.

`tools/hero.py` renders the card, generating a background from the episode
number so no two episodes ever share one - six hand-picked gradients meant
episode 7 repeated episode 1, which at this cadence is inside the first month. It carries the daily series' brand marks so the two
read as one body of work, but leads with the measured number instead of a day counter,
because the number is why anyone stops scrolling.

`tools/lint_episode.py` enforces the structure rules, the metadata, and a list of AI
tells. `.claude/skills/worked-in-dev-episode/SKILL.md` is the full operational guide.

## What this is not

**Not a benchmark shootout.** The numbers serve the idea; the idea is not "Go is fast".

**Not premature optimisation advocacy.** Several episodes should end with *at your scale,
write the simple one*. That is the honest answer more often than not, and saying it is
what makes the other episodes credible.

**Not a rewrite of the LeetCode post.** That one teaches the technique. This one shows
the bill arriving.

## Cadence and the gate

Two to three episodes a week, scheduled ahead.

That runs into the prerequisite rule directly: an episode may only publish once
every LeetCode day it cites has actually **posted**, and the daily series
publishes one day at a time. So the number of publishable episodes is capped by
how far the daily track has got, not by how fast episodes can be written.

`tools/schedule.py` reports that cap. Episodes 2 through 9 all sit inside days
1-10, which is roughly three weeks of runway at three a week — by which point the
daily series has reached the high twenties and the level-order arc has unlocked
the next batch.

**The gate is not negotiable and the cadence is.** A backlink to something the
reader cannot open costs more than a missed publishing slot.
