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
)

// Meta is the front matter block of a long-form post.
type Meta struct {
	// Title is the SEO/social-card headline. It is deliberately separate
	// from the article's own H1: the H1 is written to be read in context,
	// the meta title has to work as a standalone line in a search result.
	Title string `yaml:"meta_title"`
	// Description is the card and search-result blurb.
	Description string `yaml:"meta_description"`
	// Canonical is the URL this article should point search engines at
	// when the same piece is published in more than one place. Cross
	// posting an article to LinkedIn and Substack without one makes the
	// two copies compete as duplicates; with one, the copies credit a
	// single original.
	Canonical string `yaml:"canonical_url,omitempty"`
}

// IsZero reports whether no front matter fields were set at all, which
// is how a file written before this feature existed reads.
func (m Meta) IsZero() bool {
	return m.Title == "" && m.Description == "" && m.Canonical == ""
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
