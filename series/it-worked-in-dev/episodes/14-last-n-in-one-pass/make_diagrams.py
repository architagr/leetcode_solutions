import os, sys
sys.path.insert(0, "/Users/architagarwal/code/codestreak-growth/tools")
from diagram import *

# A log of seven lines, and the last four are wanted. Four slots is small
# enough to draw and big enough to wrap twice.
N = 4
BWS, GAP = 120, 20
PITCH = BWS + GAP
Y_TOP, Y_RING = 150, 300

Pn, C, D, R = "pend", "cur", "done", "repeat"


def stream(states, x0=60, y=Y_TOP, tail="..."):
    """The log as it arrives, left to right, with no known end."""
    out = []
    for i in range(7):
        out += box(x0 + i * PITCH, y, "L%d" % (i + 1), states.get(i, Pn), w=BWS)
    if tail:
        out.append(label(x0 + 7 * PITCH + BWS / 2, y + BH / 2 + 9, tail, 30, MUTED))
    return out


def ring(slots, w=None, states=None, y=Y_RING, x0=200):
    out = []
    states = states or {}
    for i, v in enumerate(slots):
        out += box(x0 + i * 220, y, v, states.get(i, D), w=190)
    if w is not None:
        cx = x0 + w * 220 + 95
        out.append(label(cx, y - 18, "w", 26, CUR, bold=True))
    return out


def note(y, text, color=MUTED, size=26):
    return [label(600, y, text, size, color)]


steps = [

 (stream({0: D, 1: D, 2: C}) +
  note(300, "lines arrive in order, and the end has not happened yet") +
  note(348, "\"the last four\" is defined against an end nobody has reached", MUTED),
  "a log is a stream: one pass, and no length up front"),

 (stream({i: D for i in range(7)}, tail="") +
  [connect(60 + 3 * PITCH, Y_TOP + BH + 16, 60 + 7 * PITCH - GAP, Y_TOP + BH + 16,
           color=DONE, width=5)] +
  note(290, "read it all, then take the last four") +
  note(338, "correct, one idea, and it reads like the ticket", MUTED),
  "what you would write: collect everything, slice the tail"),

 (stream({i: R for i in range(3)}, tail="") +
  stream({i: D for i in range(3, 7)}, x0=60, tail="")[3 * 2:] +
  note(290, "all[len-4:] is a window on the same array", REPEAT) +
  note(338, "so every line before it is still reachable", REPEAT) +
  note(404, "2,000,000 lines, n=200: 186 MB held to return 200", STROKE),
  "the returned tail keeps the whole log alive"),

 (stream({i: D for i in range(7)}, y=140, tail="") +
  [label(600, 236, "pass 1: count them. 7.", 26, MUTED)] +
  stream({i: D for i in range(3, 7)}, y=286, tail="")[6:] +
  [label(600, 382, "pass 2: keep from index 7 - 4", 26, MUTED)] +
  note(448, "a file can be opened twice. A socket cannot,", REPEAT) +
  note(492, "and a log still being written is a different log.", REPEAT),
  "counting first means reading it twice, which a stream forbids"),

 (ring(["L1", "L2", "L3", "L4"], w=0) +
  note(230, "four slots, and the fifth line has nowhere new to go") +
  note(432, "w is where the next line goes") +
  note(480, "which is also where the oldest line is sitting", STROKE),
  "instead: keep exactly four, and overwrite the oldest"),

 (ring(["L5", "L6", "L3", "L4"], w=2, states={0: D, 1: D, 2: R, 3: D}) +
  note(230, "L5 landed on L1, L6 landed on L2, and w moved twice") +
  note(432, "the slot w points at is always exactly four lines old", STROKE) +
  note(480, "that fixed gap is the length the function never computes", MUTED),
  "the write cursor and the oldest line never drift apart"),

 (ring(["L5", "L6", "L7", "L4"], w=3, states={3: C}) +
  [label(600, 230, "out = buf[3:] then buf[:3]  =  L4 L5 L6 L7", 26, STROKE)] +
  note(432, "one pass, no length, nothing held but four lines") +
  note(480, "2,000,000 lines, n=200: 22.3 KB instead of 186 MB", STROKE),
  "unwrap at w, and the last four come out in order"),
]
render(steps, os.path.join(os.path.dirname(os.path.abspath(__file__)), "images"))
