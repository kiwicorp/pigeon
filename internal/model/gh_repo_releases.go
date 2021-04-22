package model

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// GithubRepositoryReleases is the database model for repository releases
// objects.
type GithubRepositoryReleases struct {
	Urn string

	TotalCount int32
	Author     struct {
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
	return aws.String("content-github-releases")
}

func (m *GithubRepositoryReleases) AvMap() (map[string]*dynamodb.AttributeValue, error) {
	return dynamodbattribute.MarshalMap(m)
}

func (m *GithubRepositoryReleases) Key() map[string]*dynamodb.AttributeValue {
	return map[string]*dynamodb.AttributeValue{
		"Urn": {
			S: aws.String(m.Urn),
		},
	}
}

func (m *GithubRepositoryReleases) ExpressionAv() map[string]*dynamodb.AttributeValue {
	return nil
}

func (m *GithubRepositoryReleases) ReturnValues() string {
	return ""
}

func (m *GithubRepositoryReleases) UpdateExpression() string {
	return ""
}
