package github

import (
	"fmt"
	"html/template"
	"os"
	"path"
)

var (
	issueTemplate = template.Must(
		template.New("issue.html").
			ParseFiles(path.Join(RootPath, "./templates/issue.html")))
)

type IssueHTML struct {
	next Doer
}

type IssueTemplateData struct {
	RepoName string
	Issues
}

func (i *IssueHTML) Do() (Issues, error) {
	issues, err := i.next.Do()
	if err != nil {
		return nil, err
	}

	repos := groupIssuesByRepo(issues)

	for repoName, issues := range repos {
		err := generateTemplate(repoName, issues)
		if err != nil {
			return nil, err
		}
	}

	return issues, nil
}

func groupIssuesByRepo(issues Issues) map[string]Issues {
	repos := make(map[string]Issues)
	for _, issue := range issues {
		repoName, _ := ExtractRepoName(issue.RepositoryURL)

		issues, ok := repos[repoName]
		if !ok {
			issues = make(Issues, 0)
			repos[repoName] = issues
		}

		repos[repoName] = append(issues, issue)
	}
	return repos
}

func createRepoOutputFile(repoName string) (*os.File, error) {
	filePath := path.Join(getIssueDirectoryPath(repoName), "index.html")
	file, err := createFile(filePath)
	return file, err
}

func generateTemplate(repoName string, issues Issues) error {
	outputFile, err := createRepoOutputFile(repoName)
	if err != nil {
		return fmt.Errorf("error occured while creating file for %v: %v", repoName, err)
	}
	defer outputFile.Close()

	data := IssueTemplateData{RepoName: repoName, Issues: issues}
	if err != nil {
		return fmt.Errorf("Error occured while creating file for %v: %v", repoName, err)
	}
	if err := issueTemplate.Execute(outputFile, data); err != nil {
		return fmt.Errorf("Error occured while generating template: %v", err)
	}

	return nil
}
