// Package leetcode is a minimal client for LeetCode's unofficial,
// undocumented GraphQL API: fetching a custom problem list, a single
// question's statement, and — with a Session attached — the signed-in
// user's own accepted submissions. Premium-only fields (like official
// company tags) are not available through this client.
package leetcode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Problem is one entry from a LeetCode custom problem list.
type Problem struct {
	Number     int
	Title      string
	Slug       string
	Difficulty string
}

// Question is a single question's statement content.
type Question struct {
	Title       string
	Slug        string
	Difficulty  string
	ContentHTML string
}

// Client talks to LeetCode's GraphQL endpoint.
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	// Session, when non-nil, authenticates every request as that
	// LeetCode user. Required for anything user-scoped (submission
	// history and submitted code); the public queries ignore it.
	Session *Session
}

// NewClient returns a Client pointed at LeetCode's real GraphQL endpoint.
func NewClient() *Client {
	return &Client{
		BaseURL:    "https://leetcode.com/graphql",
		HTTPClient: &http.Client{Timeout: 20 * time.Second},
	}
}

// ParseFavoriteSlug extracts the list id from a LeetCode problem-list
// URL, e.g. "https://leetcode.com/problem-list/d0wm7eej/" -> "d0wm7eej".
func ParseFavoriteSlug(rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	parts := strings.Split(strings.Trim(u.Path, "/"), "/")
	for i, p := range parts {
		if p == "problem-list" && i+1 < len(parts) && parts[i+1] != "" {
			return parts[i+1], nil
		}
	}
	return "", fmt.Errorf("not a problem-list URL: %s", rawURL)
}

// FetchProblemList fetches every question in a LeetCode custom problem
// list, optionally filtered by difficulty ("" for all difficulties).
func (c *Client) FetchProblemList(favoriteSlug, difficulty string) ([]Problem, error) {
	query := `query favoriteQuestionList($favoriteSlug: String!, $filter: FavoriteQuestionFilterInput) {
	  favoriteQuestionList(favoriteSlug: $favoriteSlug, filter: $filter) {
	    questions {
	      questionFrontendId
	      title
	      titleSlug
	      difficulty
	    }
	  }
	}`
	filter := map[string]any{}
	if difficulty != "" {
		filter["difficulty"] = strings.ToUpper(difficulty)
	}
	variables := map[string]any{"favoriteSlug": favoriteSlug, "filter": filter}

	var resp struct {
		FavoriteQuestionList struct {
			Questions []struct {
				QuestionFrontendID string `json:"questionFrontendId"`
				Title              string `json:"title"`
				TitleSlug          string `json:"titleSlug"`
				Difficulty         string `json:"difficulty"`
			} `json:"questions"`
		} `json:"favoriteQuestionList"`
	}
	if err := c.doGraphQL(query, variables, &resp); err != nil {
		return nil, err
	}
	problems := make([]Problem, 0, len(resp.FavoriteQuestionList.Questions))
	for _, q := range resp.FavoriteQuestionList.Questions {
		n, err := strconv.Atoi(q.QuestionFrontendID)
		if err != nil {
			continue
		}
		problems = append(problems, Problem{
			Number:     n,
			Title:      q.Title,
			Slug:       q.TitleSlug,
			Difficulty: strings.ToLower(q.Difficulty),
		})
	}
	return problems, nil
}

// FetchQuestionContent fetches a single question's statement by slug.
func (c *Client) FetchQuestionContent(slug string) (Question, error) {
	query := `query questionContent($titleSlug: String!) {
	  question(titleSlug: $titleSlug) {
	    title
	    titleSlug
	    difficulty
	    content
	  }
	}`
	variables := map[string]any{"titleSlug": slug}

	var resp struct {
		Question struct {
			Title      string `json:"title"`
			TitleSlug  string `json:"titleSlug"`
			Difficulty string `json:"difficulty"`
			Content    string `json:"content"`
		} `json:"question"`
	}
	if err := c.doGraphQL(query, variables, &resp); err != nil {
		return Question{}, err
	}
	q := resp.Question
	return Question{
		Title:       q.Title,
		Slug:        q.TitleSlug,
		Difficulty:  strings.ToLower(q.Difficulty),
		ContentHTML: q.Content,
	}, nil
}

func (c *Client) doGraphQL(query string, variables map[string]any, out any) error {
	body, err := json.Marshal(map[string]any{"query": query, "variables": variables})
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, c.BaseURL, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "leetcodectl/1.0")
	if c.Session != nil {
		// LeetCode authenticates GraphQL by cookie and rejects any
		// mutation-ish query whose x-csrftoken does not echo the
		// csrftoken cookie. Referer is checked too.
		req.AddCookie(&http.Cookie{Name: "LEETCODE_SESSION", Value: c.Session.LeetCodeSession})
		if c.Session.CSRFToken != "" {
			req.AddCookie(&http.Cookie{Name: "csrftoken", Value: c.Session.CSRFToken})
			req.Header.Set("x-csrftoken", c.Session.CSRFToken)
		}
		req.Header.Set("Referer", "https://leetcode.com/")
	}

	client := c.HTTPClient
	if client == nil {
		client = http.DefaultClient
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		snippet, _ := io.ReadAll(io.LimitReader(resp.Body, 500))
		return fmt.Errorf("leetcode graphql: unexpected status %d: %s", resp.StatusCode, snippet)
	}

	var envelope struct {
		Data   json.RawMessage `json:"data"`
		Errors []struct {
			Message string `json:"message"`
		} `json:"errors"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&envelope); err != nil {
		return err
	}
	if len(envelope.Errors) > 0 {
		return fmt.Errorf("leetcode graphql: %s", envelope.Errors[0].Message)
	}
	return json.Unmarshal(envelope.Data, out)
}
