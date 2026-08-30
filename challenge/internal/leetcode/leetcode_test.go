package leetcode

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestParseFavoriteSlug(t *testing.T) {
	cases := map[string]string{
		"https://leetcode.com/problem-list/d0wm7eej/":   "d0wm7eej",
		"https://leetcode.com/problem-list/d0wm7eej":    "d0wm7eej",
		"https://leetcode.com/problem-list/abc123/?x=1": "abc123",
	}
	for url, want := range cases {
		got, err := ParseFavoriteSlug(url)
		if err != nil {
			t.Fatalf("ParseFavoriteSlug(%q): %v", url, err)
		}
		if got != want {
			t.Errorf("ParseFavoriteSlug(%q) = %q, want %q", url, got, want)
		}
	}
	if _, err := ParseFavoriteSlug("https://leetcode.com/problems/two-sum/"); err == nil {
		t.Error("expected error for a non-problem-list URL")
	}
}

func TestFetchProblemList(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Variables struct {
				FavoriteSlug string `json:"favoriteSlug"`
			} `json:"variables"`
		}
		json.NewDecoder(r.Body).Decode(&req)
		if req.Variables.FavoriteSlug != "d0wm7eej" {
			t.Errorf("favoriteSlug = %q", req.Variables.FavoriteSlug)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"favoriteQuestionList":{"questions":[
			{"questionFrontendId":"965","title":"Univalued Binary Tree","titleSlug":"univalued-binary-tree","difficulty":"Easy"},
			{"questionFrontendId":"1022","title":"Sum of Root To Leaf Binary Numbers","titleSlug":"sum-of-root-to-leaf-binary-numbers","difficulty":"Easy"}
		]}}}`))
	}))
	defer server.Close()

	client := &Client{BaseURL: server.URL, HTTPClient: server.Client()}
	problems, err := client.FetchProblemList("d0wm7eej", "easy")
	if err != nil {
		t.Fatalf("FetchProblemList: %v", err)
	}
	if len(problems) != 2 {
		t.Fatalf("got %d problems, want 2", len(problems))
	}
	if problems[0].Number != 965 || problems[0].Slug != "univalued-binary-tree" || problems[0].Difficulty != "easy" {
		t.Errorf("problems[0] = %+v", problems[0])
	}
}

func TestFetchQuestionContent(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"data":{"question":{"title":"Univalued Binary Tree","titleSlug":"univalued-binary-tree","difficulty":"Easy","content":"<p>A binary tree is univalued...</p>"}}}`))
	}))
	defer server.Close()

	client := &Client{BaseURL: server.URL, HTTPClient: server.Client()}
	q, err := client.FetchQuestionContent("univalued-binary-tree")
	if err != nil {
		t.Fatalf("FetchQuestionContent: %v", err)
	}
	if q.Title != "Univalued Binary Tree" || q.Difficulty != "easy" || q.ContentHTML == "" {
		t.Errorf("q = %+v", q)
	}
}

func TestFetchProblemListGraphQLError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errors":[{"message":"something went wrong"}],"data":null}`))
	}))
	defer server.Close()

	client := &Client{BaseURL: server.URL, HTTPClient: server.Client()}
	problems, err := client.FetchProblemList("bad-slug", "")
	if err == nil {
		t.Fatalf("FetchProblemList: expected error, got problems = %+v", problems)
	}
	if !strings.Contains(err.Error(), "something went wrong") {
		t.Errorf("error = %q, want it to contain %q", err.Error(), "something went wrong")
	}
}
