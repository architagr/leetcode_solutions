package discordpost

import (
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"leetcode_solutions/challenge/internal/queue"
)

func ts(s string) *string { return &s }

func TestSelectNextPicksOldestUnpostedForThatDestination(t *testing.T) {
	q := &queue.Queue{Entries: []queue.Entry{
		// Deliberately out of day order: selection must go by day, not file order.
		{Day: 3, Number: 257},
		{Day: 1, Number: 104, PostedAt: map[string]*string{queue.DestinationDiscord: ts("2026-09-01T09:00:00Z")}},
		{Day: 2, Number: 108, PostedAt: map[string]*string{queue.DestinationLinkedInMain: ts("2026-09-02T09:00:00Z")}},
	}}

	got, ok := SelectNext(q, queue.DestinationDiscord)
	if !ok {
		t.Fatal("SelectNext: ok = false, want true")
	}
	// Day 1 is already on discord; day 2 is only on linkedin, so it's next for discord.
	if got.Day != 2 {
		t.Errorf("SelectNext picked day %d, want 2", got.Day)
	}
}

func TestSelectNextEmptyAndFullyPosted(t *testing.T) {
	if _, ok := SelectNext(&queue.Queue{}, queue.DestinationDiscord); ok {
		t.Error("SelectNext on an empty queue = true, want false")
	}

	allPosted := &queue.Queue{Entries: []queue.Entry{
		{Day: 1, Number: 104, PostedAt: map[string]*string{queue.DestinationDiscord: ts("2026-09-01T09:00:00Z")}},
	}}
	if _, ok := SelectNext(allPosted, queue.DestinationDiscord); ok {
		t.Error("SelectNext with everything posted = true, want false")
	}
}

func TestSelectNextTreatsExplicitNullAsUnposted(t *testing.T) {
	q := &queue.Queue{Entries: []queue.Entry{
		{Day: 1, Number: 104, PostedAt: map[string]*string{queue.DestinationDiscord: nil}},
	}}
	got, ok := SelectNext(q, queue.DestinationDiscord)
	if !ok || got.Day != 1 {
		t.Errorf("SelectNext = (%+v, %v), want day 1 selected", got, ok)
	}
}

func TestValidateMessageLength(t *testing.T) {
	if err := ValidateMessage(strings.Repeat("a", MaxMessageChars)); err != nil {
		t.Errorf("exactly at the limit should be allowed: %v", err)
	}
	if err := ValidateMessage(strings.Repeat("a", MaxMessageChars-1)); err != nil {
		t.Errorf("just under the limit should be allowed: %v", err)
	}
	err := ValidateMessage(strings.Repeat("a", MaxMessageChars+1))
	if err == nil {
		t.Fatal("over the limit should be rejected")
	}
	if !strings.Contains(err.Error(), "2001") {
		t.Errorf("error should name the actual length, got %v", err)
	}
	if err := ValidateMessage("   \n  "); err == nil {
		t.Error("an effectively empty message should be rejected")
	}
}

func TestValidateMessageCountsRunesNotBytes(t *testing.T) {
	// Discord counts characters. These are 3 bytes each in UTF-8, so a
	// byte-based check would reject a message that Discord accepts.
	msg := strings.Repeat("—", MaxMessageChars)
	if err := ValidateMessage(msg); err != nil {
		t.Errorf("%d multi-byte runes should be allowed: %v", MaxMessageChars, err)
	}
}

