import os, sys
sys.path.insert(0, "/Users/architagarwal/code/codestreak-growth/tools")
from diagram import *

# Two teams of three services. The second team is laid out right to left so
# the cross-team link, when it appears, is a short vertical line from C to D.
W_ = 150
POS = {"A": (40, 110), "B": (260, 110), "C": (480, 110),
       "F": (40, 290), "E": (260, 290), "D": (480, 290)}
LINKS = [("A", "B"), ("B", "C"), ("D", "E"), ("E", "F")]
CROSS = ("C", "D")
Pn, C, D, R = "pend", "cur", "done", "repeat"
SX = 760


def line(a, b):
    ax, ay = POS[a]
    bx, by = POS[b]
    if ay == by:
        l, r = sorted([(ax, a), (bx, b)])
        return connect(l[0] + W_, ay + BH / 2, r[0], ay + BH / 2, width=4)
    return connect(ax + W_ / 2, ay + BH, bx + W_ / 2, by, width=4)


def mesh(states=None, cross=False, rows=(), title=None, notes=()):
    states = states or {}
    out = [line(a, b) for a, b in LINKS]
    if cross:
        out.append(line(*CROSS))
    for n, (x, y) in POS.items():
        out += box(x, y, n, states.get(n, Pn), w=W_)
    if title:
        out.append(label(SX, 96, title, 22, MUTED, anchor="start"))
    for i, (text, state) in enumerate(rows):
        out += box(SX, 116 + i * 74, text, state, w=400)
    for y, text, color in notes:
        out.append(label(600, y, text, 26, color))
    return out


ALL = {n: D for n in POS}
T1 = {"A": D, "B": D, "C": D}

steps = [

 (mesh(ALL, title="networks",
       rows=[("{A, B, C}", D), ("{D, E, F}", D)],
       notes=[(460, "links are two-way: A reaches C, so C reaches A", MUTED),
              (500, "two teams, nothing between them: 2 networks", STROKE)]),
  "two teams, six services, two separate networks"),

 (mesh({"A": C, "B": D, "C": D}, title="Reachable(s), per service",
       rows=[("A: {A, B, C}", D)],
       notes=[(460, "the blast-radius function, already written and tested", MUTED),
              (500, "name each network by its smallest service: A", STROKE)]),
  "what you would write: ask every service what it can reach"),

 (mesh({"A": R, "B": R, "C": R}, title="Reachable(s), per service",
       rows=[("A: {A, B, C}", D), ("B: {A, B, C}", R), ("C: {A, B, C}", R)],
       notes=[(460, "B and C search the same three services again", REPEAT),
              (500, "3 services, 3 searches of 3: 9 visits for 3", STROKE)]),
  "every member of a network repeats the same search"),

 (mesh({n: R for n in POS}, cross=True, title="one link later",
       rows=[("6 searches of 6", R), ("36 visits", R), ("5,000: 25,000,000", R)],
       notes=[(460, "one cross-team link: now one network of six", MUTED),
              (500, "499 links like it: 2.82 ms became 1.71 s", REPEAT)]),
  "join the teams and every search covers the whole mesh"),

 (mesh(T1, cross=True, title="after the first search",
       rows=[("A: {A, B, C, D, E, F}", D), ("B: already known", Pn),
             ("C: already known", Pn)],
       notes=[(460, "the first search returned the whole network", MUTED),
              (500, "B's answer was already in hand. So was C's.", STROKE)]),
  "the first search already answered every member of it"),

 (mesh({"A": D, "B": D, "C": D, "D": C}, title="one shared seen set",
       rows=[("A: new, count 1", D), ("B, C: seen, skip", Pn),
             ("D: new, count 2", C)],
       notes=[(460, "one visited set for the whole scan, not one per search", MUTED),
              (500, "every service visited once: 5,000, not 25,000,000", STROKE)]),
  "one sweep: a service nobody has reached yet starts a network"),
]
render(steps, os.path.join(os.path.dirname(os.path.abspath(__file__)), "images"))
