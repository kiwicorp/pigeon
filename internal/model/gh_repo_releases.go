package model

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	tableName = "content-github-releases"
)

// GithubRepositoryReleases is the model for a repository releases object.
type GithubRepositoryReleases struct {
	Urn        string
	TotalCount int
	Data       []GithubRepositoryReleasesData
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
	DescriptionHtml string
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
