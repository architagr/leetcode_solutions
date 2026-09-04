package hero

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

func TestRenderHTML(t *testing.T) {
	tmplPath := filepath.Join(t.TempDir(), "template.html")
	tmpl := `<h1>Day {{.Day}} / {{.Total}}</h1><h2>{{.Title}}</h2><p>{{.Topic}} - {{.Difficulty}}</p>`
	if err := os.WriteFile(tmplPath, []byte(tmpl), 0o644); err != nil {
		t.Fatal(err)
	}

	html, err := RenderHTML(tmplPath, Data{
		Day: 23, Total: 365, Topic: "Binary Tree", Difficulty: "Easy", Title: "Univalued Binary Tree",
	})
	if err != nil {
		t.Fatalf("RenderHTML: %v", err)
	}
	for _, want := range []string{"Day 23 / 365", "Univalued Binary Tree", "Binary Tree - Easy"} {
		if !strings.Contains(html, want) {
			t.Errorf("rendered HTML missing %q, got: %s", want, html)
		}
	}
}

func TestRenderHTMLRealTemplate(t *testing.T) {
	html, err := RenderHTML("../../hero_template.html", Data{
		Day: 1, Total: 365, Topic: "Arrays", Difficulty: "Easy", Title: "Two Sum",
	})
	if err != nil {
		t.Fatalf("RenderHTML with real template: %v", err)
	}
	for _, want := range []string{"1", "365", "Arrays", "Easy", "Two Sum"} {
		if !strings.Contains(html, want) {
			t.Errorf("rendered real template missing %q", want)
		}
	}
}

// --- palette ---

// relativeLuminance and contrastRatio are the WCAG 2.x definitions,
// kept here rather than in the package because nothing but these tests
// needs them: they exist to prove the palettes are readable, not to be
// called at render time.
func relativeLuminance(hexColor string) float64 {
	var r, g, b int
	if _, err := fmt.Sscanf(hexColor, "#%02x%02x%02x", &r, &g, &b); err != nil {
		panic("bad hex " + hexColor)
	}
	lin := func(c int) float64 {
		s := float64(c) / 255
		if s <= 0.03928 {
			return s / 12.92
		}
		return math.Pow((s+0.055)/1.055, 2.4)
	}
	return 0.2126*lin(r) + 0.7152*lin(g) + 0.0722*lin(b)
}

func contrastRatio(a, b string) float64 {
	la, lb := relativeLuminance(a), relativeLuminance(b)
	if la < lb {
		la, lb = lb, la
	}
	return (la + 0.05) / (lb + 0.05)
}

func TestPaletteForIsStablePerDayAndVariesBetweenDays(t *testing.T) {
	// Re-rendering a day must reproduce its image exactly, so the choice
	// is derived from the day, never drawn at random.
	if PaletteFor(7) != PaletteFor(7) {
		t.Error("PaletteFor(7) changed between calls")
	}

	// A full cycle's worth of consecutive days must all look different —
	// that's the whole point: two posts in a row must not read as one
	// post sent twice.
	seen := map[string]int{}
	for day := 1; day <= len(Palettes); day++ {
		seen[PaletteFor(day).Name]++
	}
	if len(seen) != len(Palettes) {
		t.Errorf("%d consecutive days used only %d of %d palettes: %v",
			len(Palettes), len(seen), len(Palettes), seen)
	}

	// Day 1 keeps the original scheme so the first post still matches
	// what already went out.
	if got := PaletteFor(1); got.From != "#1a1a2e" || got.To != "#16213e" {
		t.Errorf("day 1 palette = %+v, want the original midnight gradient", got)
	}
}

func TestPaletteForHandlesOutOfRangeDays(t *testing.T) {
	// Day is 1-based everywhere, but a zero or negative value must pick
	// a palette rather than panic on a negative modulus.
	for _, day := range []int{-366, -1, 0, 1, 366, 100000} {
		if p := PaletteFor(day); p.Name == "" {
			t.Errorf("PaletteFor(%d) returned the zero palette", day)
		}
	}
}

func TestPalettesKeepBrandTextReadable(t *testing.T) {
	// Every colour the template paints on the card background. The
	// accent is LeetCode orange and the brand's fixed point, so a
	// background that dulls it is not usable no matter how nice it looks.
	foregrounds := map[string]string{
		"title/day white": "#ffffff",
		"secondary grey":  "#cccccc",
		"brand accent":    "#ffa116",
	}
	// 4.5:1 is WCAG AA for normal-size text. The card's small text is
	// 16-20px, so the strict threshold is the right one to hold to.
	const minRatio = 4.5

	for _, p := range Palettes {
		for _, stop := range []string{p.From, p.To} {
			for label, fg := range foregrounds {
				if got := contrastRatio(fg, stop); got < minRatio {
					t.Errorf("palette %q: %s (%s) on %s = %.2f:1, want >= %.1f:1",
						p.Name, label, fg, stop, got, minRatio)
				}
			}
		}
	}
}

func TestPalettesAreDistinctAndWellFormed(t *testing.T) {
	names := map[string]bool{}
	gradients := map[string]bool{}
	hex := regexp.MustCompile(`^#[0-9a-f]{6}$`)
	for _, p := range Palettes {
		if names[p.Name] {
			t.Errorf("duplicate palette name %q", p.Name)
		}
		names[p.Name] = true
		if gradients[p.From+p.To] {
			t.Errorf("palette %q duplicates another palette's gradient", p.Name)
		}
		gradients[p.From+p.To] = true
		for _, c := range []string{p.From, p.To} {
			if !hex.MatchString(c) {
				t.Errorf("palette %q: %q is not a lowercase 6-digit hex colour", p.Name, c)
			}
		}
	}
}

func TestRenderHTMLPaintsTheDaysGradient(t *testing.T) {
	// The gradient has to survive html/template's CSS context escaping
	// and land in the real template's background declaration.
	day := 5
	want := PaletteFor(day)
	html, err := RenderHTML("../../hero_template.html", Data{
		Day: day, Total: 365, Topic: "Binary Tree", Difficulty: "Easy", Title: "Balanced Binary Tree",
	})
	if err != nil {
		t.Fatalf("RenderHTML: %v", err)
	}
	gradient := fmt.Sprintf("linear-gradient(135deg, %s 0%%, %s 100%%)", want.From, want.To)
	if !strings.Contains(html, gradient) {
		t.Errorf("rendered HTML missing %q", gradient)
	}
	// Branding must not move with the background.
	if !strings.Contains(html, "#ffa116") {
		t.Error("rendered HTML lost the brand accent colour")
	}
}
