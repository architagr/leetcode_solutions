#!/usr/bin/env python3
"""Refuse an episode that is structurally or editorially wrong.

Three classes of defect, all of which have actually shipped:

1. **Numbers before the reframing.** Episode 1 originally printed the
   comparison table - both implementations, the 34x - before the fast version
   had been explained. That hands the reader a conclusion and skips the part
   that transfers. The comparison must come after the fast implementation.

2. **The fast version dropped in without intuition.** Going straight from slow
   code to fast code teaches the answer to this one problem and nothing else.
   The reader gets no way to recognise the pattern next time, when there is no
   article - only their own slow code and a hunch. Every episode must carry a
   section that reasons from symptom to technique, and must ask the reader to
   attempt it before showing the answer.

3. **X spending the argument.** The first X draft was a nine-post thread
   carrying the number, the mechanism and the fix. That is the whole episode,
   given away on the surface least able to hold it, and it leaves nobody a
   reason to click. X is the top of the funnel: two posts, the surprise and
   the link.

4. **No route off the page.** A long-form draft that ends without sending the
   reader somewhere they can subscribe wastes the one moment they are most
   willing to. The routing offers a choice - Medium or the newsletter,
   whichever they already read in - rather than picking for them.

5. **AI tells.** These go out under a person's name. The vocabulary and
   sentence patterns below are the ones that make prose read as machine
   written; the list is the one the daily LeetCode skill enforces.

Run:  python3 tools/lint_episode.py [episode-folder ...]
"""
import os
import re
import sys

import yaml

# The series root, one level up from tools/.
SERIES = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
EPISODES = os.path.join(SERIES, "episodes")

# Every surface an episode publishes to. The README is the canonical hub and
# sits at the episode root beside the code it is about; the platform drafts all
# live together under posts/ so the folder does not read as nine loose files.
REQUIRED_FILES = [
    "README.md",
    "HERO.png",
    "posts/POST_DISCORD.md",
    "posts/POST_LINKEDIN.md",
    "posts/POST_LINKEDIN_ARTICLE.md",
    "posts/POST_MEDIUM.md",
    "posts/POST_SUBSTACK.md",
    "posts/POST_X.md",
]

# Frontmatter keys every prose file carries. meta_title and meta_description
# are what a link preview and a search result show - absent, the platform
# invents them from the first sentence, which is never the right sentence.
REQUIRED_FRONTMATTER = ["meta_title", "meta_description", "hero", "hashtags"]

META_TITLE_MAX = 60        # Google truncates around here
META_DESC_MIN = 110
META_DESC_MAX = 160

# X caps a post at 280 characters, and every link costs a flat 23 regardless of
# its real length, because t.co rewrites it. So a URL counts as 23 here, not as
# the 80 characters a LinkedIn newsletter link actually runs to.
X_POST_MAX = 280
X_POSTS_MAX = 2
TCO_LEN = 23

# The surfaces a reader can be sent to. A long-form draft has to name at least
# two so the reader picks the one they already read in.
ROUTE_KEYS = ["medium", "golang_journal", "substack", "linkedin_profile"]
ROUTE_MIN = 2
LONG_FORM = ["posts/POST_LINKEDIN_ARTICLE.md", "posts/POST_MEDIUM.md",
             "posts/POST_SUBSTACK.md"]

# The daily repo is where the technique was worked out, and the whole point of
# 365 days of solutions is that they are a reference set. A draft that cites a
# day without linking it wastes the backlink and the credibility.
SOLUTIONS_REPO = "github.com/architagr/leetcode_solutions"

# The line that gives the whole thing away. Discord part 1 is a question the
# server answers, so the fix must not appear above the part 2 heading.
DISCORD_SPOILER = "out[d.Path] = total"

AI_WORDS = [
    "delve", "dive into", "navigating", "underscore", "bolster", "foster",
    "harness", "leverage", "unpack", "shed light", "pave the way", "pivotal",
    "groundbreaking", "cutting-edge", "transformative", "game-changing",
    "innovative", "robust", "comprehensive", "seamless", "intricate",
    "nuanced", "vibrant", "multifaceted", "holistic", "testament",
    "crucial", "enhance", "garner", "showcase", "tapestry", "interplay",
    "align with", "enduring", "let's dive", "let's break it down",
    "in today's fast-paced", "it's important to note", "when it comes to",
    "this is where", "i hope this helps", "great question",
]

