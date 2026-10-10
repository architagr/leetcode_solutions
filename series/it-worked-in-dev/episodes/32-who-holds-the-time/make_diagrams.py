import os, sys
sys.path.insert(0, "/Users/architagarwal/code/codestreak-growth/tools")
from diagram import *

# main calls parse, parse calls lex twice, then main calls render.
# Self time: main 4, parse 3, lex 5, render 3.
SPANS = [("main", 0, 15, 0), ("parse", 2, 10, 1), ("lex", 3, 5, 2), ("lex", 6, 9, 2),
         ("render", 11, 14, 1)]
X0, U = 105, 66          # x of t=0, px per unit
ROW = [110, 190, 270]    # y per depth
BARH = 60
FILL = {"pend": "#eef1f4", "cur": CUR, "done": DONE, "repeat": "#f2c4bd"}
TXT = {"pend": STROKE, "cur": "white", "done": "white", "repeat": STROKE}


def x(t):
    return X0 + t * U


def chart(states, notes=(), extra=(), ticks=False, owners=None):
    out = []
    if ticks:  # under the bars, so they never cross a name
        for t in (0, 2, 3, 5, 6, 9, 10, 11, 14, 15):
            out.append(connect(x(t), 95, x(t), 345, color=CUR, width=2))
    for i, (name, s, e, d) in enumerate(SPANS):
        st = states.get(i, "pend")
        out.append(f'<rect x="{x(s)}" y="{ROW[d]}" width="{(e - s) * U}" height="{BARH}" rx="6" '
                   f'fill="{FILL[st]}" stroke="#b9c4cc" stroke-width="2"/>')
        out.append(label(x(s) + 12, ROW[d] + 39, name, 24, TXT[st], anchor="start"))
    for t in range(0, 16):
        out.append(label(x(t), 372, str(t), 20, MUTED))
    if owners:
        for (a, b, who) in owners:
            out.append(f'<rect x="{x(a)}" y="392" width="{(b - a) * U}" height="34" '
                       f'fill="{DONE}" stroke="white" stroke-width="2"/>')
            if b - a >= 1:
                out.append(label((x(a) + x(b)) / 2, 416, who[0], 22, "white"))
    out += list(extra)
    for y, text, color in notes:
        out.append(label(600, y, text, 26, color))
    return out


ALL = {i: "done" for i in range(len(SPANS))}
OWN = [(0, 2, "main"), (2, 3, "parse"), (3, 5, "lex"), (5, 6, "parse"), (6, 9, "lex"),
       (9, 10, "parse"), (10, 11, "main"), (11, 14, "render"), (14, 15, "main")]

steps = [

 (chart(ALL, notes=[
     (470, "one request, as the trace logs it: a start and an end per call", MUTED),
     (510, "self time: main 4, parse 3, lex 5, render 3", STROKE)]),
  "self time: how long each function ran, minus what it called"),

 (chart({0: "cur", 1: "done", 4: "done"}, notes=[
     (470, "main ran 0 to 15. Take off parse (8) and render (3).", MUTED),
     (510, "15 - 8 - 3 = 4. Only the calls one level down.", STROKE)]),
  "what you would write: duration minus the direct callees"),

 (chart({0: "cur", 1: "repeat", 2: "repeat", 3: "repeat", 4: "repeat"}, notes=[
     (470, "subtract everything inside main instead: 15 - 8 - 2 - 3 - 3", MUTED),
     (510, "= -1. Lex was already inside parse's 8. Counted twice.", REPEAT)]),
  "subtracting every call inside double counts the grandchildren"),

 (chart({0: "cur", 1: "repeat", 2: "repeat", 3: "repeat", 4: "repeat"}, notes=[
     (470, "to find its direct callees, every call reads every call inside it", MUTED),
     (510, "parse rereads both lex calls. main rereads all four.", REPEAT),
     (550, "5,000 deep: 12.5 million reads, 8.20 ms", REPEAT)]),
  "a call reads everything nested inside it: deep traces pay depth"),

 (chart(ALL, ticks=True, notes=[
     (470, "between two neighbouring events, nothing starts or ends", MUTED),
     (510, "so exactly one function is running: the innermost open call", STROKE)]),
  "cut the timeline at every event: each slice has one owner"),

 (chart(ALL, ticks=True, owners=OWN, notes=[
     (470, "replay the events with a stack; the top owns each gap", MUTED),
     (510, "m p l p l p m r m: one pass, no spans, 307x on deep recursion", STROKE)]),
  "the top of the stack owns the time since the last event"),
]
render(steps, os.path.join(os.path.dirname(os.path.abspath(__file__)), "images"))
