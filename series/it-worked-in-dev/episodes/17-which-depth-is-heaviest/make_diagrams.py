import os, sys
sys.path.insert(0, "/Users/architagarwal/code/codestreak-growth/tools")
from diagram import *

# The tied catalogue from catalog_test.go: depth 1 and depth 2 both hold 12
# SKUs. Every diagram uses it, because a tie is the only place the map version
# and the ordered versions disagree.
BWS = 170
Y = [90, 200, 310]
POS = {"root 0": (290, 0),
       "home 6": (110, 1), "garden 6": (470, 1),
       "kitchen 3": (20, 2), "bath 3": (200, 2),
       "tools 4": (380, 2), "seeds 2": (560, 2)}
KIDS = {"root 0": ["home 6", "garden 6"],
        "home 6": ["kitchen 3", "bath 3"],
        "garden 6": ["tools 4", "seeds 2"]}
DEPTH = {n: lvl for n, (_, lvl) in POS.items()}

Pn, C, D, R = "pend", "cur", "done", "repeat"
SLOT_X = 850


def tree(states=None, slots=None, slot_title=None, depth_labels=True, notes=()):
    states = states or {}
    out = []
    for parent, kids in KIDS.items():
        px, pl = POS[parent]
        for kid in kids:
            kx, kl = POS[kid]
            out.append(connect(px + BWS / 2, Y[pl] + BH,
                               kx + BWS / 2, Y[kl], width=3))
    for name, (x, level) in POS.items():
        out += box(x, Y[level], name, states.get(name, Pn), w=BWS)
    if slots is not None:
        if slot_title:
            out.append(label(SLOT_X, 66, slot_title, 22, MUTED, anchor="start"))
        for i, (text, state) in enumerate(slots):
            out += box(SLOT_X, Y[i], text, state, w=250)
            if depth_labels:
                out.append(label(SLOT_X - 14, Y[i] + BH / 2 + 8, "d%d" % i, 24,
                                 MUTED, anchor="end"))
    for y, text, color in notes:
        out.append(label(600, y, text, 26, color))
    return out


ALL = {n: D for n in POS}


def upto(d, cur=None):
    return {n: (C if DEPTH[n] == cur else D) for n in POS
            if DEPTH[n] <= d or DEPTH[n] == cur}


steps = [

 (tree(ALL, slots=[("0", D), ("12", D), ("12", D)], slot_title="SKUs per depth",
       notes=[(470, "depth 1 and depth 2 both hold 12 SKUs", MUTED),
              (510, "the report wants the shallowest: depth 1", STROKE)]),
  "two depths tie at 12 SKUs, and the shallowest one is the answer"),

 (tree(ALL, slots=[("2: 12", Pn), ("0: 0", Pn), ("1: 12", Pn)],
       slot_title="sums, as ranged", depth_labels=False,
       notes=[(470, "add each category into sums[depth], take the biggest", MUTED),
              (510, "ranging over a map, Go picks the order, not you", STROKE)]),
  "what you would write: sum into a map by depth, take the biggest"),

 (tree(ALL, slots=[("2: 12", R), ("0: 0", Pn), ("1: 12", Pn)],
       slot_title="this call's order", depth_labels=False,
       notes=[(470, "strict &gt; keeps the first 12 the loop meets", MUTED),
              (510, "1,000 calls: 883 said depth 1, 117 said depth 2", REPEAT)]),
  "same catalogue, same code: whichever 12 comes first wins"),

 (tree(upto(1, cur=2), slots=[("reached 1st", D), ("reached 2nd", D), ("reached 3rd", C)],
       slot_title="as the walk went",
       notes=[(470, "the walk reaches depth 0 before 1, and 1 before 2", MUTED),
              (510, "the order was produced. The map threw it away.", STROKE)]),
  "the walk met the depths in order, and the map dropped it"),

 (tree(ALL, slots=[("0", Pn), ("12  best", D), ("12  not &gt;", R)],
       slot_title="sums[d], read in order",
       notes=[(470, "depths are 0, 1, 2 with no gaps: an index, not a key", MUTED),
              (510, "read left to right, strict &gt; keeps depth 1. 2.33x.", STROKE)]),
  "depth is an index: a slice, read in order, keeps the shallowest"),

 (tree(upto(1, cur=2), slots=[("0", Pn), ("12  best", D), ("12  not &gt;", C)],
       slot_title="compared as it closes",
       notes=[(470, "close a level, compare it, forget it", MUTED),
              (510, "no recursion: 1.94x on a 2,000-deep chain", STROKE)]),
  "day 33's walk: compare each level the moment it closes"),
]
render(steps, os.path.join(os.path.dirname(os.path.abspath(__file__)), "images"))