AI_PATTERNS = [
    (r"[Ii]t's not just .{1,40} - it's", "the 'not just X, it's Y' construction"),
    (r"[Ii]t's not just .{1,40} — it's", "the 'not just X, it's Y' construction"),
    (r"[Nn]ot only .{1,40}, but", "the 'not only X but Y' construction"),
    (r"isn't about .{1,40}\. It's about", "the 'not about X, it's about Y' construction"),
    (r"[“”‘’]", "curly quotes - use straight ones"),
]

# The section headings that carry the reasoning. An episode needs one of the
# intuition headings and one of the attempt headings.
INTUITION_HEADINGS = ["from the symptom to the shape", "building the intuition",
                      "the part i want you to take away"]

# The derivation has to actually derive. These are the beats it walks through -
# name the repetition, notice the answer already existed, name the data's
# shape, write the equation, conclude the order. An episode missing the middle
# of that is announcing the technique rather than reaching it.
DERIVATION_BEATS = [
    ("the issue", ["the issue, said plainly", "the issue is", "said plainly"]),
    ("already had the answer", ["already had the answer", "already existed",
                                "we already computed"]),
    ("name the data shape", ["what shape is this data", "that is a tree",
                             "shape is this"]),
    ("the property", ["depend only on the answers for its children",
                      "depends only on its children"]),
    ("when it does not apply", ["when this does not apply", "does not apply",
                                "no longer closes"]),
]
ATTEMPT_HEADINGS = ["try it before", "stop here and write", "stop and write",
                    "try it before you scroll", "stop here and try"]


def frontmatter(text):
    if not text.lstrip().startswith("---") and "\n---\n" not in text[:1200]:
        return None
    m = re.search(r"^---\n(.*?)\n---\n", text, re.S | re.M)
    return m.group(1) if m else None


def check_metadata(path, errors):
    text = open(path).read()
    fm = frontmatter(text)
    if fm is None:
        errors.append("%s: no frontmatter block" % path)
        return
    for key in REQUIRED_FRONTMATTER:
        if not re.search(r"^%s:" % key, fm, re.M):
            errors.append("%s: frontmatter missing %s" % (path, key))
    m = re.search(r'^meta_title:\s*"(.*)"', fm, re.M)
    if m and len(m.group(1)) > META_TITLE_MAX:
        errors.append("%s: meta_title is %d chars, over %d - search results cut it"
                      % (path, len(m.group(1)), META_TITLE_MAX))
    m = re.search(r'^meta_description:\s*"(.*)"', fm, re.M)
    if m:
        n = len(m.group(1))
        if not META_DESC_MIN <= n <= META_DESC_MAX:
            errors.append("%s: meta_description is %d chars, want %d-%d"
                          % (path, n, META_DESC_MIN, META_DESC_MAX))
    if "#" not in (re.search(r"^hashtags:.*", fm, re.M) or
                   type("", (), {"group": lambda s, *a: ""})()).group(0):
        errors.append("%s: hashtags line carries no tags" % path)
    if not re.search(r"!\[[^\]]*\]\((\.\./)?HERO\.png\)", text):
        errors.append("%s: does not embed HERO.png" % path)


def check_prose(path, errors):
    text = open(path).read()
    # Skip fenced code - Go identifiers are not prose and must not be flagged.
    prose = re.sub(r"```.*?```", "", text, flags=re.S)
    prose = re.sub(r"<!--.*?-->", "", prose, flags=re.S)
    low = prose.lower()
    for word in AI_WORDS:
        if word in low:
            errors.append("%s: AI tell %r" % (path, word))
    for pattern, why in AI_PATTERNS:
        if re.search(pattern, prose):
            errors.append("%s: %s" % (path, why))


