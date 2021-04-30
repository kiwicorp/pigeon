package model

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	// fixme 26/04/2021: hardcoded table name
	tableName = "content-github-releases"
)

// GithubRepositoryReleases is the model for a repository releases object.
type GithubRepositoryReleases struct {
	Urn        string
	TotalCount *int
}

// GithubRepositoryReleasesData represents the data that changed since the last
// check.
type GithubRepositoryReleasesData struct {
	Author struct {
		Email string
		Login string
	}
	CreatedAt       time.Time
	Description     string
	DescriptionHtml string `graphql:"descriptionHTML"`
	IsDraft         bool
	IsPrerelease    bool
	Name            string
	PublishedAt     time.Time
	Url             string
	TagName         string
	IsLatest        bool
}

func (m *GithubRepositoryReleases) TableName() *string {
	return aws.String(tableName)
}

func (m *GithubRepositoryReleases) Key() map[string]*dynamodb.AttributeValue {
	return map[string]*dynamodb.AttributeValue{
		// fixme 24/04/2021: hardcoded hash key
		"Urn": {
			S: aws.String(m.Urn),
		},
	}
}

// NewGithubRepositoryReleases creates a new github repository releases object
// and assigns it a urn.
func NewGithubRepositoryReleases(owner, name string) *GithubRepositoryReleases {
	return &GithubRepositoryReleases{
		Urn: fmt.Sprintf("urn:pigeon-selftech-io:content:github_releases:%s/%s", owner, name),
	}
}
