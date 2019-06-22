package github

import "fmt"

var (
	issueLabel      string
	outputDirectory string
	issuesQuery     string
)

func Run(query, label, outputDir string) error {
	issuesQuery = query
	issueLabel = label
	outputDirectory = outputDir

	issues := make(Issues, 0)
	wrapper := &IssuesIndex{
		next: &IssueHTML{
			next: &DirectoryMaker{
				next: issues,
			},
		},
	}
	issues, err := wrapper.Do()
	if err != nil {
		return fmt.Errorf("problem occured: %v", err)
	}
	return nil
}
