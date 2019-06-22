package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Issue struct {
	URL           string   `json:"html_url"`
	State         string   `json:"state"`
	CommentsCount int      `json:"comments"`
	CreatedAt     string   `json:"created_at"`
	UpdatedAt     string   `json:"updated_at"`
	Labels        []*Label `json:"labels"`
	RepositoryURL string   `json:"repository_url"`
}

type Issues []*Issue

func (i Issues) Do() (Issues, error) {
	issuesURL := BuildURL()
	res, err := http.Get(issuesURL)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resObject := Response{}

	if err := json.NewDecoder(res.Body).Decode(&resObject); err != nil {
		return nil, fmt.Errorf("couldn't parse response for url %v: %v", issuesURL, err)
	}

	return resObject.Issues, nil
}

func BuildURL() string {
	return fmt.Sprintf(`%s?q=%s+state:open+label:"%s"`, BaseURL, issuesQuery, issueLabel)
}
