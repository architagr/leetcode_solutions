import os, sys
sys.path.insert(0, "/Users/architagarwal/code/codestreak-growth/tools")
from diagram import *

# Two feeds, each already sorted, whose timestamps interleave - and a tie at 4,
# because the tie is half the argument.
OWN = ["1", "4", "6", "9"]
FOL = ["2", "4", "7", "8"]

BWS, GAP = 120, 20          # a narrower box: eight of them have to fit
PITCH = BWS + GAP
LEFT = 200                  # room for the row label
Y_OWN, Y_FOL, Y_OUT = 120, 236, 410

Pn, C, D, R = "pend", "cur", "done", "repeat"


def row(y, vals, states, x0=LEFT, name=None):
    out = []
    if name:
        out.append(label(x0 - 26, y + BH / 2 + 9, name, 26, MUTED, anchor="end"))
    for i, v in enumerate(vals):
        out += box(x0 + i * PITCH, y, v, states.get(i, Pn), w=BWS)
    return out


def cursor(y, i, text, x0=LEFT):
    """A caret over one slot, naming which cursor is parked there."""
    cx = x0 + i * PITCH + BWS / 2
    return [label(cx, y - 16, text, 24, CUR, bold=True)]


def note(y, text, color=MUTED, size=26):
    return [label(600, y, text, size, color)]


# ---------------------------------------------------------------- the inputs
inputs_plain = (row(Y_OWN, OWN, {i: D for i in range(4)}, name="own") +
                row(Y_FOL, FOL, {i: D for i in range(4)}, name="followed") +
                note(360, "both queries ended in ORDER BY at", MUTED))

# ------------------------------------------------- concatenate, then sort it
CAT = OWN + FOL
cat_states = {i: D for i in range(4)}
cat_states.update({i: R for i in range(4, 8)})

concat = (row(Y_OWN, CAT, cat_states, x0=60, name=None) +
          note(Y_OWN - 30, "append(own, followed...) - two runs, one array") +
          note(330, "9 before 2: the array is no longer sorted at all", REPEAT))

sorted_row = (row(Y_OWN, ["1", "2", "4", "4", "6", "7", "8", "9"],
                  {i: D for i in range(8)}, x0=60) +
              note(Y_OWN - 30, "sort.Slice puts it back in order") +
              note(330, "200,050 events: 10,455,458 comparisons", REPEAT) +
              note(378, "to rebuild an order the database had already", MUTED))

# ------------------------------------------------------------------ the merge
def merge_step(i, j, out_vals, mark_tie=False, notes=()):
    st_own = {k: D for k in range(i)}
    st_fol = {k: D for k in range(j)}
    if i < len(OWN):
        st_own[i] = C
    if j < len(FOL):
        st_fol[j] = C
    body = (row(Y_OWN, OWN, st_own, name="own") +
            row(Y_FOL, FOL, st_fol, name="followed"))
    if i < len(OWN):
        body += cursor(Y_OWN, i, "i")
    if j < len(FOL):
        body += cursor(Y_FOL, j, "j")
    body += row(Y_OUT, out_vals + ["?"] * (8 - len(out_vals)),
                {k: D for k in range(len(out_vals))}, x0=60, name=None)
    body += [label(60, Y_OUT - 16, "out", 24, MUTED, anchor="start")]
    for y, t, c in notes:
        body += note(y, t, c)
    return body


step_first = merge_step(0, 0, ["1"], notes=[
    (360, "1 &lt;= 2, so own goes first. One comparison, one event.", MUTED)])

step_tie = merge_step(1, 1, ["1", "2"], notes=[
    (360, "4 &lt;= 4 is true, so own wins the tie - on every run", MUTED),
    (556, "&lt;= rather than &lt; is the whole tie-break policy", STROKE)])

step_tail = (row(Y_OWN, OWN, {0: D, 1: D, 2: D, 3: C}, name="own") +
             cursor(Y_OWN, 3, "i") +
             row(Y_FOL, FOL, {k: D for k in range(4)}, name="followed") +
             [label(LEFT + 4 * PITCH - 6, Y_FOL + BH / 2 + 9, "empty", 24, MUTED,
                    anchor="start")] +
             row(Y_OUT, ["1", "2", "4", "4", "6", "7", "8", "9"],
                 {k: D for k in range(7)}, x0=60) +
             [label(60, Y_OUT - 16, "out", 24, MUTED, anchor="start")] +
             note(360, "what is left in own is already in order, and later") +
             note(556, "so it is copied whole, not compared event by event", STROKE))

# --------------------------------------------------------------- the lopsided
lopsided = (
    [label(174, Y_OWN + BH / 2 + 9, "own", 26, MUTED, anchor="end")] +
    box(200, Y_OWN, "200,000 events", D, w=820) +
    [label(174, Y_FOL + BH / 2 + 9, "followed", 26, MUTED, anchor="end")] +
    box(200, Y_FOL, "50", D, w=120) +
    [label(340, Y_FOL + BH / 2 + 8, "all of them older than everything in own",
           24, MUTED, anchor="start")] +
    note(360, "the 50 are placed first, and then followed is empty") +
    note(410, "105 comparisons, then one copy of 199,995 events", STROKE) +
    note(462, "sort.Slice on the same input: 10,455,458 comparisons", REPEAT))

steps = [
    (inputs_plain, "two queries, and the database ordered each one already"),
    (concat, "the obvious move: concatenate the two result sets"),
    (sorted_row, "then sort, which is where the ordering is paid for twice"),
    (step_first, "the merge instead: take the smaller of the two fronts"),
    (step_tie, "a tie goes to own, by policy rather than by luck"),
    (step_tail, "one side runs out, and the rest is attached in one move"),
    (lopsided, "the quiet source is where the two versions part company"),
]
render(steps, os.path.join(os.path.dirname(os.path.abspath(__file__)), "images"))
