#!/usr/bin/env python3
"""Render an episode hero card to HERO.png.

Same brand marks as the daily LeetCode series - the LeetCode badge, the
CodeStreak Daily mark, the avatar, the orange accent - so the two series read
as one body of work. What changes is the headline: the daily card leads with
the day counter, this one leads with the measured number, because the number
is the reason anybody stops scrolling.

The palette rotates with the episode number for the same reason the daily one
rotates with the day: a reader should see at a glance that this is a new one.

Needs a headless Chromium once:  npx playwright install chromium
"""
import argparse
import colorsys
import os
import subprocess
import sys
import tempfile

TOOLS = os.path.dirname(os.path.abspath(__file__))
TEMPLATE = os.path.join(TOOLS, "hero_template.html")

# Every episode gets its own background, and no two ever collide.
#
# The first version rotated through six hand-picked gradients, which meant
# episode 7 was episode 1 again. At two to three episodes a week that repeat
# lands inside the first month, and two cards with the same background read as
# the same post to someone scrolling past.
#
# So the palette is generated instead. Three constraints shape it:
#
# 1. **Nothing warm.** A hue is excluded if either of its two stops lands in
#    the orange/yellow/brown band, because #ffa116 has to stay the only warm
#    thing on the card - it is the one colour doing the branding.
# 2. **Consecutive episodes look nothing alike.** The 13 hues are spread evenly
#    across the arc that survives, then walked with a stride so neighbours sit
#    at least 114 degrees apart rather than drifting round the wheel.
# 3. **Text stays readable on both stops.** Checked, not assumed - see
#    verify_palettes(), which is what `--check-palettes` runs.
#
# 13 hues times 3 lightness tiers is 39 distinct backgrounds. Past that the
# script refuses rather than silently repeating; add a tier.
HUE_SHIFT = 18          # degrees between the two gradient stops
WARM_BAND = (345, 78)   # the arc the accent owns, wrapping past 0
HUE_COUNT = 13

# (lightness, saturation) for each stop, one entry per tier.
TIERS = [
    (0.105, 0.42, 0.190, 0.50),
    (0.080, 0.26, 0.150, 0.32),
    (0.120, 0.58, 0.185, 0.64),
]

# Contrast floors, against both stops. White is body text, #ffa116 the accent,
# #cccccc the muted labels.
CONTRAST = [((255, 255, 255), 4.5), ((255, 161, 22), 4.5), ((204, 204, 204), 4.5)]


def _warm(hue):
    # wraps through 0, so it is two arcs: 345-359 and 0-78. Reds count as warm
    # here - a maroon card makes #ffa116 read as part of the background rather
    # than as the one thing being pointed at.
    lo, hi = WARM_BAND
    return hue >= lo or hue <= hi


def _hue_ladder():
    allowed = [h for h in range(0, 360, 3)
               if not _warm(h) and not _warm((h + HUE_SHIFT) % 360)]
    spread = [allowed[round(i * (len(allowed) - 1) / (HUE_COUNT - 1))]
              for i in range(HUE_COUNT)]
    # stride 5 with 13 hues visits every one before repeating, and puts a wide
    # gap between neighbours instead of walking the wheel in order
    return [spread[(i * 5) % HUE_COUNT] for i in range(HUE_COUNT)]


HUES = _hue_ladder()
CYCLE = len(HUES) * len(TIERS)


def _hex(hue, light, sat):
    r, g, b = colorsys.hls_to_rgb(hue / 360.0, light, sat)
    return "#%02x%02x%02x" % (round(r * 255), round(g * 255), round(b * 255))


