import os, sys
sys.path.insert(0, "/Users/architagarwal/code/codestreak-growth/tools")
from diagram import *

Pn, C, D, R = "pend", "cur", "done", "repeat"


def stack(x, title, items, w=300):
    out = [label(x, 96, title, 22, MUTED, anchor="start")]
    for i, (t, st) in enumerate(items):
        out += box(x, 112 + i * 70, t, st, w=w)
    return out


def scene(parts, notes=()):
    out = []
    for p in parts:
        out += p
    for y, text, color in notes:
        out.append(label(600, y, text, 26, color))
    return out


steps = [

 (scene([stack(60, "the document", [("hello", D)], w=420),
         stack(560, "after each edit", [("hello world", D), ("hello world!", D), ("hi world!", C)], w=560)],
        notes=[(430, "type ' world', type '!', replace 'hello' with 'hi'", MUTED),
               (470, "undo walks back one edit at a time; redo walks forward", STROKE)]),
  "undo and redo: step back through edits, most recent first"),

 (scene([stack(60, "undo stack", [("hello", R), ("hello world", R), ("hello world!", R)], w=420),
         stack(560, "the document", [("hi world!", C)], w=420)],
        notes=[(430, "before every edit, keep a copy of the whole document", MUTED),
               (470, "undo: put the last copy back. It cannot be wrong.", STROKE)]),
  "what you would write: a snapshot of the whole document per edit"),

 (scene([stack(60, "undo stack, doc_50k", [("76.5 KB", R), ("76.5 KB", R), ("... 2,000 of them", R)], w=420),
         stack(560, "held", [("122 MB", R)], w=420)],
        notes=[(430, "2,000 edits of a few bytes each keep 2,000 whole copies", MUTED),
               (470, "and no edit can happen in place: history points at it", REPEAT)]),
  "the history holds the whole document once per keystroke"),

 (scene([stack(60, "what each edit changed", [("insert ' world' at 5", D), ("insert '!' at 11", D),
                                                   ("delete 'hello' at 0", D), ("insert 'hi' at 0", C)], w=460)],
        notes=[(430, "an edit changes a few bytes, and says which", MUTED),
               (470, "keep that, and its inverse is known: delete what went in", STROKE)]),
  "each edit already says exactly what it changed"),

 (scene([stack(60, "undo stack", [("+' world' @5", D), ("+'!' @11", D), ("-'hello' @0", D), ("+'hi' @0", C)], w=420),
         stack(560, "undo: pop, apply the inverse", [("delete 'hi' at 0", C), ("insert 'hello' at 0", D)], w=560)],
        notes=[(470, "78.7 KB held instead of 122 MB, and edits in place: 4.63x", STROKE)]),
  "a stack of edits: undo pops the most recent and applies its inverse"),
]
render(steps, os.path.join(os.path.dirname(os.path.abspath(__file__)), "images"))
