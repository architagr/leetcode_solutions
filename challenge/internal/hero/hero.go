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
	if err := tmpl.Execute(&buf, data); err != nil {
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
