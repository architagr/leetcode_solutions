// Package postmeta reads the optional YAML front matter at the top of a
// long-form post file — the SEO metadata LinkedIn and Substack ask for
// separately from the body when you publish an article.
//
// The metadata lives in the post file rather than in queue.yaml because
// it is written with the article, is reviewed with the article, and is
// meaningless without it. The queue tracks scheduling; this tracks
// content.
//
// Front matter is optional throughout. Every article written before this
// existed has none, and a batch of days must still prepare cleanly when
// some of them predate the field — a missing meta title is something to
// report, not something to fail on.
package postmeta

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"gopkg.in/yaml.v3"
)

// Length budgets, in characters. These are where the platforms and
// Google truncate with an ellipsis, not hard limits either service
// enforces — so going over is a warning about how the card will look in
// a feed or a search result, never an error that blocks publishing.
const (
	MaxTitleChars       = 60
	MaxDescriptionChars = 155

	// MaxTags is a house rule rather than a Substack limit: past a
	// handful, tags stop describing the piece and start looking like
	// keyword stuffing to a reader who can see them all.
	MaxTags = 5
)

// Meta is the front matter block of a long-form post.
//
// There is deliberately no canonical URL field. Neither platform this
// serves lets a publisher set one: LinkedIn's article composer has no
// such field, and Substack treats its own domain as canonical and
// exposes only an SEO title and subtitle. Carrying the value anyway
// would have meant a front matter field that nothing could ever apply.
// The duplicate-content problem it would have solved is handled the only
// way these platforms allow — both copies link prominently back to the
// repo in the body.
type Meta struct {
	// Title is the SEO/social-card headline. It is deliberately separate
	// from the article's own H1: the H1 is written to be read in context,
	// the meta title has to work as a standalone line in a search result.
	Title string `yaml:"meta_title"`
	// Description is the card and search-result blurb.
	Description string `yaml:"meta_description"`
	// Tags are Substack's post tags, set in the publish dialog rather
	// than written into the body. They are the reason a hashtag line
	// does not belong in POST_SUBSTACK.md: Substack indexes these, and
	// treats a #hashtag in the prose as prose.
	Tags []string `yaml:"tags,omitempty"`
}

// IsZero reports whether no front matter fields were set at all, which
// is how a file written before this feature existed reads.
func (m Meta) IsZero() bool {
	return m.Title == "" && m.Description == "" && len(m.Tags) == 0
}

// Warnings describes anything about the metadata that will publish, but
// won't look right. Never an error: a long title still posts, it just
// gets cut off in a feed card.
func (m Meta) Warnings() []string {
	var out []string
	if m.Title == "" {
		out = append(out, "no meta_title — the platform will fall back to the first heading, which is usually too long for a card")
	} else if n := utf8.RuneCountInString(m.Title); n > MaxTitleChars {
		out = append(out, fmt.Sprintf("meta_title is %d characters, over the %d that fit before a search result truncates it", n, MaxTitleChars))
	}
	if len(m.Tags) > MaxTags {
		out = append(out, fmt.Sprintf("%d tags — Substack shows the first %d, and a long list reads as keyword stuffing", len(m.Tags), MaxTags))
	}
	if m.Description == "" {
		out = append(out, "no meta_description — the platform will excerpt the opening line, which is rarely the line you'd choose")
	} else if n := utf8.RuneCountInString(m.Description); n > MaxDescriptionChars {
		out = append(out, fmt.Sprintf("meta_description is %d characters, over the %d that fit before a search result truncates it", n, MaxDescriptionChars))
	}
	return out
}

// Parse splits a post file into its front matter and its body.
//
// Front matter is the leading block delimited by --- lines, as in every
// static site generator. A file with no such block parses as empty
// metadata plus the whole file as the body, with no error — that is the
// normal shape of every article written before this package existed.
func Parse(content string) (Meta, string, error) {
	// Editors and git checkouts both introduce CRLF; normalising here
	// means the delimiter match doesn't silently fail on a Windows-edited
	// file and drop the metadata.
	normalised := strings.ReplaceAll(content, "\r\n", "\n")
	trimmed := strings.TrimLeft(normalised, "\ufeff \t\n")

	if !strings.HasPrefix(trimmed, "---\n") {
		return Meta{}, strings.TrimSpace(normalised), nil
	}

	rest := trimmed[len("---\n"):]
	end := strings.Index(rest, "\n---")
	if end < 0 {
		// An opening delimiter with no closing one means the author
		// started a block and didn't finish it. Treating the whole file
		// as body would silently publish the YAML, so say so instead.
		return Meta{}, "", fmt.Errorf("front matter opens with --- but is never closed")
	}

	var m Meta
	if err := yaml.Unmarshal([]byte(rest[:end]), &m); err != nil {
		return Meta{}, "", fmt.Errorf("parsing front matter: %w", err)
	}

	body := rest[end+len("\n---"):]
	// The closing delimiter may or may not be followed by a newline;
	// trimming the body handles both without a special case.
	return m, strings.TrimSpace(body), nil
}
