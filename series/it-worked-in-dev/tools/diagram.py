"""Walkthrough diagrams, in the same house style as the LeetCode series.

Consistency across the two properties is deliberate: a reader who has seen the
daily posts should recognise these on sight.

Two guards exist because both have been the cause of shipped-broken images:
a caption wider than the canvas is cut mid-word, and a label wider than its
box is clipped. Both raise rather than render.
"""
import os
import re
import subprocess
import xml.etree.ElementTree as ET

W, H = 1200, 640
DONE = "#3178c6"     # resolved / computed
CUR = "#e8890c"      # being worked on now
PEND = "#d3d3d3"     # not yet reached
REPEAT = "#c0392b"   # work being done again
STROKE = "#2c3e50"
EDGE = "#8a9a9a"
MUTED = "#9aa4ac"

CAPTION_MAX = 68


def _clen(t):
    return len(re.sub(r"&[a-z]+;", "X", t))


def check_caption(t):
    n = _clen(t)
    if n > CAPTION_MAX:
        raise SystemExit(f"CAPTION TOO LONG ({n} > {CAPTION_MAX}):\n  {t}")


def _hdr(step, total):
    return (
        f'<svg xmlns="http://www.w3.org/2000/svg" width="{W}" height="{H}" viewBox="0 0 {W} {H}">'
        f'<rect width="{W}" height="{H}" fill="white"/>'
        '<defs><marker id="a" viewBox="0 0 10 10" refX="9" refY="5" markerWidth="7" markerHeight="7" '
        f'orient="auto-start-reverse"><path d="M 0 0 L 10 5 L 0 10 z" fill="{EDGE}"/></marker>'
        '<marker id="up" viewBox="0 0 10 10" refX="9" refY="5" markerWidth="7" markerHeight="7" '
        f'orient="auto-start-reverse"><path d="M 0 0 L 10 5 L 0 10 z" fill="{DONE}"/></marker></defs>'
        f'<text x="40" y="52" font-family="monospace" font-size="30" fill="{MUTED}">STEP {step} / {total}</text>'
    )


def caption(t):
    return f'<text x="28" y="{H-34}" font-family="monospace" font-size="27" fill="{STROKE}">{t}</text>'


def label(x, y, t, size=26, color=STROKE, anchor="middle", bold=False):
    w = ' font-weight="bold"' if bold else ""
    return (f'<text x="{x}" y="{y}" font-family="monospace" font-size="{size}" '
            f'text-anchor="{anchor}" fill="{color}"{w}>{t}</text>')


BW, BH = 190, 62


def box(x, y, text, state="pend", sub=None, w=None):
    """A directory node. state: pend | cur | done | repeat."""
    w = w or BW
    if _clen(text) * 16 > w - 16:
        raise SystemExit(f"LABEL {text!r} does not fit a {w}px box")
    fill = {"done": DONE, "cur": "white", "pend": "#f4f6f8", "repeat": "#fdecea"}[state]
    sw = 6 if state in ("cur", "repeat") else 3
    sc = CUR if state == "cur" else (REPEAT if state == "repeat" else "#b9c4cc")
    tc = "white" if state == "done" else STROKE
    out = [f'<rect x="{x}" y="{y}" width="{w}" height="{BH}" rx="9" fill="{fill}" '
           f'stroke="{sc}" stroke-width="{sw}"/>',
           f'<text x="{x+w/2}" y="{y+BH/2+9}" font-family="monospace" font-size="26" '
           f'text-anchor="middle" fill="{tc}">{text}</text>']
    if sub:
        # To the right of the box, never below it: connectors leave from the
        # bottom edge and a label there lands on top of them.
        out.append(label(x + w + 14, y + BH / 2 + 8, sub, 24,
                         DONE if state == "done" else MUTED, anchor="start"))
    return out


def connect(x1, y1, x2, y2, color=EDGE, marker=None, width=4):
    m = f' marker-end="url(#{marker})"' if marker else ""
    return (f'<line x1="{x1}" y1="{y1}" x2="{x2}" y2="{y2}" stroke="{color}" '
            f'stroke-width="{width}"{m}/>')


def render(steps, outdir, tmpdir="/tmp"):
    os.makedirs(outdir, exist_ok=True)
    total = len(steps)
    made = []
    for i, (body, cap) in enumerate(steps, 1):
        check_caption(cap)
        svg = _hdr(i, total) + "\n" + "\n".join(body) + "\n" + caption(cap) + "\n</svg>\n"
        sp = f"{tmpdir}/cs-{i}.svg"
        open(sp, "w").write(svg)
        ET.parse(sp)
        op = os.path.join(outdir, f"walkthrough-{i}.png")
        subprocess.run(["rsvg-convert", "-w", "1200", "--keep-aspect-ratio", "-b", "white",
                        "-o", op, sp], check=True)
        made.append(op)
        os.remove(sp)

    from PIL import Image
    import numpy as np
    for op in made:
        a = np.array(Image.open(op).convert("L")) < 245
        if a[:, 0].sum() or a[:, -1].sum() or a[0, :].sum() or a[-1, :].sum():
            raise SystemExit(f"EDGE CONTACT: {op}")
    print(f"  {len(made)} diagrams -> {outdir}")
