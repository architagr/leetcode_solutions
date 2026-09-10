package xpost

import (
	"encoding/json"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"unicode/utf8"
)

func testCreds() Credentials {
	return Credentials{ConsumerKey: "ck", ConsumerSecret: "cs", AccessToken: "at", AccessSecret: "as"}
}

// redirect points the package at a test server for the duration of a
// test, restoring the real endpoints afterwards.
func redirect(t *testing.T, tweetURL, mediaURL string) {
	t.Helper()
	oldTweet, oldMedia := tweetEndpoint, mediaEndpoint
	tweetEndpoint, mediaEndpoint = tweetURL, mediaURL
	t.Cleanup(func() { tweetEndpoint, mediaEndpoint = oldTweet, oldMedia })
}

func writeHero(t *testing.T) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "HERO.png")
	if err := os.WriteFile(path, []byte("not-really-a-png"), 0o644); err != nil {
		t.Fatal(err)
	}
	return path
}

// The signature is the one part of this package that fails silently and
// expensively: a wrong one is a 401 from a scheduled run at 06:00. The
// expected value here was produced by an independent HMAC-SHA1
// implementation, so a Go-side change to the base string construction
// shows up as a mismatch rather than as agreement with itself.
func TestSignWithMatchesAnIndependentlyComputedSignature(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "https://api.x.com/2/tweets", strings.NewReader(`{"text":"hi"}`))
	if err != nil {
		t.Fatal(err)
	}
	if err := signWith(req, testCreds(), "abc123", 1600000000); err != nil {
		t.Fatal(err)
	}

	got := req.Header.Get("Authorization")
	const want = `oauth_signature="mYcpossuO%2FUyWMTN3Ci7vpAnDeQ%3D"`
	if !strings.Contains(got, want) {
		t.Errorf("Authorization = %s\nwant it to contain %s", got, want)
	}
	if !strings.HasPrefix(got, "OAuth ") {
		t.Errorf("Authorization = %q, want it to start with %q", got, "OAuth ")
	}
	for _, field := range []string{`oauth_consumer_key="ck"`, `oauth_token="at"`, `oauth_nonce="abc123"`, `oauth_timestamp="1600000000"`, `oauth_signature_method="HMAC-SHA1"`, `oauth_version="1.0"`} {
		if !strings.Contains(got, field) {
			t.Errorf("Authorization is missing %s: %s", field, got)
		}
	}
}

// A JSON or multipart body is not part of an OAuth 1.0a base string, so
// two requests differing only in body must sign identically. Getting
// this wrong is the classic cause of intermittent 401s.
func TestSignWithIgnoresTheRequestBody(t *testing.T) {
	sigFor := func(body string) string {
		req, err := http.NewRequest(http.MethodPost, "https://api.x.com/2/tweets", strings.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}
		if err := signWith(req, testCreds(), "abc123", 1600000000); err != nil {
			t.Fatal(err)
		}
		return req.Header.Get("Authorization")
	}
	if a, b := sigFor(`{"text":"one"}`), sigFor(`{"text":"a completely different body"}`); a != b {
		t.Errorf("signature changed with the body:\n%s\n%s", a, b)
	}
}

func TestSignWithIncludesQueryParameters(t *testing.T) {
	plain, err := http.NewRequest(http.MethodPost, "https://api.x.com/2/tweets", nil)
	if err != nil {
		t.Fatal(err)
	}
	withQuery, err := http.NewRequest(http.MethodPost, "https://api.x.com/2/tweets?expansions=author_id", nil)
	if err != nil {
		t.Fatal(err)
	}
	if err := signWith(plain, testCreds(), "abc123", 1600000000); err != nil {
		t.Fatal(err)
	}
	if err := signWith(withQuery, testCreds(), "abc123", 1600000000); err != nil {
		t.Fatal(err)
	}
	if plain.Header.Get("Authorization") == withQuery.Header.Get("Authorization") {
		t.Error("query parameters did not affect the signature, but they are part of the base string")
	}
}

func TestEncodeIsRFC3986NotQueryEscape(t *testing.T) {
	// url.QueryEscape would render the space as + and escape the tilde.
	if got, want := encode("a b~c"), "a%20b~c"; got != want {
		t.Errorf("encode(%q) = %q, want %q", "a b~c", got, want)
	}
}

