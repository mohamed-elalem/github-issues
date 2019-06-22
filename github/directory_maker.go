package github

import (
	"fmt"
	"os"
	"path"
	"strings"
)

var (
	createdDirectories map[string]string
)

type DirectoryMaker struct {
	next Doer
}

func init() {
	createdDirectories = make(map[string]string)
}

func (d *DirectoryMaker) Do() (issues Issues, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("error has occured: %v", err)
		}
	}()

	issues, err = d.next.Do()
	if err != nil {
		return
	}

	if len(issues) == 0 {
		return
	}

	if err = createReposDirectories(issues); err != nil {
		return
	}

	return issues, nil
}

func createReposDirectories(issues Issues) error {
	for _, issue := range issues {
		repoURL := issue.RepositoryURL
		repo, err := getRepoFromCache(repoURL)
		if err != nil {
			return fmt.Errorf("problem occured in cached data: %v", err)
		}
		if err := createDirectoryForRepo(repo); err != nil {
			return fmt.Errorf("Error occured in directory creation: %v", err)
		}
	}

	return nil
}

func createDirectoryForRepo(repository *Repository) error {
	if _, ok := createdDirectories[repository.Name]; ok {
		return nil
	}

	directoryPath := getIssueDirectoryPath(repository.Name)

	if err := createDirectory(directoryPath); err != nil {
		return fmt.Errorf("error occured while creating directory %v: %v", directoryPath, err)
	}

	createdDirectories[repository.Name] = directoryPath

	return nil
}

func getIssueDirectoryPath(repoName string) string {
	directoryBasename := formatPath(repoName)
	directoryPath := path.Join(outputDirectory, directoryBasename)
	return directoryPath
}

func formatPath(path string) string {
	return strings.Replace(path, "/", "-", -1)
}

func createDirectory(path string) error {
	return os.MkdirAll(path, 0755)
}
