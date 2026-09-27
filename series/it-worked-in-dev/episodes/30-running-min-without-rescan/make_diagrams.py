import os, sys
sys.path.insert(0, "/Users/architagarwal/code/codestreak-growth/tools")
from diagram import *

Pn, C, D, R = "pend", "cur", "done", "repeat"
# opening 1000; transactions -200, -900, +500; balances 800, -100, 400
COLS = [(60, "amount"), (330, "balance"), (600, "lowest so far")]


def table(rows, title=None, cols=COLS):
    out = []
    if title:
        out.append(label(60, 100, title, 22, MUTED, anchor="start"))
    for x, h in cols:
        out.append(label(x, 140, h, 22, MUTED, anchor="start"))
    for i, cells in enumerate(rows):
        for (x, _), (text, st) in zip(cols, cells):
            out += box(x, 156 + i * 72, text, st, w=240)
    return out


def scene(parts, notes=()):
    out = []
    for p in parts:
        out += p
    for y, text, color in notes:
        out.append(label(600, y, text, 26, color))
    return out


steps = [

 (scene([table([[("open 1,000", D), ("1,000", D), ("1,000", D)],
                 [("-200", D), ("800", D), ("800", D)],
                 [("-900", C), ("-100", C), ("-100", R)]], title="the plan")],
        notes=[(470, "the warning: the lowest balance the plan ever reaches", MUTED),
               (510, "here -100: the plan goes overdrawn after the rent", STROKE)]),
  "the overdraft warning: the lowest balance the plan reaches"),

 (scene([table([[("-200", D), ("800", R), ("800", R)],
                 [("-900", D), ("-100", R), ("-100", R)],
                 [("+500", C), ("400", R), ("-100", R)]], title="after every change: walk the whole plan")],
        notes=[(470, "add or undo, then rescan from the opening balance", MUTED),
               (510, "a replayed log of 106,103 changes: 3.04 s", REPEAT)]),
  "what you would write: rescan the plan on every change"),

 (scene([table([[("-200", D), ("800", D), ("low: -100", R)],
                 [("-900", R), ("undone", R), ("still -100", R)]],
               title="one tracked minimum, then undo the -900")],
        notes=[(470, "the balance can be subtracted back; the minimum cannot", MUTED),
               (510, "it says -100. The plan's lowest is now 800.", REPEAT)]),
  "the obvious shortcut: one minimum field, wrong after an undo"),

 (scene([table([[("-200", D), ("800", D), ("800", D)],
                 [("-900", D), ("-100", D), ("-100", D)],
                 [("+500", C), ("400", C), ("-100", C)]], title="each row, as it stood when it was added")],
        notes=[(470, "the answer after each transaction was known when it was added", MUTED),
               (510, "an undo only ever goes back to an answer already worked out", STROKE)]),
  "every row's answer was known the moment the row was added"),

 (scene([table([[("-200", D), ("800", D), ("800", C)],
                 [("-900", R), ("-100", R), ("-100", R)]],
               title="undo the -900: drop the top row")],
        notes=[(470, "the lowest is the top row's: 800. Nothing is recomputed.", STROKE),
               (510, "13 ns a change instead of 28.6 µs. 2224x.", STROKE)]),
  "carry the answer per row: undo pops, and the answer is on top"),
]
render(steps, os.path.join(os.path.dirname(os.path.abspath(__file__)), "images"))
