package github_releases

import (
	_ "embed"
	"html/template"
	"strings"

	"github.com/selftechio/pigeon/internal/model"
)

var (
	//go:embed template.go.tpl
	templateString string

	// fixme 29/04/2021: hardcoded template name
	tpl = template.Must(template.New("github_content").Parse(templateString))
)

type GithubReleasesTemplateData struct {
	Owner    string
	Name     string
	Releases []model.GithubRepositoryReleasesData
}

// CreateMailBody creates the mail body for a github releases content type.
func CreateMailBody(data *GithubReleasesTemplateData) (string, error) {
	stringBuilder := new(strings.Builder)
	err := tpl.Execute(stringBuilder, data)
	if err != nil {
		return "", err
	}
	return stringBuilder.String(), nil
}
