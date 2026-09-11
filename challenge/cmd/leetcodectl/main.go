// Command leetcodectl is the deterministic-logic half of the
// /leetcode-content Claude Code skill: each subcommand takes a single
// JSON argument and prints a single JSON result, so the skill can drive
// it without any argv-flag parsing.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"leetcode_solutions/challenge/internal/cli"
	"leetcode_solutions/challenge/internal/hero"
	"leetcode_solutions/challenge/internal/queue"
	"leetcode_solutions/challenge/internal/xpost"
)

func main() {
	if len(os.Args) < 2 {
		fail(fmt.Errorf("usage: leetcodectl <command> [json-args]"))
	}
	cmd := os.Args[1]
	payload := []byte("{}")
	if len(os.Args) > 2 {
		payload = []byte(os.Args[2])
	}

	result, err := dispatch(cmd, payload)
	if err != nil {
		fail(err)
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(result); err != nil {
		fail(err)
	}
}

func dispatch(cmd string, payload []byte) (any, error) {
	switch cmd {
	case "resolve":
		var in struct {
			RepoRoot   string `json:"repoRoot"`
			MapPath    string `json:"mapPath"`
			Number     int    `json:"number"`
			Difficulty string `json:"difficulty"`
			Slug       string `json:"slug"`
		}
		if err := json.Unmarshal(payload, &in); err != nil {
			return nil, err
		}
		return cli.Resolve(in.RepoRoot, in.MapPath, in.Number, in.Difficulty, in.Slug)

	case "scaffold-from-submission":
		var in struct {
			RepoRoot    string `json:"repoRoot"`
			SessionPath string `json:"sessionPath"`
			Number      int    `json:"number"`
			Difficulty  string `json:"difficulty"`
			Slug        string `json:"slug"`
			Title       string `json:"title"`
		}
		if err := json.Unmarshal(payload, &in); err != nil {
			return nil, err
		}
		return cli.ScaffoldFromSubmission(in.RepoRoot, in.SessionPath, in.Number, in.Difficulty, in.Slug, in.Title)

	case "reorg":
		var in struct {
			RepoRoot   string `json:"repoRoot"`
			FromPath   string `json:"fromPath"`
			Difficulty string `json:"difficulty"`
			Number     int    `json:"number"`
			Slug       string `json:"slug"`
		}
		if err := json.Unmarshal(payload, &in); err != nil {
			return nil, err
		}
		newPath, err := cli.Reorg(in.RepoRoot, in.FromPath, in.Difficulty, in.Number, in.Slug)
		if err != nil {
			return nil, err
		}
		return map[string]string{"path": newPath}, nil

	case "queue-has":
		var in struct {
			QueuePath string `json:"queuePath"`
			Number    int    `json:"number"`
		}
		if err := json.Unmarshal(payload, &in); err != nil {
			return nil, err
		}
		has, err := cli.QueueHas(in.QueuePath, in.Number)
		if err != nil {
			return nil, err
		}
		return map[string]bool{"has": has}, nil

	case "queue-append":
		var in struct {
			QueuePath string      `json:"queuePath"`
			Entry     queue.Entry `json:"entry"`
		}
		if err := json.Unmarshal(payload, &in); err != nil {
			return nil, err
		}
		day, err := cli.QueueAppend(in.QueuePath, in.Entry)
		if err != nil {
			return nil, err
		}
		return map[string]int{"day": day}, nil

	case "companies-lookup":
		var in struct {
			DatasetPath string `json:"datasetPath"`
			Number      int    `json:"number"`
		}
		if err := json.Unmarshal(payload, &in); err != nil {
			return nil, err
		}
		list, err := cli.CompaniesLookup(in.DatasetPath, in.Number)
		if err != nil {
			return nil, err
		}
		return map[string][]string{"companies": list}, nil

	case "gitmap-update":
		var in struct {
			RepoRoot string `json:"repoRoot"`
			MapPath  string `json:"mapPath"`
		}
		if err := json.Unmarshal(payload, &in); err != nil {
			return nil, err
		}
		updated, newest, err := cli.GitmapUpdate(in.RepoRoot, in.MapPath)
		if err != nil {
			return nil, err
		}
		return map[string]any{"updated": updated, "newestCommit": newest}, nil

	case "hero-render-html":
		var in struct {
			TemplatePath string    `json:"templatePath"`
			OutPath      string    `json:"outPath"`
			Data         hero.Data `json:"data"`
		}
		if err := json.Unmarshal(payload, &in); err != nil {
			return nil, err
		}
		if err := cli.HeroRenderHTML(in.TemplatePath, in.Data, in.OutPath); err != nil {
			return nil, err
		}
		return map[string]string{"outPath": in.OutPath}, nil

	case "hero-generate":
		var in struct {
			TemplatePath string    `json:"templatePath"`
			OutPath      string    `json:"outPath"`
			Data         hero.Data `json:"data"`
		}
		if err := json.Unmarshal(payload, &in); err != nil {
			return nil, err
		}
		return cli.HeroGenerate(in.TemplatePath, in.Data, in.OutPath)

	case "hero-screenshot":
		var in struct {
			HTMLPath string `json:"htmlPath"`
			OutPath  string `json:"outPath"`
		}
		if err := json.Unmarshal(payload, &in); err != nil {
			return nil, err
		}
		if err := hero.Screenshot(in.HTMLPath, in.OutPath); err != nil {
			return nil, err
		}
		return map[string]string{"outPath": in.OutPath}, nil

	case "fetch-list":
		var in struct {
			URL        string `json:"url"`
			Difficulty string `json:"difficulty"`
		}
		if err := json.Unmarshal(payload, &in); err != nil {
			return nil, err
		}
		return cli.FetchList(in.URL, in.Difficulty)

	case "fetch-question":
		var in struct {
			Slug string `json:"slug"`
		}
		if err := json.Unmarshal(payload, &in); err != nil {
			return nil, err
		}
		return cli.FetchQuestion(in.Slug)

	case "post-discord":
		var in struct {
			RepoRoot  string `json:"repoRoot"`
			QueuePath string `json:"queuePath"`
			// Force posts even though a day already went out today. The
			// daily cron leaves it off; it's for a manual catch-up.
			Force bool `json:"force"`
		}
		if err := json.Unmarshal(payload, &in); err != nil {
			return nil, err
		}
		// The webhook URL is a bearer credential for the channel, so it
		// comes from the environment rather than the JSON argument —
		// argv is visible in process listings and gets echoed into logs.
		webhookURL := os.Getenv("DISCORD_WEBHOOK_URL")
		if webhookURL == "" {
			return nil, fmt.Errorf("DISCORD_WEBHOOK_URL is not set")
		}
		var opts []cli.PostOption
		if in.Force {
			opts = append(opts, cli.AllowSameDay())
		}
		return cli.PostDiscord(in.RepoRoot, in.QueuePath, webhookURL, opts...)

	case "verify-x":
		// Read-only: confirms the four OAuth values authenticate and
		// reports which account they post as. Publishes nothing, so it
		// is safe to run as often as setup needs.
		creds := xpost.Credentials{
			ConsumerKey:    os.Getenv("X_API_KEY"),
			ConsumerSecret: os.Getenv("X_API_SECRET"),
			AccessToken:    os.Getenv("X_ACCESS_TOKEN"),
			AccessSecret:   os.Getenv("X_ACCESS_TOKEN_SECRET"),
		}
		acct, err := xpost.Verify(creds)
		if err != nil {
			return nil, err
		}
		return struct {
			ID       string `json:"id"`
			Username string `json:"username"`
			Name     string `json:"name"`
			Message  string `json:"message"`
		}{acct.ID, acct.Username, acct.Name, fmt.Sprintf("credentials authenticate as @%s", acct.Username)}, nil

	case "post-x":
		var in struct {
			RepoRoot  string `json:"repoRoot"`
			QueuePath string `json:"queuePath"`
			// Handle is only used to build the result URL; posting is
			// authorised by the tokens, not by this.
			Handle string `json:"handle"`
			// Force posts even though a day already went out today. The
			// daily cron leaves it off; it's for a manual catch-up.
			Force bool `json:"force"`
		}
		if err := json.Unmarshal(payload, &in); err != nil {
			return nil, err
		}
		// The four OAuth values are account credentials, so they come
		// from the environment rather than the JSON argument — argv is
		// visible in process listings and gets echoed into logs.
		creds := xpost.Credentials{
			ConsumerKey:    os.Getenv("X_API_KEY"),
			ConsumerSecret: os.Getenv("X_API_SECRET"),
			AccessToken:    os.Getenv("X_ACCESS_TOKEN"),
			AccessSecret:   os.Getenv("X_ACCESS_TOKEN_SECRET"),
		}
		if in.Handle == "" {
			in.Handle = os.Getenv("X_HANDLE")
		}
		var opts []cli.PostOption
		if in.Force {
			opts = append(opts, cli.AllowSameDay())
		}
		return cli.PostX(in.RepoRoot, in.QueuePath, in.Handle, creds, opts...)

	case "linkedin-batch":
		var in struct {
			RepoRoot    string `json:"repoRoot"`
			QueuePath   string `json:"queuePath"`
			Destination string `json:"destination"`
			Count       int    `json:"count"`
			OutPath     string `json:"outPath"`
		}
		if err := json.Unmarshal(payload, &in); err != nil {
			return nil, err
		}
		if in.Destination == "" {
			in.Destination = queue.DestinationLinkedInMain
		}
		if in.Count == 0 {
			in.Count = 7
		}
		if in.OutPath == "" {
			in.OutPath = fmt.Sprintf("/tmp/linkedin-batch-%s.md", time.Now().Format("2006-01-02"))
		}
		return cli.LinkedInBatch(in.RepoRoot, in.QueuePath, in.Destination, in.Count, in.OutPath)

	case "x-batch":
		var in struct {
			RepoRoot  string `json:"repoRoot"`
			QueuePath string `json:"queuePath"`
			Count     int    `json:"count"`
			OutPath   string `json:"outPath"`
		}
		if err := json.Unmarshal(payload, &in); err != nil {
			return nil, err
		}
		if in.Count == 0 {
			in.Count = 7
		}
		if in.OutPath == "" {
			in.OutPath = fmt.Sprintf("/tmp/x-batch-%s.md", time.Now().Format("2006-01-02"))
		}
		return cli.XBatch(in.RepoRoot, in.QueuePath, in.Count, in.OutPath)

	case "substack-batch":
		var in struct {
			RepoRoot  string `json:"repoRoot"`
			QueuePath string `json:"queuePath"`
			Count     int    `json:"count"`
			OutPath   string `json:"outPath"`
		}
		if err := json.Unmarshal(payload, &in); err != nil {
			return nil, err
		}
		if in.Count == 0 {
			in.Count = 4
		}
		if in.OutPath == "" {
			in.OutPath = fmt.Sprintf("/tmp/substack-batch-%s.md", time.Now().Format("2006-01-02"))
		}
		return cli.SubstackBatch(in.RepoRoot, in.QueuePath, in.Count, in.OutPath)

	case "mark-content-ready":
		var in struct {
			QueuePath string `json:"queuePath"`
			Numbers   []int  `json:"numbers"`
		}
		if err := json.Unmarshal(payload, &in); err != nil {
			return nil, err
		}
		return cli.MarkContentReady(in.QueuePath, in.Numbers)

	case "mark-posted":
		var in struct {
			QueuePath   string `json:"queuePath"`
			Destination string `json:"destination"`
			Numbers     []int  `json:"numbers"`
		}
		if err := json.Unmarshal(payload, &in); err != nil {
			return nil, err
		}
		return cli.MarkPosted(in.QueuePath, in.Destination, in.Numbers)

	default:
		return nil, fmt.Errorf("unknown command %q", cmd)
	}
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