// A link costs 23 whatever its length, and the deep links these posts
// carry run past 120 characters — counting those in full would spend
// half the budget on something X charges 23 for.
func TestPostLengthChargesLinksTheTcoWeight(t *testing.T) {
	const link = "https://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/binary_tree_preorder_traversal/SOLUTION.md"
	if len(link) <= URLChars {
		t.Fatal("test precondition: the link should be longer than the weight X charges for it")
	}

	if got, want := PostLength(link), URLChars; got != want {
		t.Errorf("PostLength(a bare link) = %d, want %d", got, want)
	}
	if got, want := PostLength("see "+link), len("see ")+URLChars; got != want {
		t.Errorf("PostLength(prose + link) = %d, want %d", got, want)
	}
	if got, want := PostLength(link+" and "+link), 2*URLChars+len(" and "); got != want {
		t.Errorf("PostLength(two links) = %d, want %d", got, want)
	}
	if got, want := PostLength("no links here"), 13; got != want {
		t.Errorf("PostLength(no links) = %d, want %d", got, want)
	}
}

// The post that motivated this: real prose plus a real deep link, over
// the limit counted naively and comfortably inside it counted as X does.
func TestValidatePostAcceptsAPostThatOnlyFitsWithLinkWeighting(t *testing.T) {
	post := "Day 8/365 · Binary Tree Preorder Traversal (Easy)\n\n" +
		strings.Repeat("a", 130) +
		"\n\nhttps://github.com/architagr/leetcode_solutions/blob/main/easy_problems/101_200/binary_tree_preorder_traversal/SOLUTION.md"

	if utf8.RuneCountInString(post) <= MaxPostChars {
		t.Fatal("test precondition: the post should be over the limit when counted naively")
	}
	if err := ValidatePost(post); err != nil {
		t.Errorf("a post that fits once links are weighted must be accepted, got %v", err)
	}
}

func TestValidatePostRejectsEmptyAndOverLength(t *testing.T) {
	if err := ValidatePost("   \n "); err == nil {
		t.Error("an all-whitespace post must be rejected")
	}
	if err := ValidatePost(strings.Repeat("a", MaxPostChars)); err != nil {
		t.Errorf("a post at exactly the limit must be accepted, got %v", err)
	}
	if err := ValidatePost(strings.Repeat("a", MaxPostChars+1)); err == nil {
		t.Error("a post one character over the limit must be rejected")
	}
}

// Length is counted in runes, not bytes: an em dash is three bytes and
// one character, and counting bytes would reject posts X would accept.
func TestValidatePostCountsRunesNotBytes(t *testing.T) {
	post := strings.Repeat("—", MaxPostChars)
	if len(post) <= MaxPostChars {
		t.Fatal("test precondition: the post should be longer in bytes than in runes")
	}
	if err := ValidatePost(post); err != nil {
		t.Errorf("a 280-rune post must be accepted, got %v", err)
	}
}

func TestCredentialsValidateNamesEveryMissingVariable(t *testing.T) {
	err := Credentials{}.Validate()
	if err == nil {
		t.Fatal("empty credentials must be rejected")
	}
	for _, name := range []string{"X_API_KEY", "X_API_SECRET", "X_ACCESS_TOKEN", "X_ACCESS_TOKEN_SECRET"} {
		if !strings.Contains(err.Error(), name) {
			t.Errorf("error does not name %s: %v", name, err)
		}
	}
	if err := testCreds().Validate(); err != nil {
		t.Errorf("complete credentials must be accepted, got %v", err)
	}
}

