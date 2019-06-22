package github

import (
	"fmt"
	"html/template"
	"path"
)

var (
	issuesIndexTemplate = template.Must(
		template.New("issues.html").
			Funcs(template.FuncMap{
				"IssuePath": formatPath,
			}).
			ParseFiles(path.Join(RootPath, "templates/issues.html")))
)

type IssuesIndex struct {
	next Doer
}

func (i *IssuesIndex) Do() (Issues, error) {
	issues, err := i.next.Do()
	if err != nil {
		return nil, err
	}

	repos := groupIssuesByRepo(issues)

	outputFile, err := createFile(path.Join(outputDirectory, "index.html"))
	if err != nil {
		return nil, fmt.Errorf("error while creating index.html for issues: %v", err)
	}
	defer outputFile.Close()

	if err := issuesIndexTemplate.Execute(outputFile, repos); err != nil {
		return nil, fmt.Errorf("error parsing issues index: %v", err)
	}

	return issues, nil
}
