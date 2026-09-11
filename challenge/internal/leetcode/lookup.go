package leetcode

import (
	"strconv"
	"strings"
)

// Meta is the identity of a question: enough to place its folder in the
// canonical layout.
type Meta struct {
	Number     int
	Title      string
	Slug       string
	Difficulty string
}

// FetchQuestionMeta looks a question up by its title slug and returns
// its frontend number and difficulty. found is false when LeetCode has
// no question at that slug, which is the normal outcome for a folder
// name that is not quite the real slug.
//
// No authentication required.
func (c *Client) FetchQuestionMeta(slug string) (Meta, bool, error) {
	query := `query questionMeta($titleSlug: String!) {
	  question(titleSlug: $titleSlug) {
	    questionFrontendId
	    title
	    titleSlug
	    difficulty
	  }
	}`

	var resp struct {
		Question *struct {
			QuestionFrontendID string `json:"questionFrontendId"`
			Title              string `json:"title"`
			TitleSlug          string `json:"titleSlug"`
			Difficulty         string `json:"difficulty"`
		} `json:"question"`
	}
	if err := c.doGraphQL(query, map[string]any{"titleSlug": slug}, &resp); err != nil {
		// A slug that does not exist comes back as a GraphQL error on
		// some LeetCode deployments and as a null question on others.
		if strings.Contains(strings.ToLower(err.Error()), "no question") {
			return Meta{}, false, nil
		}
		return Meta{}, false, err
	}
	if resp.Question == nil || resp.Question.TitleSlug == "" {
		return Meta{}, false, nil
	}
	n, err := strconv.Atoi(resp.Question.QuestionFrontendID)
	if err != nil {
		return Meta{}, false, nil
	}
	return Meta{
		Number:     n,
		Title:      resp.Question.Title,
		Slug:       resp.Question.TitleSlug,
		Difficulty: strings.ToLower(resp.Question.Difficulty),
	}, true, nil
}
