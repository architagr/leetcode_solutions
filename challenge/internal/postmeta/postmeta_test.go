package postmeta

import (
	"strings"
	"testing"
)

func TestParseReadsFrontMatterAndBody(t *testing.T) {
	m, body, err := Parse(`---
meta_title: "Recover a contaminated tree without rebuilding it"
meta_description: "The recovered values are 2i+1 and 2i+2, so the path to any target is its binary representation."
---

# Day 9

The body starts here.
`)
	if err != nil {
		t.Fatal(err)
	}
	if m.Title != "Recover a contaminated tree without rebuilding it" {
		t.Errorf("Title = %q", m.Title)
	}
	if !strings.HasPrefix(m.Description, "The recovered values are 2i+1") {
		t.Errorf("Description = %q", m.Description)
	}
	if !strings.HasPrefix(body, "# Day 9") {
		t.Errorf("body = %q, want it to start at the heading", body)
	}
	if strings.Contains(body, "meta_title") {
		t.Error("the front matter leaked into the body")
	}
}

// Every article written before this package existed has no front matter,
// and a batch containing one must still prepare.
func TestParseFileWithoutFrontMatterIsNotAnError(t *testing.T) {
	m, body, err := Parse("# Just an article\n\nNo front matter here.\n")
	if err != nil {
		t.Fatalf("a file without front matter must parse: %v", err)
	}
	if !m.IsZero() {
		t.Errorf("Meta = %+v, want zero", m)
	}
	if !strings.HasPrefix(body, "# Just an article") {
		t.Errorf("body = %q, want the whole file", body)
	}
}

// A --- inside the article (a horizontal rule, which these posts use)
// must not be mistaken for front matter.
func TestParseIgnoresADelimiterThatIsNotAtTheTop(t *testing.T) {
	m, body, err := Parse("# Article\n\n---\n\nA section break.\n")
	if err != nil {
		t.Fatal(err)
	}
	if !m.IsZero() {
		t.Errorf("Meta = %+v, want zero — the --- is a horizontal rule", m)
	}
	if !strings.Contains(body, "A section break.") {
		t.Errorf("body = %q, want the whole file", body)
	}
}

// Publishing the raw YAML because the author forgot a closing delimiter
// is worse than refusing to prepare the batch.
func TestParseUnclosedFrontMatterIsAnError(t *testing.T) {
	if _, _, err := Parse("---\nmeta_title: oops\n\n# Article\n"); err == nil {
		t.Fatal("front matter that is never closed must be an error")
	}
}

func TestParseMalformedYAMLIsAnError(t *testing.T) {
	if _, _, err := Parse("---\nmeta_title: \"unterminated\nmeta_description: x\n---\nbody\n"); err == nil {
		t.Fatal("unparseable front matter must be an error")
	}
}

func TestParseHandlesCRLFAndALeadingBOM(t *testing.T) {
	m, body, err := Parse("\ufeff---\r\nmeta_title: Windows\r\n---\r\n\r\nBody.\r\n")
	if err != nil {
		t.Fatal(err)
	}
	if m.Title != "Windows" {
		t.Errorf("Title = %q, want Windows", m.Title)
	}
	if body != "Body." {
		t.Errorf("body = %q, want Body.", body)
	}
}

func TestWarningsFlagMissingAndOverLongFields(t *testing.T) {
	if got := (Meta{}).Warnings(); len(got) != 2 {
		t.Errorf("empty meta warnings = %v, want one for each of title and description", got)
	}

	fine := Meta{Title: "A short enough title", Description: "A description comfortably inside the budget."}
	if got := fine.Warnings(); len(got) != 0 {
		t.Errorf("well-formed meta warnings = %v, want none", got)
	}

	long := Meta{
		Title:       strings.Repeat("a", MaxTitleChars+1),
		Description: strings.Repeat("b", MaxDescriptionChars+1),
	}
	got := long.Warnings()
	if len(got) != 2 {
		t.Fatalf("over-long meta warnings = %v, want two", got)
	}
	for _, w := range got {
		if !strings.Contains(w, "truncates") {
			t.Errorf("warning should say what goes wrong: %q", w)
		}
	}
}

// Warnings are about how a card renders, so they count characters, not
// bytes — an em dash must not cost three of the budget.
func TestWarningsCountRunesNotBytes(t *testing.T) {
	m := Meta{
		Title:       strings.Repeat("—", MaxTitleChars),
		Description: strings.Repeat("—", MaxDescriptionChars),
	}
	if got := m.Warnings(); len(got) != 0 {
		t.Errorf("warnings = %v, want none for a title and description at exactly the rune budget", got)
	}
}

func TestParseReadsSubstackTags(t *testing.T) {
	m, _, err := Parse("---\nmeta_title: T\ntags: [golang, binary-tree, recursion]\n---\n\nBody.\n")
	if err != nil {
		t.Fatal(err)
	}
	if len(m.Tags) != 3 || m.Tags[0] != "golang" || m.Tags[2] != "recursion" {
		t.Errorf("Tags = %v, want [golang binary-tree recursion]", m.Tags)
	}
	if m.IsZero() {
		t.Error("meta carrying tags is not zero")
	}
}

// Tags are the reason there is no hashtag line in a Substack post, so a
// file carrying only tags must still read as having front matter.
func TestTagsAloneAreNotZeroMeta(t *testing.T) {
	if (Meta{Tags: []string{"golang"}}).IsZero() {
		t.Error("Meta with tags reported as zero")
	}
}

func TestWarningsFlagTooManyTags(t *testing.T) {
	tags := make([]string, MaxTags+1)
	for i := range tags {
		tags[i] = "tag"
	}
	m := Meta{Title: "A title", Description: "A description.", Tags: tags}
	got := m.Warnings()
	if len(got) != 1 || !strings.Contains(got[0], "keyword stuffing") {
		t.Errorf("warnings = %v, want one about too many tags", got)
	}

	m.Tags = tags[:MaxTags]
	if got := m.Warnings(); len(got) != 0 {
		t.Errorf("warnings = %v, want none at exactly the tag budget", got)
	}
}
