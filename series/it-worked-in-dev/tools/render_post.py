#!/usr/bin/env python3
"""Resolve a draft's {{placeholders}} and strip its editor notes.

Drafts keep URLs as {{github_growth}} so a moved newsletter is one edit in
LINKS.yaml rather than a hunt through every episode. This produces the copy
that actually gets pasted into LinkedIn or Medium: placeholders replaced, the
HTML comment at the top removed, frontmatter removed.

    python3 tools/render_post.py <draft.md>              # to stdout
    python3 tools/render_post.py <episode-folder>        # every draft, to out/

Frontmatter is not thrown away silently - meta_title, meta_description and the
hashtags are printed to stderr, because they have to be pasted into fields the
body cannot carry.
"""
import os
import re
import sys

try:
    import yaml
except ImportError:
    sys.exit("needs pyyaml:  pip3 install pyyaml")

SERIES = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))


def links():
    data = yaml.safe_load(open(os.path.join(SERIES, "LINKS.yaml")))
    out = dict(data.get("hub") or {})
    out.update(data.get("spokes") or {})
    return out


def render(path, url_map):
    text = open(path).read()

    unresolved = []

    def sub(m):
        key = m.group(1)
        value = url_map.get(key)
        if not value:
            unresolved.append(key)
            return m.group(0)
        return value

    # Substitute before splitting off the frontmatter: `canonical` carries a
    # {{github_growth}} and was coming out with the braces still in it.
    text = re.sub(r"\{\{(\w+)\}\}", sub, text)

    front = ""
    m = re.search(r"^---\n(.*?)\n---\n", text, re.S | re.M)
    if m:
        front = m.group(1)
        text = text[:m.start()] + text[m.end():]

    text = re.sub(r"^<!--.*?-->\s*", "", text, flags=re.S)
    if unresolved:
        sys.exit("%s: unresolved %s - set it in LINKS.yaml"
                 % (path, ", ".join(sorted(set(unresolved)))))

    return text.strip() + "\n", front


def report(path, front):
    name = os.path.basename(path)
    print("\n--- %s: paste these into the platform's own fields ---" % name,
          file=sys.stderr)
    for key in ("meta_title", "meta_description", "tags", "hashtags", "canonical"):
        m = re.search(r"^%s:\s*(.*)$" % key, front, re.M)
        if m:
            print("  %-16s %s" % (key, m.group(1).strip()), file=sys.stderr)


def main():
    if len(sys.argv) != 2:
        sys.exit(__doc__)
    target = sys.argv[1]
    url_map = links()

    if os.path.isdir(target):
        posts = os.path.join(target, "posts")
        src = posts if os.path.isdir(posts) else target
        out_dir = os.path.join(target, "out")
        os.makedirs(out_dir, exist_ok=True)
        for f in sorted(os.listdir(src)):
            if not f.startswith("POST_") or not f.endswith(".md"):
                continue
            body, front = render(os.path.join(src, f), url_map)
            dest = os.path.join(out_dir, f)
            open(dest, "w").write(body)
            print("wrote", os.path.relpath(dest, SERIES))
            report(f, front)
        return

    body, front = render(target, url_map)
    report(target, front)
    sys.stdout.write(body)


if __name__ == "__main__":
    main()
