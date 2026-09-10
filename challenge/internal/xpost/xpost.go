// Package xpost posts one queued challenge entry per run to X
// (twitter.com/x.com) through the official X API v2.
//
// Everything here goes through the documented API with OAuth 1.0a user
// credentials. That matters beyond correctness: X's automation rules
// draw the line at *unauthorised* automation — browser drivers, scraped
// sessions, anything pretending to be a person at a keyboard — while an
// app posting the account owner's own content under the owner's own
// tokens is the supported path. It is also the only path that gets a
// real status code back instead of a silent shadow-limit.
//
// The body is the entry's POST_X.md, sent verbatim: that file is
// authored to be exactly what lands on the timeline. It is written per
// day rather than derived from POST_DISCORD.md on purpose — near
// identical text fanned out across several networks is the duplicate
// content pattern X's platform manipulation policy actually polices.
package xpost

import (
	"bytes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

// MaxPostChars is the character budget for a single post on the free and
// basic API tiers. X Premium raises the ceiling for long posts, but the
// API path for those is a separate field and a separate approval, so the
// tooling holds every day to the limit that always works.
//
// X counts most characters as one, but weights CJK at two and collapses
// every URL to a fixed 23 characters regardless of real length. Counting
// runes is therefore an approximation — deliberately a conservative one
// for this content, since a post here is ASCII prose plus one long
// GitHub URL that X will shorten to well under what we charged for it.
const MaxPostChars = 280

// httpClient is a package-level client so a slow or hung X API doesn't
// wedge a scheduled run forever. Media upload gets the longer timeout of
// the two because it ships an image.
var httpClient = &http.Client{Timeout: 60 * time.Second}

// API endpoints, as vars rather than consts so tests can point them at a
// httptest server. Nothing outside tests should reassign them.
var (
	tweetEndpoint = "https://api.x.com/2/tweets"
	mediaEndpoint = "https://api.x.com/2/media/upload"
)

// Credentials are the four OAuth 1.0a values from the X developer
// portal: the app's consumer key/secret and the posting account's access
// token/secret. Posting is a user-context action, so a bearer token
// (app-only auth) cannot do it — all four are required.
type Credentials struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
}

// Validate reports which credential fields are missing, naming them so a
// failed scheduled run says which secret to go and set rather than
// surfacing as a 401 from X.
func (c Credentials) Validate() error {
	var missing []string
	if strings.TrimSpace(c.ConsumerKey) == "" {
		missing = append(missing, "X_API_KEY")
	}
	if strings.TrimSpace(c.ConsumerSecret) == "" {
		missing = append(missing, "X_API_SECRET")
	}
	if strings.TrimSpace(c.AccessToken) == "" {
		missing = append(missing, "X_ACCESS_TOKEN")
	}
	if strings.TrimSpace(c.AccessSecret) == "" {
		missing = append(missing, "X_ACCESS_TOKEN_SECRET")
	}
	if len(missing) > 0 {
		return fmt.Errorf("missing X credentials: %s", strings.Join(missing, ", "))
	}
	return nil
}

// ValidatePost checks a post body against X's limits before any network
// call. Over-length content is a bug in the source file, not something
// to silently truncate: a truncated post stops mid-sentence in public.
func ValidatePost(text string) error {
	if strings.TrimSpace(text) == "" {
		return fmt.Errorf("post text is empty")
	}
	if n := utf8.RuneCountInString(text); n > MaxPostChars {
		return fmt.Errorf("post is %d characters, over X's %d limit", n, MaxPostChars)
	}
	return nil
}

