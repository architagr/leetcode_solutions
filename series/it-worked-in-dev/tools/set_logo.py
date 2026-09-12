#!/usr/bin/env python3
"""Embed a brand mark into the hero template as a data URI.

The template carries its images inline rather than referencing files, because
the card is screenshotted from a temp file in /tmp and a relative <img src>
would resolve against that directory and silently render nothing.

Usage:
    python3 tools/set_logo.py golang_journal path/to/logo.png

Slots:
    golang_journal   the second newsletter mark, under CodeStreak Daily
"""
import base64
import mimetypes
import os
import sys

TOOLS = os.path.dirname(os.path.abspath(__file__))
TEMPLATE = os.path.join(TOOLS, "hero_template.html")

SLOTS = {
    "golang_journal": "{GOLANG_JOURNAL_SRC}",
}

# The mark renders into a 44px box at 2x for retina, so anything under 88px
# square is visibly soft on the card. Square, because the box crops to square
# and a wide logo loses its edges to object-fit: cover.
MIN_PX = 88


def main():
    if len(sys.argv) != 3 or sys.argv[1] not in SLOTS:
        sys.exit(__doc__)
    slot, path = sys.argv[1], sys.argv[2]

    if not os.path.exists(path):
        sys.exit("no such file: %s" % path)
    kind, _ = mimetypes.guess_type(path)
    if not kind or not kind.startswith("image/"):
        sys.exit("%s is not an image (guessed %r)" % (path, kind))

    try:
        from PIL import Image
        w, h = Image.open(path).size
        if min(w, h) < MIN_PX:
            sys.exit("%s is %dx%d; the mark renders at 44px on a 2x card, so "
                     "anything under %dpx square looks soft" % (path, w, h, MIN_PX))
        if abs(w - h) > max(w, h) * 0.1:
            print("warning: %s is %dx%d, not square. The slot crops to square "
                  "with object-fit: cover, so the long edges will be cut."
                  % (path, w, h))
    except ImportError:
        print("note: pillow not installed, skipping the size check")

    data = base64.b64encode(open(path, "rb").read()).decode()
    uri = "data:%s;base64,%s" % (kind, data)

    html = open(TEMPLATE).read()
    token = SLOTS[slot]
    if token not in html:
        sys.exit("slot %s is already filled. Reset it in %s first."
                 % (slot, TEMPLATE))
    open(TEMPLATE, "w").write(html.replace(token, uri))
    print("embedded %s into %s (%d KB)" % (path, slot, len(data) // 1024))


if __name__ == "__main__":
    main()
