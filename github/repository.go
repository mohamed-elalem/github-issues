package github

import (
	"fmt"
	"net/url"
	"strings"
)

type Repository struct {
	Name string
}

var (
	reposCache      map[string]*Repository
	repoIssuesCount map[string]int
)

func init() {
	reposCache = make(map[string]*Repository)
	repoIssuesCount = make(map[string]int)
}

func cacheRepo(repoURL string) (*Repository, bool, error) {
	repoName, err := ExtractRepoName(repoURL)
	if err != nil {
		return nil, false, fmt.Errorf("couldn't get repo name: %v", err)
	}
	repository, ok := reposCache[repoName]
	if !ok {
		repository = &Repository{repoName}
		reposCache[repoURL] = repository
	}
	return repository, !ok, nil
}

func getRepoFromCache(repoURL string) (*Repository, error) {
	repo, ok := reposCache[repoURL]
	if ok {
		return repo, nil
	}

	repo, _, err := cacheRepo(repoURL)

	if err != nil {
		return nil, err
	}

	return repo, nil
}

func ExtractRepoName(repoURL string) (string, error) {
	parsedURL, err := url.Parse(repoURL)
	if err != nil {
		return "", err
	}
	urlBasename := strings.TrimSpace(parsedURL.EscapedPath())
	urlPath := strings.Split(urlBasename[1:], "/")
	repoName := strings.Join(urlPath[1:], "_")
	return repoName, nil
}