// Post publishes text to X, attaching the image at heroPath if one is
// given. An empty heroPath posts text only. It returns the new post's
// ID so the caller can log a link to what actually went out.
func Post(creds Credentials, text, heroPath string) (postID string, err error) {
	if err := creds.Validate(); err != nil {
		return "", err
	}
	if err := ValidatePost(text); err != nil {
		return "", err
	}

	body := map[string]any{"text": text}
	if heroPath != "" {
		mediaID, err := UploadMedia(creds, heroPath)
		if err != nil {
			return "", fmt.Errorf("uploading hero image: %w", err)
		}
		body["media"] = map[string]any{"media_ids": []string{mediaID}}
	}

	payload, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("encoding payload: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, tweetEndpoint, bytes.NewReader(payload))
	if err != nil {
		return "", fmt.Errorf("building request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	// A JSON body contributes nothing to an OAuth 1.0a signature — the
	// base string covers the query string and the oauth_* parameters
	// only — so the body is signed by omission, not by hashing.
	if err := sign(req, creds); err != nil {
		return "", err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("posting to X: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(io.LimitReader(resp.Body, 8192))
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", apiError(resp, respBody)
	}

	var out struct {
		Data struct {
			ID string `json:"id"`
		} `json:"data"`
	}
	if err := json.Unmarshal(respBody, &out); err != nil {
		return "", fmt.Errorf("X accepted the post but its response did not parse: %w", err)
	}
	return out.Data.ID, nil
}

// UploadMedia uploads an image and returns the media ID to attach to a
// post. It uses the v2 simple upload, which takes the whole file in one
// multipart request — the hero images are a few hundred KB, far under
// the point where chunked upload becomes necessary.
func UploadMedia(creds Credentials, path string) (string, error) {
	imageBytes, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("reading %s: %w", path, err)
	}

	var buf bytes.Buffer
	form := multipart.NewWriter(&buf)
	part, err := form.CreateFormFile("media", filepath.Base(path))
	if err != nil {
		return "", fmt.Errorf("creating media part: %w", err)
	}
	if _, err := part.Write(imageBytes); err != nil {
		return "", fmt.Errorf("writing media: %w", err)
	}
	// Without a category X stores the upload but won't let a post
	// reference it, and the failure surfaces later as an opaque error on
	// the post itself rather than here.
	if err := form.WriteField("media_category", "tweet_image"); err != nil {
		return "", fmt.Errorf("writing media_category: %w", err)
	}
	if err := form.Close(); err != nil {
		return "", fmt.Errorf("closing multipart form: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, mediaEndpoint, &buf)
	if err != nil {
		return "", fmt.Errorf("building upload request: %w", err)
	}
	req.Header.Set("Content-Type", form.FormDataContentType())
	// Multipart bodies are excluded from the signature base string for
	// the same reason JSON ones are: only query and oauth_* parameters
	// are signed.
	if err := sign(req, creds); err != nil {
		return "", err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("uploading media to X: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(io.LimitReader(resp.Body, 8192))
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", apiError(resp, respBody)
	}

	// v2 returns the id under data; the older v1.1 shape put
	// media_id_string at the top level. Both are accepted because this
	// endpoint moved once already and an id in the wrong place should
	// not cost a day's post.
	var out struct {
		Data struct {
			ID string `json:"id"`
		} `json:"data"`
		MediaIDString string `json:"media_id_string"`
	}
	if err := json.Unmarshal(respBody, &out); err != nil {
		return "", fmt.Errorf("X accepted the upload but its response did not parse: %w", err)
	}
	if out.Data.ID != "" {
		return out.Data.ID, nil
	}
	if out.MediaIDString != "" {
		return out.MediaIDString, nil
	}
	return "", fmt.Errorf("X accepted the upload but returned no media id: %s", strings.TrimSpace(string(respBody)))
}

// apiError turns a non-2xx response into an error worth reading in a
// workflow log. The rate limit case is called out separately because it
// is the one failure that is about pacing rather than a broken request,
// and the reset header says how long to wait.
func apiError(resp *http.Response, body []byte) error {
	trimmed := strings.TrimSpace(string(body))
	switch resp.StatusCode {
	case http.StatusTooManyRequests:
		reset := resp.Header.Get("x-rate-limit-reset")
		if ts, err := strconv.ParseInt(reset, 10, 64); err == nil {
			return fmt.Errorf("X rate limited this app until %s: %s",
				time.Unix(ts, 0).UTC().Format(time.RFC3339), trimmed)
		}
		return fmt.Errorf("X rate limited this app: %s", trimmed)
	case http.StatusUnauthorized:
		return fmt.Errorf("X rejected the credentials (401) — check the four OAuth values and that the app has Read and write permission: %s", trimmed)
	case http.StatusForbidden:
		return fmt.Errorf("X refused the request (403) — usually a duplicate post, a suspended app, or an access tier without write access: %s", trimmed)
	default:
		return fmt.Errorf("X returned %d: %s", resp.StatusCode, trimmed)
	}
}

// sign adds an OAuth 1.0a Authorization header to req, per RFC 5849.
func sign(req *http.Request, creds Credentials) error {
	nonce, err := newNonce()
	if err != nil {
		return err
	}
	return signWith(req, creds, nonce, time.Now().Unix())
}

// signWith is sign with the nonce and timestamp injected, so a test can
// check the signature against a fixed expected value.
func signWith(req *http.Request, creds Credentials, nonce string, timestamp int64) error {
	oauthParams := map[string]string{
		"oauth_consumer_key":     creds.ConsumerKey,
		"oauth_nonce":            nonce,
		"oauth_signature_method": "HMAC-SHA1",
		"oauth_timestamp":        strconv.FormatInt(timestamp, 10),
		"oauth_token":            creds.AccessToken,
		"oauth_version":          "1.0",
	}

	// The base string covers the oauth_* parameters and the query
	// string, sorted together as one set.
	signingParams := make(map[string]string, len(oauthParams))
	for k, v := range oauthParams {
		signingParams[k] = v
	}
	for k, vs := range req.URL.Query() {
		if len(vs) > 0 {
			signingParams[k] = vs[0]
		}
	}

	baseURL := (&url.URL{Scheme: req.URL.Scheme, Host: req.URL.Host, Path: req.URL.Path}).String()
	base := strings.ToUpper(req.Method) + "&" + encode(baseURL) + "&" + encode(joinParams(signingParams))

	key := encode(creds.ConsumerSecret) + "&" + encode(creds.AccessSecret)
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(base))
	oauthParams["oauth_signature"] = base64.StdEncoding.EncodeToString(mac.Sum(nil))

	names := make([]string, 0, len(oauthParams))
	for k := range oauthParams {
		names = append(names, k)
	}
	sort.Strings(names)
	pairs := make([]string, 0, len(names))
	for _, k := range names {
		pairs = append(pairs, fmt.Sprintf("%s=%q", encode(k), encode(oauthParams[k])))
	}
	req.Header.Set("Authorization", "OAuth "+strings.Join(pairs, ", "))
	return nil
}

// joinParams renders the normalised parameter string: every pair
// percent-encoded, sorted by encoded key, joined with &.
func joinParams(params map[string]string) string {
	encoded := make([]string, 0, len(params))
	for k, v := range params {
		encoded = append(encoded, encode(k)+"="+encode(v))
	}
	sort.Strings(encoded)
	return strings.Join(encoded, "&")
}

// encode is RFC 3986 percent-encoding. url.QueryEscape is not a
// substitute: it renders a space as + and leaves ~ alone, and both
// differences change the signature.
func encode(s string) string {
	const unreserved = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-._~"
	var b strings.Builder
	for _, c := range []byte(s) {
		if strings.IndexByte(unreserved, c) >= 0 {
			b.WriteByte(c)
			continue
		}
		fmt.Fprintf(&b, "%%%02X", c)
	}
	return b.String()
}

func newNonce() (string, error) {
	buf := make([]byte, 16)
	if _, err := rand.Read(buf); err != nil {
		return "", fmt.Errorf("generating OAuth nonce: %w", err)
	}
	return hex.EncodeToString(buf), nil
}