def check_structure(path, errors):
    """The ordering rule: intuition, then attempt, then fast code, then numbers."""
    text = open(path).read()
    low = text.lower()

    if not any(h in low for h in INTUITION_HEADINGS):
        errors.append("%s: no intuition section. The reader needs the reasoning "
                      "from symptom to technique, not just the fast code." % path)
    missing = [name for name, phrases in DERIVATION_BEATS
               if not any(x in low for x in phrases)]
    if missing:
        errors.append(
            "%s: the intuition section skips %s. It has to walk from the "
            "symptom to the technique, not announce the technique."
            % (path, ", ".join(missing)))
    if not any(h in low for h in ATTEMPT_HEADINGS):
        errors.append("%s: never asks the reader to attempt it before showing "
                      "the answer. Reading the answer costs them the rep." % path)

    # The comparison table is the one with both columns. Find where it sits
    # relative to the fast implementation.
    # Match the HEADER of the two-implementation table, not any row that
    # happens to have enough pipes. Adding a human-units column to the
    # brute-force-only table gave it five pipes and tripped a pipe count.
    comparison = None
    for m in re.finditer(r"^\|[^\n]*\|", text, re.M):
        row = m.group(0).lower()
        if "brute force" in row and ("postorder" in row or "ratio" in row):
            comparison = m.start()
            break
    fast = low.find("the version that scales")
    if comparison is not None and fast != -1 and comparison < fast:
        errors.append(
            "%s: the two-implementation comparison table appears at offset %d, "
            "before the fast version is explained at offset %d. Publishing the "
            "numbers first hands over the conclusion and skips the reasoning."
            % (path, comparison, fast))


def check_x(path, errors):
    """X is the hook, not the episode."""
    text = open(path).read()
    text = re.sub(r"<!--.*?-->", "", text, flags=re.S)
    text = re.sub(r"^---\n.*?\n---\n", "", text, flags=re.S)

    posts = re.split(r"^\*\*\d+/\d+\*\*.*$", text, flags=re.M)[1:]
    if len(posts) > X_POSTS_MAX:
        errors.append(
            "%s: %d posts, cap is %d. X is the top of the funnel - spending the "
            "argument there leaves nobody a reason to click through."
            % (path, len(posts), X_POSTS_MAX))
    for i, post in enumerate(posts, 1):
        body = post.strip().strip("-").strip()
        body = re.sub(r"!\[[^\]]*\]\([^)]*\)", "", body)      # image embeds
        body = re.sub(r"\[attach [^\]]*\]", "", body)          # attach markers
        counted = re.sub(r"\{\{\w+\}\}|https?://\S+", "x" * TCO_LEN, body).strip()
        if len(counted) > X_POST_MAX:
            errors.append("%s: post %d is %d characters, over %d"
                          % (path, i, len(counted), X_POST_MAX))


def check_provenance(path, errors):
    """A cited day has to be a link, not a name."""
    text = open(path).read()
    if SOLUTIONS_REPO not in text:
        errors.append(
            "%s: never links %s. The episode is built on days from the daily "
            "series - cite them with URLs so a reader can go and read them."
            % (path, SOLUTIONS_REPO))


def check_route(path, errors):
    """Every long-form draft ends by sending the reader somewhere."""
    text = open(path).read()
    links = yaml.safe_load(open(os.path.join(SERIES, "LINKS.yaml")))
    spokes = links.get("spokes") or {}
    found = {k for k in ROUTE_KEYS
             if spokes.get(k) and spokes[k].rstrip("/") in text}
    if len(found) < ROUTE_MIN:
        errors.append(
            "%s: names %d of the places a reader can follow this (%s), want at "
            "least %d. End by offering the choice - Medium or the newsletter - "
            "rather than picking for them."
            % (path, len(found), ", ".join(sorted(found)) or "none", ROUTE_MIN))


def check_units(path, errors):
    """Nanoseconds are for reproducing a benchmark, not for reading one.

    go test reports ns and the first draft printed them straight through.
    "2,633,461 ns" is a figure nobody converts in their head; "2.63 ms" has a
    size. Raw values still belong in the episode - they are what makes the
    numbers checkable - but on a line that says so, not in the middle of an
    argument.

    tools/humanize.py does the conversion, and picks one unit for a whole
    column so rows compare without the reader doing arithmetic.
    """
    text = re.sub(r"```.*?```", "", open(path).read(), flags=re.S)
    for line in text.splitlines():
        if re.search(r"\braw\b", line, re.I):
            continue
        for m in re.finditer(r"(\d[\d,]{3,})\s*ns\b", line):
            if not re.search(r"\d\s*(µs|us|ms|s|min|hours)\b", line):
                errors.append(
                    "%s: %r is raw nanoseconds with no readable unit beside "
                    "it. Convert with tools/humanize.py, and put bare ns on a "
                    "line marked as the raw figures."
                    % (path, m.group(0)))


