package leetcode

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Session holds the browser cookies that authenticate a LeetCode user.
// Both values are copied out of a logged-in browser's cookie jar for
// leetcode.com; they are secrets and must never be committed.
type Session struct {
	// LeetCodeSession is the "LEETCODE_SESSION" cookie.
	LeetCodeSession string `json:"leetcodeSession"`
	// CSRFToken is the "csrftoken" cookie, echoed back in the
	// x-csrftoken header on every request.
	CSRFToken string `json:"csrfToken"`
}

// LoadSession reads a Session from a JSON file. An empty path (or a
// path that does not exist) falls back to the LEETCODE_SESSION and
// LEETCODE_CSRF environment variables, so a cookie can be supplied
// without ever touching the filesystem.
func LoadSession(path string) (Session, error) {
	if path != "" {
		raw, err := os.ReadFile(path)
		if err == nil {
			var s Session
			if err := json.Unmarshal(raw, &s); err != nil {
				return Session{}, fmt.Errorf("parsing session file %s: %w", path, err)
			}
			if err := s.validate(path); err != nil {
				return Session{}, err
			}
			return s, nil
		}
		if !os.IsNotExist(err) {
			return Session{}, fmt.Errorf("reading session file %s: %w", path, err)
		}
	}

	s := Session{
		LeetCodeSession: strings.TrimSpace(os.Getenv("LEETCODE_SESSION")),
		CSRFToken:       strings.TrimSpace(os.Getenv("LEETCODE_CSRF")),
	}
	if err := s.validate("$LEETCODE_SESSION/$LEETCODE_CSRF"); err != nil {
		return Session{}, err
	}
	return s, nil
}

func (s Session) validate(origin string) error {
	if strings.TrimSpace(s.LeetCodeSession) == "" {
		return fmt.Errorf("no LeetCode session cookie in %s: log in at leetcode.com, copy the LEETCODE_SESSION and csrftoken cookies", origin)
	}
	return nil
}
