// Package discordpost posts one queued challenge entry per run to a
// Discord channel via an incoming webhook.
//
// The message body is the entry's POST_DISCORD.md, sent verbatim: that
// file is authored to be exactly what lands in the channel, so nothing
// here parses, trims or rewrites it. The only thing this package adds is
// the hero image, attached as a file because these images live in the
// repo rather than behind a public URL an embed could point at.
package discordpost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode/utf8"
)

// MaxMessageChars is Discord's hard limit on a message's content field.
// Discord counts characters, not bytes.
const MaxMessageChars = 2000

// httpClient is a package-level client so a slow or hung Discord doesn't
// wedge a scheduled run forever.
var httpClient = &http.Client{Timeout: 30 * time.Second}

// ValidateMessage checks a message body against Discord's limits before
// any network call. Over-length content is a bug in the source file, not
// something to silently truncate: truncating would publish a message
// that stops mid-sentence.
func ValidateMessage(message string) error {
	if strings.TrimSpace(message) == "" {
		return fmt.Errorf("message is empty")
	}
	if n := utf8.RuneCountInString(message); n > MaxMessageChars {
		return fmt.Errorf("message is %d characters, over Discord's %d limit", n, MaxMessageChars)
	}
	return nil
}

// payload is the JSON body of a webhook execution.
type payload struct {
	Content string  `json:"content"`
	Embeds  []embed `json:"embeds,omitempty"`
}

type embed struct {
	Image *embedImage `json:"image,omitempty"`
}

type embedImage struct {
	URL string `json:"url"`
}

// Post sends message to a Discord incoming webhook, attaching the image
// at heroPath if one is given. An empty heroPath posts text only.
func Post(webhookURL, message, heroPath string) error {
	if err := ValidateMessage(message); err != nil {
		return err
	}

	body := payload{Content: message}
	var heroBytes []byte
	var heroName string
	if heroPath != "" {
		var err error
		heroBytes, err = os.ReadFile(heroPath)
		if err != nil {
			return fmt.Errorf("reading hero image: %w", err)
		}
		heroName = filepath.Base(heroPath)
		body.Embeds = []embed{{Image: &embedImage{URL: "attachment://" + heroName}}}
	}

	payloadJSON, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("encoding payload: %w", err)
	}

	var buf bytes.Buffer
	form := multipart.NewWriter(&buf)
	if err := form.WriteField("payload_json", string(payloadJSON)); err != nil {
		return fmt.Errorf("writing payload field: %w", err)
	}
	if heroName != "" {
		part, err := form.CreateFormFile("files[0]", heroName)
		if err != nil {
			return fmt.Errorf("creating attachment part: %w", err)
		}
		if _, err := part.Write(heroBytes); err != nil {
			return fmt.Errorf("writing attachment: %w", err)
		}
	}
	if err := form.Close(); err != nil {
		return fmt.Errorf("closing multipart form: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, webhookURL, &buf)
	if err != nil {
		return fmt.Errorf("building request: %w", err)
	}
	req.Header.Set("Content-Type", form.FormDataContentType())

	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("posting to Discord: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	respBody, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
	if resp.StatusCode == http.StatusTooManyRequests {
		return fmt.Errorf("discord rate limited this webhook (retry after %s): %s",
			resp.Header.Get("Retry-After"), strings.TrimSpace(string(respBody)))
	}
	return fmt.Errorf("discord returned %d: %s", resp.StatusCode, strings.TrimSpace(string(respBody)))
}