def check_ratios(path, folder, errors):
    """Every NNx figure in prose has to exist in RESULTS.md.

    A meta_description claimed "801 directories benchmarked 34x slower than
    8,191", which ran two different comparisons together: the 34.5x is brute
    against postorder on one shape, while 801-chain against 8,191-balanced is
    3.0x. Both numbers were real, the sentence was not. Pinning every ratio to
    the benchmark file at least stops one being invented or surviving an edit
    that changed it.
    """
    results = os.path.join(folder, "RESULTS.md")
    if not os.path.exists(results):
        return
    measured = set(re.findall(r"(\d+(?:\.\d+)?)x", open(results).read()))
    measured |= {m.split(".")[0] for m in measured}
    text = re.sub(r"```.*?```", "", open(path).read(), flags=re.S)
    for claim in set(re.findall(r"(\d+(?:\.\d+)?)x\b", text)):
        if claim not in measured and claim.split(".")[0] not in measured:
            errors.append(
                "%s: claims %sx, which is not in RESULTS.md. Either the number "
                "is invented or the benchmark moved under it." % (path, claim))


def check_images(path, folder, errors):
    """Every image reference resolves, and no placeholder markers survive.

    The drafts used to carry [IMAGE: walkthrough-2.png] as a note-to-self about
    what to upload. It renders as literal text, so the diagrams were missing
    from every draft including its own GitHub preview, and nothing said so.
    """
    text = open(path).read()
    for marker in re.findall(r"\[IMAGE:[^\]]*\]", text):
        errors.append("%s: %s is a placeholder, not an image reference. Use "
                      "![caption](../images/...) so the draft renders."
                      % (path, marker))
    for ref in re.findall(r"!\[[^\]]*\]\(([^)]+)\)", text):
        if ref.startswith(("http://", "https://", "{{")):
            errors.append("%s: image %r is a remote URL. Local paths only - a "
                          "raw.githubusercontent link 404s once the repo moves."
                          % (path, ref))
            continue
        target = os.path.normpath(os.path.join(os.path.dirname(path), ref))
        if not os.path.exists(target):
            errors.append("%s: image %r resolves to %s, which does not exist"
                          % (path, ref, os.path.relpath(target, folder)))


def check_discord(path, errors):
    """Part 1 asks the server; part 2 answers it. Not one broadcast."""
    text = open(path).read()
    parts = re.findall(r"^## Part (\d)", text, re.M)
    if parts != ["1", "2"]:
        errors.append(
            "%s: wants a 'Part 1' and a 'Part 2' heading, found %r. Discord is "
            "the one surface where the attempt beat can be real - post the "
            "benchmark, let people answer, reveal the next day." % (path, parts))
        return
    first = text[:text.index("## Part 2")]
    if DISCORD_SPOILER in first:
        errors.append(
            "%s: part 1 contains %r, which is the fix. Part 1 has to withhold "
            "it or there is nothing to answer." % (path, DISCORD_SPOILER))
    if not re.search(r"\?", first):
        errors.append("%s: part 1 asks the server nothing" % path)


def check_episode(folder, errors):
    name = os.path.basename(folder)
    for f in REQUIRED_FILES:
        if not os.path.exists(os.path.join(folder, f)):
            errors.append("%s: missing %s" % (name, f))
    for f in REQUIRED_FILES:
        path = os.path.join(folder, f)
        if not f.endswith(".md") or not os.path.exists(path):
            continue
        check_metadata(path, errors)
        check_prose(path, errors)
        check_images(path, folder, errors)
        check_ratios(path, folder, errors)
        check_units(path, errors)
        if f in LONG_FORM:
            check_route(path, errors)
            check_provenance(path, errors)
        if f == "posts/POST_X.md":
            check_x(path, errors)
        if f == "posts/POST_DISCORD.md":
            check_discord(path, errors)
    readme = os.path.join(folder, "README.md")
    if os.path.exists(readme):
        check_structure(readme, errors)


def main():
    folders = sys.argv[1:] or sorted(
        os.path.join(EPISODES, d) for d in os.listdir(EPISODES)
        if os.path.isdir(os.path.join(EPISODES, d)))
    errors = []
    for folder in folders:
        check_episode(folder, errors)
        print("checked", os.path.basename(folder))
    if errors:
        print()
        for e in errors:
            print("  FAIL", e)
        sys.exit(1)
    print("\nall episodes pass")


if __name__ == "__main__":
    main()
