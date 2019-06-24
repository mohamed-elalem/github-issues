package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var allowedError = http.StatusUnprocessableEntity

type Issue struct {
	URL           string   `json:"html_url"`
	State         string   `json:"state"`
	CommentsCount int      `json:"comments"`
	CreatedAt     string   `json:"created_at"`
	UpdatedAt     string   `json:"updated_at"`
	Labels        []*Label `json:"labels"`
	RepositoryURL string   `json:"repository_url"`
	Title         string   `json:"title"`
}

type Issues []*Issue

func (i Issues) Do() (Issues, error) {
	resObject := Response{}
	for i := 1; i <= 10; i++ {
		issuesURL := BuildURL(i)
		res, err := http.Get(issuesURL)
		if err != nil {
			return nil, err
		}

		if res.StatusCode == allowedError {
			break
		}

		defer res.Body.Close()

		tmpResObject := Response{}

		if err := json.NewDecoder(res.Body).Decode(&tmpResObject); err != nil {
			return nil, fmt.Errorf("couldn't parse response for url %v: %v", issuesURL, err)
		}

		issues := resObject.Issues
		issues = append(issues, tmpResObject.Issues...)
		resObject.Issues = issues
	}
	return resObject.Issues, nil
}

func BuildURL(page int) string {
	return fmt.Sprintf(`%s?q=%s+state:open+label:"%s"&per_page=100&page=%d`, BaseURL, issuesQuery, issueLabel, page)
}
