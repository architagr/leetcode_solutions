package leetcode

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// mockGraphQL serves the two-call submission flow: the submission list
// first, then the code for whichever id was asked for.
func mockGraphQL(t *testing.T, submissions string, codeByID map[string]string) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			Query     string         `json:"query"`
			Variables map[string]any `json:"variables"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Errorf("bad request body: %v", err)
		}
		if c, err := r.Cookie("LEETCODE_SESSION"); err != nil || c.Value != "sess-abc" {
			t.Errorf("request did not carry the session cookie: %v", err)
		}
		if got := r.Header.Get("x-csrftoken"); got != "csrf-xyz" {
			t.Errorf("x-csrftoken = %q, want csrf-xyz", got)
		}
		w.Header().Set("Content-Type", "application/json")
		if _, ok := body.Variables["submissionId"]; ok {
			id := json.Number("")
			switch v := body.Variables["submissionId"].(type) {
			case float64:
				id = json.Number(jsonItoa(int(v)))
			}
			code := codeByID[string(id)]
			w.Write([]byte(`{"data":{"submissionDetails":{"code":` + jsonQuote(code) + `,"lang":{"name":"golang"}}}}`))
			return
		}
		w.Write([]byte(`{"data":{"questionSubmissionList":{"submissions":` + submissions + `}}}`))
	}))
}

func jsonQuote(s string) string { b, _ := json.Marshal(s); return string(b) }
func jsonItoa(i int) string     { b, _ := json.Marshal(i); return string(b) }

func authedClient(url string) *Client {
	c := NewClient()
	c.BaseURL = url
	c.Session = &Session{LeetCodeSession: "sess-abc", CSRFToken: "csrf-xyz"}
	return c
}

func TestFetchAcceptedSubmissionPrefersGo(t *testing.T) {
	list := `[
	  {"id":"31","statusDisplay":"Wrong Answer","lang":"golang","timestamp":"1770000300"},
	  {"id":"30","statusDisplay":"Accepted","lang":"python3","timestamp":"1770000200"},
	  {"id":"29","statusDisplay":"Accepted","lang":"golang","timestamp":"1770000100"}
	]`
	srv := mockGraphQL(t, list, map[string]string{"29": "func twoSum() {}", "30": "def twoSum(): pass"})
	defer srv.Close()

	sub, found, err := authedClient(srv.URL).FetchAcceptedSubmission("two-sum", "golang")
	if err != nil || !found {
		t.Fatalf("found=%v err=%v", found, err)
	}
	// The Go one is older than the accepted python3 one; preferLang must win.
	if sub.ID != "29" || sub.Lang != "golang" {
		t.Fatalf("got id=%s lang=%s, want 29/golang", sub.ID, sub.Lang)
	}
	if sub.Code != "func twoSum() {}" {
		t.Fatalf("code = %q", sub.Code)
	}
	if got := sub.Timestamp.Format("2006-01-02"); got != "2026-02-02" {
		t.Fatalf("timestamp = %s", got)
	}
}

func TestFetchAcceptedSubmissionFallsBackToOtherLang(t *testing.T) {
	list := `[{"id":"30","statusDisplay":"Accepted","lang":"python3","timestamp":"1770000200"}]`
	srv := mockGraphQL(t, list, map[string]string{"30": "def twoSum(): pass"})
	defer srv.Close()

	sub, found, err := authedClient(srv.URL).FetchAcceptedSubmission("two-sum", "golang")
	if err != nil || !found {
		t.Fatalf("found=%v err=%v", found, err)
	}
	if sub.Lang != "python3" {
		t.Fatalf("lang = %s, want python3", sub.Lang)
	}
}

func TestFetchAcceptedSubmissionNotSolved(t *testing.T) {
	srv := mockGraphQL(t, `[{"id":"1","statusDisplay":"Wrong Answer","lang":"golang","timestamp":"1770000200"}]`, nil)
	defer srv.Close()

	_, found, err := authedClient(srv.URL).FetchAcceptedSubmission("two-sum", "golang")
	if err != nil {
		t.Fatalf("err = %v", err)
	}
	if found {
		t.Fatal("found = true, want false when nothing is accepted")
	}
}

func TestFetchAcceptedSubmissionNeedsAuth(t *testing.T) {
	c := NewClient()
	if _, _, err := c.FetchAcceptedSubmission("two-sum", "golang"); err == nil {
		t.Fatal("want an error from an unauthenticated client")
	}
}
