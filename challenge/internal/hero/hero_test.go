package hero

import (
	"os"
	"path/filepath"
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
