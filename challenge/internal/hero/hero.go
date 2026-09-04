// Package hero renders the "365 Days of LeetCode Challenge" hero image
// shown at the top of each LinkedIn/Discord post: an HTML template is
// filled with the day counter, topic, difficulty and title, then
// screenshotted to a PNG by an external headless browser.
package hero

import (
	"bytes"
	"fmt"
	"html/template"
	"os/exec"
	"path/filepath"
)

// Data is the set of placeholders the hero template fills in.
type Data struct {
	Day        int
	Total      int
	Topic      string
	Difficulty string
	Title      string
}

// RenderHTML fills the template at templatePath with data and returns
// the rendered HTML.
func RenderHTML(templatePath string, data Data) (string, error) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, view{Data: data, Palette: PaletteFor(data.Day)}); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// Screenshot renders htmlPath (a local file) to a PNG at outPath using a
// headless Chromium instance via Playwright's CLI. Requires
// `npx playwright install chromium` to have been run once on the
// machine — see challenge/README.md. Not covered by automated tests
// since it depends on an external browser binary; verify manually.
func Screenshot(htmlPath, outPath string) error {
	absHTML, err := filepath.Abs(htmlPath)
	if err != nil {
		return err
	}
	cmd := exec.Command("npx", "playwright", "screenshot",
		"--viewport-size=1200,630",
		"file://"+absHTML, outPath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("playwright screenshot: %w: %s", err, out)
	}
	return nil
}

// Palette is one hero card's background gradient.
//
// Only the background varies between days. The brand marks — the
// LeetCode-orange accent, the logo badge, the newsletter mark, the
// author avatar, the type and the layout — are fixed in the template,
// so a reader recognises the series at a glance while still seeing that
// today's post is a new one.
type Palette struct {
	Name string
	From string
	To   string
}

// Palettes are the backgrounds a hero card rotates through, all dark
// enough that white text and the #ffa116 accent clear WCAG AA on both
// gradient stops (enforced by TestPalettesKeepBrandTextReadable).
//
// Hues stay away from the orange/yellow/brown end: the accent has to
// stay obviously the accent, and an orange-on-brown card would blunt the
// one colour doing the branding.
var Palettes = []Palette{
	{Name: "midnight", From: "#1a1a2e", To: "#16213e"},
	{Name: "ocean", From: "#0b1e33", To: "#14406b"},
	{Name: "teal", From: "#0d2427", To: "#12414a"},
	{Name: "forest", From: "#102419", To: "#17402c"},
	{Name: "plum", From: "#241429", To: "#3f2050"},
	{Name: "wine", From: "#2b1220", To: "#4a1c33"},
	{Name: "slate", From: "#1b2027", To: "#2d3743"},
	{Name: "indigo", From: "#191634", To: "#2b2270"},
}

// PaletteFor picks the background for a given challenge day.
//
// It rotates rather than picking at random: random can repeat two days
// running, which is the exact confusion this is meant to remove, and it
// would make re-rendering a day produce a different image than the one
// already posted. Rotating guarantees len(Palettes) distinct days in a
// row and is reproducible.
func PaletteFor(day int) Palette {
	i := (day - 1) % len(Palettes)
	if i < 0 {
		i += len(Palettes)
	}
	return Palettes[i]
}

// view is what the template actually renders against: the caller's Data
// plus the palette derived from it. Deriving it here rather than asking
// callers to pass it means no caller can forget and render a card with
// no background.
type view struct {
	Data
	Palette Palette
}
