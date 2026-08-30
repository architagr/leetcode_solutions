// Command leetcodectl is the deterministic-logic half of the
// /leetcode-content Claude Code skill: each subcommand takes a single
// JSON argument and prints a single JSON result, so the skill can drive
// it without any argv-flag parsing.
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"leetcode_solutions/challenge/internal/cli"
	"leetcode_solutions/challenge/internal/hero"
	"leetcode_solutions/challenge/internal/queue"
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

	default:
		return nil, fmt.Errorf("unknown command %q", cmd)
	}
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