func TestPostSendsTextAndReturnsTheNewPostID(t *testing.T) {
	var gotBody map[string]any
	var gotAuth string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
		json.NewDecoder(r.Body).Decode(&gotBody)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"data":{"id":"1799999999","text":"ok"}}`)
	}))
	defer srv.Close()
	redirect(t, srv.URL, srv.URL)

	id, err := Post(testCreds(), "Day 9 · a post", "")
	if err != nil {
		t.Fatal(err)
	}
	if id != "1799999999" {
		t.Errorf("post id = %q, want 1799999999", id)
	}
	if gotBody["text"] != "Day 9 · a post" {
		t.Errorf("text = %v, want the post body verbatim", gotBody["text"])
	}
	if _, ok := gotBody["media"]; ok {
		t.Error("a text-only post must not send a media field")
	}
	if !strings.HasPrefix(gotAuth, "OAuth ") {
		t.Errorf("Authorization = %q, want an OAuth 1.0a header", gotAuth)
	}
}

func TestPostUploadsTheHeroImageAndAttachesIt(t *testing.T) {
	var uploadedName, uploadedCategory string
	var gotBody map[string]any

	media := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, params, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
		if err != nil {
			t.Errorf("upload content type: %v", err)
			return
		}
		mr := multipart.NewReader(r.Body, params["boundary"])
		for {
			part, err := mr.NextPart()
			if err != nil {
				break
			}
			switch part.FormName() {
			case "media":
				uploadedName = part.FileName()
			case "media_category":
				b, _ := io.ReadAll(part)
				uploadedCategory = string(b)
			}
		}
		io.WriteString(w, `{"data":{"id":"media-42"}}`)
	}))
	defer media.Close()

	tweets := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewDecoder(r.Body).Decode(&gotBody)
		io.WriteString(w, `{"data":{"id":"1800000000"}}`)
	}))
	defer tweets.Close()
	redirect(t, tweets.URL, media.URL)

	if _, err := Post(testCreds(), "with a hero", writeHero(t)); err != nil {
		t.Fatal(err)
	}
	if uploadedName != "HERO.png" {
		t.Errorf("uploaded file name = %q, want HERO.png", uploadedName)
	}
	if uploadedCategory != "tweet_image" {
		t.Errorf("media_category = %q, want tweet_image", uploadedCategory)
	}
	mediaField, ok := gotBody["media"].(map[string]any)
	if !ok {
		t.Fatalf("post body has no media field: %v", gotBody)
	}
	ids, ok := mediaField["media_ids"].([]any)
	if !ok || len(ids) != 1 || ids[0] != "media-42" {
		t.Errorf("media_ids = %v, want [media-42]", mediaField["media_ids"])
	}
}

// The upload endpoint moved from v1.1 to v2 and the id changed place in
// the response. Accepting both shapes means a further shuffle does not
// cost a day's post.
func TestUploadMediaAcceptsTheLegacyResponseShape(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `{"media_id_string":"legacy-7","media_id":7}`)
	}))
	defer srv.Close()
	redirect(t, srv.URL, srv.URL)

	id, err := UploadMedia(testCreds(), writeHero(t))
	if err != nil {
		t.Fatal(err)
	}
	if id != "legacy-7" {
		t.Errorf("media id = %q, want legacy-7", id)
	}
}

func TestUploadMediaErrorsWhenNoIDComesBack(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `{"data":{}}`)
	}))
	defer srv.Close()
	redirect(t, srv.URL, srv.URL)

	if _, err := UploadMedia(testCreds(), writeHero(t)); err == nil {
		t.Error("an upload that returns no media id must be an error")
	}
}

// A day whose post fails must not be marked, so every failure has to
// come back as an error rather than an empty id and a nil.
func TestPostSurfacesAPIErrors(t *testing.T) {
	cases := []struct {
		name       string
		status     int
		headers    map[string]string
		body       string
		wantSubstr string
	}{
		{"rate limited", http.StatusTooManyRequests, map[string]string{"x-rate-limit-reset": "1600000000"}, `{"title":"Too Many Requests"}`, "2020-09-13T12:26:40Z"},
		{"bad credentials", http.StatusUnauthorized, nil, `{"title":"Unauthorized"}`, "401"},
		{"duplicate post", http.StatusForbidden, nil, `{"detail":"duplicate content"}`, "duplicate"},
		{"server error", http.StatusInternalServerError, nil, `boom`, "500"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				for k, v := range tc.headers {
					w.Header().Set(k, v)
				}
				w.WriteHeader(tc.status)
				io.WriteString(w, tc.body)
			}))
			defer srv.Close()
			redirect(t, srv.URL, srv.URL)

			_, err := Post(testCreds(), "a post", "")
			if err == nil {
				t.Fatalf("status %d must be an error", tc.status)
			}
			if !strings.Contains(err.Error(), tc.wantSubstr) {
				t.Errorf("error = %v, want it to mention %q", err, tc.wantSubstr)
			}
		})
	}
}

// Validation runs before any network call, so a bad post never burns an
// API write or reaches the timeline half-formed.
func TestPostValidatesBeforeCallingTheAPI(t *testing.T) {
	called := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
	}))
	defer srv.Close()
	redirect(t, srv.URL, srv.URL)

	if _, err := Post(testCreds(), strings.Repeat("a", MaxPostChars+1), ""); err == nil {
		t.Error("an over-length post must be rejected")
	}
	if _, err := Post(Credentials{}, "fine", ""); err == nil {
		t.Error("missing credentials must be rejected")
	}
	if called {
		t.Error("the API was called despite validation failing")
	}
}

func TestPostMissingHeroFileIsAnError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `{"data":{"id":"1"}}`)
	}))
	defer srv.Close()
	redirect(t, srv.URL, srv.URL)

	if _, err := Post(testCreds(), "msg", filepath.Join(t.TempDir(), "nope.png")); err == nil {
		t.Fatal("a missing hero file must be an error, not a silent skip")
	}
}

func redirectMe(t *testing.T, meURL string) {
	t.Helper()
	old := meEndpoint
	meEndpoint = meURL
	t.Cleanup(func() { meEndpoint = old })
}

func TestVerifyReportsTheAuthenticatedAccount(t *testing.T) {
	var method, auth string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method, auth = r.Method, r.Header.Get("Authorization")
		io.WriteString(w, `{"data":{"id":"42","username":"architagr","name":"Archit Agarwal"}}`)
	}))
	defer srv.Close()
	redirectMe(t, srv.URL)

	acct, err := Verify(testCreds())
	if err != nil {
		t.Fatal(err)
	}
	if acct.Username != "architagr" || acct.ID != "42" || acct.Name != "Archit Agarwal" {
		t.Errorf("account = %+v, want the authenticated user", acct)
	}
	// Verify must never publish anything.
	if method != http.MethodGet {
		t.Errorf("method = %s, want GET — Verify must not write", method)
	}
	if !strings.HasPrefix(auth, "OAuth ") {
		t.Errorf("Authorization = %q, want an OAuth 1.0a header", auth)
	}
}

func TestVerifySurfacesBadCredentials(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, `{"title":"Unauthorized"}`)
	}))
	defer srv.Close()
	redirectMe(t, srv.URL)

	_, err := Verify(testCreds())
	if err == nil {
		t.Fatal("a 401 must be an error")
	}
	if !strings.Contains(err.Error(), "Read and write") {
		t.Errorf("a 401 during setup should point at the permission trap, got %v", err)
	}
}

func TestVerifyChecksCredentialsBeforeCallingTheAPI(t *testing.T) {
	called := false
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
	}))
	defer srv.Close()
	redirectMe(t, srv.URL)

	if _, err := Verify(Credentials{ConsumerKey: "ck"}); err == nil {
		t.Error("incomplete credentials must be rejected")
	}
	if called {
		t.Error("the API was called with incomplete credentials")
	}
}

// The oauth1-permissions 403 is the one setup failure that looks like
// every other one — the credentials authenticate, verify-x passes, and
// only the write is refused. The error has to say what to do about it,
// because the fix (regenerate the tokens) is not what the message
// "check your permissions" would lead you to do.
func TestPostExplainsTheReadOnlyTokenTrap(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		io.WriteString(w, `{"detail":"Your client app is not configured with the appropriate oauth1 app permissions for this endpoint.","status":403,"title":"Forbidden","type":"https://api.x.com/2/problems/oauth1-permissions"}`)
	}))
	defer srv.Close()
	redirect(t, srv.URL, srv.URL)

	_, err := Post(testCreds(), "a post", "")
	if err == nil {
		t.Fatal("a 403 must be an error")
	}
	if !strings.Contains(err.Error(), "REGENERATE") {
		t.Errorf("the error must say to regenerate the tokens, got %v", err)
	}
}
