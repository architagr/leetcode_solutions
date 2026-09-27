import os, sys
sys.path.insert(0, "/Users/architagarwal/code/codestreak-growth/tools")
from diagram import *

Pn, C, D, R = "pend", "cur", "done", "repeat"


def row(y, chars, states, x0=60, step=110, w=90):
    out = []
    for i, (c, st) in enumerate(zip(chars, states)):
        out += box(x0 + i * step, y, c, st, w=w)
    return out


def scene(parts, notes=()):
    out = []
    for p in parts:
        out += p
    for y, text, color in notes:
        out.append(label(600, y, text, 26, color))
    return out


steps = [

 (scene([[label(60, 110, "discount(max(cart[total], 500), {tier: gold})", 26, STROKE, anchor="start")],
         [label(60, 170, "the linter reads only the brackets:", 22, MUTED, anchor="start")],
         row(200, list("(([]){})"), [D] * 8)],
        notes=[(430, "every opener closed, by its own kind, in reverse order", MUTED),
               (470, "a file that does not nest is rejected before parsing", STROKE)]),
  "a rules file is valid only if its brackets nest properly"),

 (scene([row(120, list("([)]"), [D, D, R, R]),
         [label(60, 250, "round: 1 1 0 0    square: 0 1 1 0", 26, STROKE, anchor="start")]],
        notes=[(430, "count each kind: every count ends at zero, none dips below", MUTED),
               (470, "so \"([)]\" passes. The ] closes a ( that is still inside it.", REPEAT)]),
  "what ships first: count each kind, and ([)] passes"),

 (scene([[label(60, 100, "pass 1", 22, MUTED, anchor="start")], row(116, list("([{}])"), [Pn, Pn, R, R, Pn, Pn]),
         [label(60, 214, "pass 2", 22, MUTED, anchor="start")], row(230, list("([])"), [Pn, R, R, Pn]),
         [label(60, 328, "pass 3", 22, MUTED, anchor="start")], row(344, list("()"), [R, R])],
        notes=[(470, "correct: delete (), [] and {} until nothing changes", MUTED),
               (510, "one layer per pass. 5,000 deep: 21.7 ms and 25.4 MB.", REPEAT)]),
  "what you would write next: erase adjacent pairs until none are left"),

 (scene([row(150, list("([)]"), [D, D, C, Pn]),
         [label(60, 290, "open and not yet closed:  (  [", 26, STROKE, anchor="start")]],
        notes=[(430, "when ) arrives, only the most recent opener can match it", MUTED),
               (470, "and that is [, which has not closed yet", STROKE)]),
  "a closer can only match the bracket opened most recently"),

 (scene([row(150, list("([)]"), [D, D, R, Pn]),
         [label(60, 280, "stack:", 22, MUTED, anchor="start")],
         row(300, list("(["), [D, C], x0=160)],
        notes=[(470, "push openers. ) needs ( on top, and the top is [: reject.", STROKE),
               (510, "one pass, stops at the first mismatch. 38.2 µs at 5,000 deep.", STROKE)]),
  "a stack of what is still open: one pass, and it stops at the error"),
]
render(steps, os.path.join(os.path.dirname(os.path.abspath(__file__)), "images"))
