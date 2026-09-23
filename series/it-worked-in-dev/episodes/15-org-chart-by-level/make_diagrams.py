import os, sys
sys.path.insert(0, "/Users/architagarwal/code/codestreak-growth/tools")
from diagram import *

# A seven-person chart. Preorder visit numbers are marked because the whole
# episode is about when a level's last member shows up in that order.
BWS = 150
Y = [78, 178, 278, 378]
POS = {
    "CEO":  (525, 0),
    "Ops":  (300, 1), "Eng": (750, 1),
    "Pay":  (150, 2), "Web": (450, 2), "SRE": (750, 2),
    "API":  (150, 3),
}
KIDS = {"CEO": ["Ops", "Eng"], "Ops": ["Pay", "Web"], "Eng": ["SRE"], "Pay": ["API"]}
VISIT = {"CEO": 1, "Ops": 2, "Pay": 3, "API": 4, "Web": 5, "Eng": 6, "SRE": 7}

Pn, C, D, R = "pend", "cur", "done", "repeat"


def tree(states=None, subs=None, notes=()):
    states, subs = states or {}, subs or {}
    out = []
    for parent, kids in KIDS.items():
        px, pl = POS[parent]
        for kid in kids:
            kx, kl = POS[kid]
            out.append(connect(px + BWS / 2, Y[pl] + BH,
                               kx + BWS / 2, Y[kl], width=3))
    for name, (x, level) in POS.items():
        out += box(x, Y[level], name, states.get(name, Pn), sub=subs.get(name), w=BWS)
    for y, text, color in notes:
        out.append(label(600, y, text, 26, color))
    return out


def levels(fill, notes=()):
    """The page's own picture: one row per level, filled as it is delivered."""
    out = []
    rows = [("level 0", ["CEO"]), ("level 1", ["Ops", "Eng"]),
            ("level 2", ["Pay", "Web", "SRE"]), ("level 3", ["API"])]
    for i, (name, members) in enumerate(rows):
        y = 110 + i * 92
        out.append(label(230, y + 38, name, 26, MUTED, anchor="end"))
        for j, m in enumerate(members):
            out += box(260 + j * 170, y, m, D if i < fill else Pn, w=150)
    for y, text, color in notes:
        out.append(label(600, y, text, 26, color))
    return out


ALL = {n: D for n in POS}
SUBS = {n: "#%d" % v for n, v in VISIT.items()}

steps = [
 (levels(1, notes=[(500, "each level appears as soon as it is ready", MUTED)]),
  "the page paints the org chart one level at a time"),

 (tree(ALL, SUBS, notes=[(478, "preorder: down one branch, all the way, then back up", MUTED)]),
  "the recursive walk visits in this order, marked #1 to #7"),

 (tree({"Ops": D, "Eng": D}, SUBS,
       notes=[(478, "Ops is visit #2. Eng is visit #6, of 7.", STROKE),
              (520, "level 1 is not finished until almost the whole walk is", MUTED)]),
  "level 1 has two people, and they arrive six visits apart"),

 (tree({"Ops": D, "Eng": D, "SRE": R}, SUBS,
       notes=[(478, "after #6 the walk holds all of level 1, and cannot know it", REPEAT),
              (520, "#7 might have been another direct report", MUTED),
              (556, "55,987 people: level 1 complete at 46,657, provable at 55,987", STROKE)]),
  "having the answer and being able to say so are different things"),

 (tree({"CEO": C}, notes=[(478, "queue: [CEO].  len is 1, so level 0 is those 1.", STROKE),
                          (520, "emit it. Everything appended now is level 1.", MUTED)]),
  "the queue reads its own length, and that is the level boundary"),

 (tree({"CEO": D, "Ops": C, "Eng": C},
       notes=[(478, "queue: [Ops, Eng].  len is 2, so level 1 is those 2.", STROKE),
              (520, "three people looked at, and two levels are on screen", MUTED)]),
  "each round drains exactly one level and loads the next"),

 (tree({"CEO": D, "Ops": D, "Eng": D, "Pay": C, "Web": C, "SRE": C},
       notes=[(478, "the top three levels of a 55,987-person chart:", MUTED),
              (518, "43 people looked at, not 55,987", STROKE),
              (556, "producing the whole chart, though, is a dead heat", MUTED)]),
  "the first three levels cost the first three levels"),
]
render(steps, os.path.join(os.path.dirname(os.path.abspath(__file__)), "images"))
