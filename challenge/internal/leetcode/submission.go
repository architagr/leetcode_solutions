package leetcode

import (
	"fmt"
	"strconv"
	"time"
)

// Submission is one accepted submission of the authenticated user's own
// code for a question.
type Submission struct {
	ID        string
	Lang      string // LeetCode's lang slug, e.g. "golang"
	Code      string
	Timestamp time.Time
}

// FetchAcceptedSubmission returns the authenticated user's most recent
// accepted submission for a question, preferring one written in
// preferLang (a LeetCode lang slug such as "golang") and falling back to
// the newest accepted submission in any language.
//
// found is false when the user has no accepted submission for the
// question at all — i.e. they have not solved it.
func (c *Client) FetchAcceptedSubmission(slug, preferLang string) (sub Submission, found bool, err error) {
	if c.Session == nil {
		return Submission{}, false, fmt.Errorf("FetchAcceptedSubmission needs an authenticated client")
	}

	const listQuery = `query submissionList($offset: Int!, $limit: Int!, $questionSlug: String!) {
	  questionSubmissionList(offset: $offset, limit: $limit, questionSlug: $questionSlug) {
	    submissions {
	      id
	      statusDisplay
	      lang
	      timestamp
	    }
	  }
	}`

	var listResp struct {
		QuestionSubmissionList struct {
			Submissions []struct {
				ID            string `json:"id"`
				StatusDisplay string `json:"statusDisplay"`
				Lang          string `json:"lang"`
				Timestamp     string `json:"timestamp"`
			} `json:"submissions"`
		} `json:"questionSubmissionList"`
	}
	vars := map[string]any{"offset": 0, "limit": 20, "questionSlug": slug}
	if err := c.doGraphQL(listQuery, vars, &listResp); err != nil {
		return Submission{}, false, err
	}

	// Submissions come back newest-first, so the first accepted match in
	// each pass is already the most recent one.
	var chosen *struct {
		ID            string `json:"id"`
		StatusDisplay string `json:"statusDisplay"`
		Lang          string `json:"lang"`
		Timestamp     string `json:"timestamp"`
	}
	for i := range listResp.QuestionSubmissionList.Submissions {
		s := &listResp.QuestionSubmissionList.Submissions[i]
		if s.StatusDisplay != "Accepted" {
			continue
		}
		if preferLang != "" && s.Lang == preferLang {
			chosen = s
			break
		}
		if chosen == nil {
			chosen = s
		}
	}
	if chosen == nil {
		return Submission{}, false, nil
	}

	const detailQuery = `query submissionDetails($submissionId: Int!) {
	  submissionDetails(submissionId: $submissionId) {
	    code
	    lang { name }
	  }
	}`
	id, err := strconv.Atoi(chosen.ID)
	if err != nil {
		return Submission{}, false, fmt.Errorf("unexpected submission id %q: %w", chosen.ID, err)
	}
	var detail struct {
		SubmissionDetails struct {
			Code string `json:"code"`
			Lang struct {
				Name string `json:"name"`
			} `json:"lang"`
		} `json:"submissionDetails"`
	}
	if err := c.doGraphQL(detailQuery, map[string]any{"submissionId": id}, &detail); err != nil {
		return Submission{}, false, err
	}
	if detail.SubmissionDetails.Code == "" {
		return Submission{}, false, fmt.Errorf("submission %s for %s came back with no code", chosen.ID, slug)
	}

	out := Submission{
		ID:   chosen.ID,
		Lang: chosen.Lang,
		Code: detail.SubmissionDetails.Code,
	}
	if out.Lang == "" {
		out.Lang = detail.SubmissionDetails.Lang.Name
	}
	if secs, err := strconv.ParseInt(chosen.Timestamp, 10, 64); err == nil {
		out.Timestamp = time.Unix(secs, 0).UTC()
	}
	return out, true, nil
}
