import os, sys
sys.path.insert(0, "/Users/architagarwal/code/codestreak-growth/tools")
from diagram import *

# The lopsided thread: the newest reply to the root is a dead end, and the
# deep replies hang off an older sibling. It is the shape the tempting version
# gets wrong, so every diagram uses it.
BWS = 170
Y = [78, 178, 278, 378]
POS = {"root": (330, 0), "old": (150, 1), "new": (560, 1),
       "old-1": (150, 2), "old-2": (150, 3)}
KIDS = {"root": ["old", "new"], "old": ["old-1"], "old-1": ["old-2"]}

Pn, C, D, R = "pend", "cur", "done", "repeat"
SLOT_X = 830


def tree(states=None, subs=None, slots=None, notes=(), slot_title=None):
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
    if slots is not None:
        if slot_title:
            out.append(label(SLOT_X, 58, slot_title, 22, MUTED, anchor="start"))
        for i, (text, state) in enumerate(slots):
            out += box(SLOT_X, Y[i], text, state, w=250)
            out.append(label(SLOT_X - 14, Y[i] + BH / 2 + 8, "d%d" % i, 24,
                             MUTED, anchor="end"))
    for y, text, color in notes:
        out.append(label(600, y, text, 26, color))
    return out


ALL = {n: D for n in POS}
VISIBLE = {"root": D, "new": D, "old-1": D, "old-2": D}
EMPTY = [("-", Pn)] * 4
FULL = [("root", D), ("new", D), ("old-1", D), ("old-2", D)]

steps = [

 (tree(VISIBLE, slots=FULL, slot_title="collapsed view",
       notes=[(490, "one line per nesting level: the newest reply there", MUTED),
              (530, "\"old\" is hidden. \"new\", beside it, is not.", STROKE)]),
  "collapsed, a thread shows one line per nesting level"),

 (tree(ALL, slots=FULL, slot_title="the four lines",
       notes=[(490, "build every level, then take the last of each", MUTED),
              (530, "55,987 comments held, to return 7 lines", REPEAT)]),
  "what you would write: every level, then the last of each"),

 (tree({"root": R, "new": R}, slots=[("root", D), ("new", D), ("-", Pn), ("-", Pn)],
       slot_title="what it finds",
       notes=[(490, "follow the newest reply down, and it stops at \"new\"", REPEAT),
              (530, "the deeper levels are in the OLDER branch beside it", STROKE)]),
  "the tempting version: follow the newest reply all the way down"),

 (tree(ALL, slots=FULL, slot_title="one slot per depth",
       notes=[(490, "every comment overwrites the slot for its own depth", MUTED),
              (530, "240 bytes instead of 2.87 MB, and 4.78x faster", STROKE)]),
  "keep one slot per depth and let the last writer win"),

 (tree(ALL, subs={"old": "x", "old-1": "x", "old-2": "x"},
       slots=FULL, slot_title="7 kept",
       notes=[(490, "but the overwritten value was still built first", REPEAT),
              (530, "55,987 previews built. Seven survive.", STROKE)]),
  "the loser is produced and then thrown away, every time"),

 (tree(ALL, subs={"root": "#1", "new": "#2", "old": "#3",
                  "old-1": "#4", "old-2": "#5"},
       slots=EMPTY, slot_title="still empty",
       notes=[(490, "walk the replies newest first, not oldest first", MUTED),
              (530, "so \"new\" is reached before \"old\", at depth 1", STROKE)]),
  "reverse the walk: the newest reply is visited first"),

 (tree(VISIBLE, subs={"root": "#1", "new": "#2", "old-1": "#4", "old-2": "#5"},
       slots=FULL, slot_title="written once each",
       notes=[(490, "record only when the depth has no entry yet", MUTED),
              (530, "7 previews built instead of 55,987. 22.5x, 1.12 KB.", STROKE)]),
  "first arrival at a depth wins, so nothing is built twice"),
]
render(steps, os.path.join(os.path.dirname(os.path.abspath(__file__)), "images"))
