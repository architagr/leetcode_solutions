import os, sys
sys.path.insert(0, "/Users/architagarwal/code/codestreak-growth/tools")
from diagram import *

# products -> laptops -> gaming, and products -> phones
P = (505, 95)
L = (270, 185)
G = (270, 275)
PH = (760, 185)
EDGES = [(L, P), (G, L), (PH, P)]
BOXES = [(P, "products"), (L, "laptops"), (G, "gaming"), (PH, "phones")]

def menu(states, subs=None, note=None, out=None):
    subs = subs or {}
    res = []
    for (x, y), (px, py) in EDGES:
        res.append(connect(px + BW/2, py + BH, x + BW/2, y, width=3))
    for i, ((x, y), name) in enumerate(BOXES):
        res += box(x, y, name, states.get(i, "pend"), subs.get(i))
    if out:
        res.append(label(600, 395, out, 27, STROKE, bold=True))
    if note:
        res.append(label(600, 465, note, 26, MUTED))
    return res

Pn, C, D, R = "pend", "cur", "done", "repeat"

steps = [
 (menu({}, note="the renderer wants one flat row per item, in order"),
  "nested config in, flat list out"),

 (menu({2: C}, {2: "[gaming]"}, note="correct, finished, owned by this call"),
  "the deepest call builds its own slice and returns it"),

 (menu({1: C, 2: R}, {1: "[laptops, gaming]"},
       note="gaming's row is copied into a new slice"),
  "the parent merges it into a slice of its own"),

 (menu({0: C, 1: R, 2: R}, {0: "[products, laptops, gaming]"},
       note="and copied again. Once per level above it."),
  "and the grandparent merges that one too"),

 (menu({0: R, 1: R, 2: R, 3: R},
       note="180,300 row copies at 600 deep"),
  "every row is rewritten once per ancestor"),

 (menu({0: C}, out="[products]", note="one slice, and depth is a parameter"),
  "instead: one destination, written as you walk"),

 (menu({0: D, 1: D, 2: D, 3: D},
       out="[products, laptops, gaming, phones]",
       note="nothing copied out of one slice into another"),
  "same list, appended once each"),
]
render(steps, os.path.join(os.path.dirname(os.path.abspath(__file__)), "images"))