func TestPostSendsMessageAndHeroAttachment(t *testing.T) {
	dir := t.TempDir()
	heroPath := filepath.Join(dir, "HERO.png")
	if err := os.WriteFile(heroPath, []byte("\x89PNG\r\n\x1a\nfake-image-bytes"), 0o644); err != nil {
		t.Fatal(err)
	}

	var gotPayload, gotFile string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, params, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
		if err != nil {
			t.Errorf("Content-Type: %v", err)
		}
		mr := multipart.NewReader(r.Body, params["boundary"])
		for {
			part, err := mr.NextPart()
			if err == io.EOF {
				break
			}
			if err != nil {
				t.Fatalf("NextPart: %v", err)
			}
			body, _ := io.ReadAll(part)
			switch part.FormName() {
			case "payload_json":
				gotPayload = string(body)
			case "files[0]":
				gotFile = string(body)
			}
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	if err := Post(server.URL, "hello from day 1", heroPath); err != nil {
		t.Fatalf("Post: %v", err)
	}
	if !strings.Contains(gotPayload, "hello from day 1") {
		t.Errorf("payload_json missing the message: %s", gotPayload)
	}
	if !strings.Contains(gotPayload, "attachment://HERO.png") {
		t.Errorf("payload_json should reference the attached hero image: %s", gotPayload)
	}
	if !strings.Contains(gotFile, "fake-image-bytes") {
		t.Errorf("hero image bytes were not attached, got %q", gotFile)
	}
}

func TestPostWithoutHeroImageStillSends(t *testing.T) {
	var gotPayload string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, params, _ := mime.ParseMediaType(r.Header.Get("Content-Type"))
		mr := multipart.NewReader(r.Body, params["boundary"])
		for {
			part, err := mr.NextPart()
			if err == io.EOF {
				break
			}
			if err != nil {
				t.Fatalf("NextPart: %v", err)
			}
			if part.FormName() == "payload_json" {
				body, _ := io.ReadAll(part)
				gotPayload = string(body)
			}
		}
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	if err := Post(server.URL, "no hero today", ""); err != nil {
		t.Fatalf("Post without hero: %v", err)
	}
	if !strings.Contains(gotPayload, "no hero today") {
		t.Errorf("payload_json missing the message: %s", gotPayload)
	}
	if strings.Contains(gotPayload, "attachment://") {
		t.Errorf("payload must not reference an attachment when none was sent: %s", gotPayload)
	}
}

func TestPostSurfacesNon2xx(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"message":"Invalid Webhook Token"}`))
	}))
	defer server.Close()

	err := Post(server.URL, "anything", "")
	if err == nil {
		t.Fatal("a 400 response must be an error")
	}
	if !strings.Contains(err.Error(), "400") || !strings.Contains(err.Error(), "Invalid Webhook Token") {
		t.Errorf("error should carry the status and Discord's reason, got: %v", err)
	}
}

func TestPostSurfacesRateLimit(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Retry-After", "42")
		w.WriteHeader(http.StatusTooManyRequests)
		_, _ = w.Write([]byte(`{"message":"You are being rate limited.","retry_after":42}`))
	}))
	defer server.Close()

	err := Post(server.URL, "anything", "")
	if err == nil {
		t.Fatal("a 429 response must be an error")
	}
	if !strings.Contains(err.Error(), "rate limited") || !strings.Contains(err.Error(), "42") {
		t.Errorf("error should name the rate limit and retry-after, got: %v", err)
	}
}

func TestPostMissingHeroFileIsAnError(t *testing.T) {
	err := Post("http://example.invalid", "msg", filepath.Join(t.TempDir(), "nope.png"))
	if err == nil {
		t.Fatal("a missing hero file must be an error, not a silent skip")
	}
}

func TestPostedOnFindsAnEntrySentThatUTCDay(t *testing.T) {
	q := &queue.Queue{Entries: []queue.Entry{
		{Day: 1, Number: 104, PostedAt: map[string]*string{queue.DestinationDiscord: ts("2026-09-03T06:16:03Z")}},
		{Day: 2, Number: 108, PostedAt: map[string]*string{queue.DestinationDiscord: ts("2026-09-04T05:31:13Z")}},
		{Day: 3, Number: 257},
	}}

	// A second run later the same UTC day must see day 2 as today's post.
	got, ok := PostedOn(q, queue.DestinationDiscord, time.Date(2026, 9, 4, 23, 59, 0, 0, time.UTC))
	if !ok {
		t.Fatal("PostedOn: ok = false, want true")
	}
	if got.Day != 2 {
		t.Errorf("PostedOn found day %d, want 2", got.Day)
	}

	// The next UTC day is clear again, even a minute past midnight.
	if _, ok := PostedOn(q, queue.DestinationDiscord, time.Date(2026, 9, 5, 0, 1, 0, 0, time.UTC)); ok {
		t.Error("PostedOn on a fresh day = true, want false")
	}
}

func TestPostedOnComparesInUTCAndIgnoresOtherDestinations(t *testing.T) {
	q := &queue.Queue{Entries: []queue.Entry{
		{Day: 1, Number: 104, PostedAt: map[string]*string{queue.DestinationLinkedInMain: ts("2026-09-04T05:00:00Z")}},
		{Day: 2, Number: 108, PostedAt: map[string]*string{queue.DestinationDiscord: ts("bogus")}},
	}}

	// Another destination's post that day must not block discord.
	if _, ok := PostedOn(q, queue.DestinationDiscord, time.Date(2026, 9, 4, 6, 0, 0, 0, time.UTC)); ok {
		t.Error("PostedOn matched a different destination, want false")
	}

	// A now in a non-UTC zone still compares by UTC calendar day: 2026-09-05
	// 04:00 +05:30 is 2026-09-04 22:30 UTC.
	q.Entries[1].PostedAt[queue.DestinationDiscord] = ts("2026-09-04T05:31:13Z")
	ist := time.FixedZone("IST", 5*60*60+30*60)
	if _, ok := PostedOn(q, queue.DestinationDiscord, time.Date(2026, 9, 5, 4, 0, 0, 0, ist)); !ok {
		t.Error("PostedOn with a non-UTC now = false, want true")
	}
}
