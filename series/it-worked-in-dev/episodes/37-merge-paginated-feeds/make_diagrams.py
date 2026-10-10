import os, sys
sys.path.insert(0, "/Users/architagarwal/code/codestreak-growth/tools")
from diagram import *

# Four followed accounts, each feed newest first, in minutes ago. The first
# page of five: 1m (C), 2m (A), 5m (B), 6m (B), 9m (A).
FEEDS = {"A": [2, 9, 15, 40], "B": [5, 6, 30, 31], "C": [1, 22, 23, 50], "D": [12, 13, 14, 60]}
COLX = {"A": 90, "B": 280, "C": 470, "D": 660}
ROWY = [120, 190, 260, 330]
PW = 130
TX = 930
Pn, C, D, R = "pend", "cur", "done", "repeat"


def scene(states, page=(), notes=(), page_label="timeline, page 1"):
    out = []
    for f, x in COLX.items():
        out.append(label(x + PW / 2, 100, "feed " + f, 22, MUTED))
        for i, m in enumerate(FEEDS[f]):
            out += box(x, ROWY[i], "%dm" % m, states.get((f, i), Pn), w=PW)
    out.append(label(TX + PW / 2, 100, page_label, 22, MUTED))
    for i, (txt, st) in enumerate(page):
        out += box(TX, 120 + i * 62, txt, st, w=PW + 40) if False else box(TX, 112 + i * 64, txt, st, w=PW)
    for y, text, color in notes:
        out.append(label(600, y, text, 26, color))
    return out


FIRST5 = [("C", 0), ("A", 0), ("B", 0), ("B", 1), ("A", 1)]
steps = [

 (scene({}, notes=[
     (480, "each followed account's posts, newest first, from the posts service", MUTED),
     (520, "the timeline wants the newest 50 of all of them", STROKE)]),
  "the home timeline: one page from k feeds that are each sorted"),

 (scene({("A", i): R for i in range(4)} | {("B", i): R for i in range(4)},
        page=[("A+B", C)], notes=[
     (480, "merge A with B, then that with C, then with D: day 22, k times", MUTED),
     (520, "every merge copies everything merged so far", REPEAT)]),
  "what you would write: merge the feeds in, two at a time"),

 (scene({(f, i): R for f in "AB" for i in range(4)} | {("C", i): D for i in range(4)},
        page=[("A+B+C", C)], notes=[
     (480, "feed A is copied once per feed that comes after it", MUTED),
     (520, "5,000 feeds: 1.25 billion copies, 20 GB, 3.03 s for 50 posts", REPEAT)]),
  "folding copies the early feeds once for every later one"),

 (scene({("A", 0): C, ("B", 0): C, ("C", 0): C, ("D", 0): C}, notes=[
     (480, "the newest post of all is the newest of the four fronts", MUTED),
     (520, "a merge only ever compares fronts. The rest is in order already.", STROKE)]),
  "the newest post overall is always one of the fronts"),

 (scene({("C", 0): D, ("A", 0): C, ("B", 0): C, ("C", 1): C, ("D", 0): C},
        page=[("1m C", D)], notes=[
     (480, "keep the k fronts in a heap, newest at the root", MUTED),
     (520, "take the root; only that feed's front changes: C moves to 22m", STROKE)]),
  "a heap of fronts: take the newest, one feed moves along"),

 (scene({("C", 0): D, ("A", 0): D, ("B", 0): D, ("B", 1): D, ("A", 1): D,
         ("A", 2): C, ("B", 2): C, ("C", 1): C, ("D", 0): C},
        page=[("1m C", D), ("2m A", D), ("5m B", D), ("6m B", D), ("9m A", D)], notes=[
     (480, "stop when the page is full: the rest of every feed is never read", MUTED),
     (520, "5,000 feeds: 130 µs, 82.9 KB. 23,272x the fold.", STROKE)]),
  "stop at a full page: most of every feed is never touched"),
]
render(steps, os.path.join(os.path.dirname(os.path.abspath(__file__)), "images"))