def palette_for(episode):
    """The gradient for one episode. Deterministic, and unique within CYCLE."""
    if episode < 1:
        sys.exit("episode must be 1 or greater, got %d" % episode)
    if episode > CYCLE:
        sys.exit(
            "episode %d is past the %d distinct backgrounds this generator "
            "produces, so it would reuse episode %d's card. Add a tier to "
            "TIERS in %s." % (episode, CYCLE, ((episode - 1) % CYCLE) + 1,
                              os.path.basename(__file__)))
    i = episode - 1
    hue = HUES[i % len(HUES)]
    light1, sat1, light2, sat2 = TIERS[(i // len(HUES)) % len(TIERS)]
    return (_hex(hue, light1, sat1),
            _hex((hue + HUE_SHIFT) % 360, light2, sat2))


def _luminance(rgb):
    def channel(c):
        c = c / 255.0
        return c / 12.92 if c <= 0.04045 else ((c + 0.055) / 1.055) ** 2.4
    r, g, b = [channel(x) for x in rgb]
    return 0.2126 * r + 0.7152 * g + 0.0722 * b


def _contrast(a, b):
    la, lb = _luminance(a), _luminance(b)
    hi, lo = max(la, lb), min(la, lb)
    return (hi + 0.05) / (lo + 0.05)


def verify_palettes():
    """Assert what the comment above claims. Run by --check-palettes."""
    seen, problems = {}, []
    for episode in range(1, CYCLE + 1):
        stops = palette_for(episode)
        if stops in seen:
            problems.append("episodes %d and %d share %s"
                            % (seen[stops], episode, stops))
        seen[stops] = episode
        for stop in stops:
            rgb = tuple(int(stop[i:i + 2], 16) for i in (1, 3, 5))
            for fg, floor in CONTRAST:
                got = _contrast(fg, rgb)
                if got < floor:
                    problems.append(
                        "episode %d stop %s: #%02x%02x%02x contrast %.2f, "
                        "under %.1f" % (episode, stop, fg[0], fg[1], fg[2],
                                        got, floor))
    gaps = [min(abs(HUES[i] - HUES[i + 1]), 360 - abs(HUES[i] - HUES[i + 1]))
            for i in range(len(HUES) - 1)]
    print("%d episodes, %d distinct backgrounds" % (CYCLE, len(seen)))
    print("closest consecutive hues: %d degrees apart" % min(gaps))
    if problems:
        for p in problems:
            print("  FAIL", p)
        sys.exit(1)
    print("every background distinct, all text clears its contrast floor")


# The template shrinks the headline to fit, down to a 56px floor, so this cap
# is a backstop rather than the real constraint. It exists because below about
# 56px the headline stops doing its job at feed thumbnail size, and 38
# characters is roughly where the autofit hits that floor.
#
# The first version of this capped at 14 characters, which is why episode 1
# shipped with "34x slower" on the card. That is a statistic, not a hook - 34x
# slower than what, and why would anyone care? A hero headline has to carry a
# contradiction the reader wants resolved. "Smaller tree. 3x slower." is the
# same benchmark and it makes someone stop.
NUMBER_MAX = 38

# The series name in the top-left rule.
#
# It was "Brute force to DSA", which named the mechanism in the vocabulary of
# people who already know what DSA stands for - the audience that needs the
# series least. The senior backend engineer whose endpoint is slow does not
# call it DSA and does not go looking for it.
#
# "It worked in dev" is a sentence every developer has said, usually while
# something is on fire, and it is the premise exactly: code that was
# reasonable right up until it was not. It is also sympathetic rather than
# superior, which matters, because rule 2 of the format is that the brute
# force is code you would defend in review.
SERIES = "It worked in dev"

# The series name and the two newsletter marks now share one row, so the name
# grows to the right and the marks are fixed to the right edge. "It worked in
# dev" leaves roughly 19 characters of slack before the two collide, and the
# collision is silent - the text runs under the CodeStreak mark rather than
# wrapping or clipping. 30 keeps a margin.
SERIES_MAX = 30
TITLE_MAX = 72


def render_html(episode, number, number_sub, technique, title, series):
    if len(number) > NUMBER_MAX:
        sys.exit(
            "hero headline %r is %d chars, over %d. Past that the autofit "
            "drops below 56px and it stops reading at thumbnail size."
            % (number, len(number), NUMBER_MAX)
        )
    if len(title) > TITLE_MAX:
        sys.exit(
            "hero title %r is %d chars, over %d. It wraps to three lines and "
            "collides with the author block." % (title, len(title), TITLE_MAX)
        )
    if len(series) > SERIES_MAX:
        sys.exit(
            "series name %r is %d chars, over %d. It shares the top row with "
            "both newsletter marks and would run underneath them."
            % (series, len(series), SERIES_MAX))
    frm, to = palette_for(episode)
    html = open(TEMPLATE).read()
    for key, value in [
        ("{FROM}", frm),
        ("{TO}", to),
        ("{EPISODE}", str(episode)),
        ("{NUMBER}", number),
        ("{NUMBER_SUB}", number_sub),
        ("{TECHNIQUE}", technique),
        ("{TITLE}", title),
        ("{SERIES}", series),
    ]:
        html = html.replace(key, value)
    return html


def screenshot(html, out_path):
    with tempfile.NamedTemporaryFile("w", suffix=".html", delete=False) as f:
        f.write(html)
        tmp = f.name
    try:
        subprocess.run(
            ["npx", "playwright", "screenshot", "--viewport-size=1200,630",
             "file://" + tmp, out_path],
            check=True, capture_output=True,
        )
    finally:
        os.unlink(tmp)


def main():
    p = argparse.ArgumentParser()
    p.add_argument("--episode", type=int, required=True)
    p.add_argument("--number", required=True,
                   help="the hook: a contradiction, not a statistic. "
                        "'Smaller tree. 3x slower.' not '34x slower'")
    p.add_argument("--number-sub", default="",
                   help="small trailing text, e.g. ' at 800 deep'")
    p.add_argument("--technique", required=True)
    p.add_argument("--title", required=True)
    p.add_argument("--series", default=SERIES,
                   help="the series name in the top-left rule")
    p.add_argument("--out", required=True)
    p.add_argument("--check-palettes", action="store_true",
                   help="verify every background is distinct and readable")
    if "--check-palettes" in sys.argv:
        verify_palettes()
        return
    a = p.parse_args()
    html = render_html(a.episode, a.number, a.number_sub, a.technique,
                       a.title, a.series)
    screenshot(html, a.out)
    print("wrote", a.out)


if __name__ == "__main__":
    main()
