package github_releases

import (
	"time"

	"github.com/shurcooL/githubv4"
)

// QueryVariables is an interface that defines the behaviour of query variables.
//
// The `ToMap` method is used to create the variables map used for executing
// queries that require variables.
type QueryVariables interface {
	ToMap() map[string]interface{}
}

// RepositoryReleasesTotalCountQuery is a GraphQL query that retrieves the total
// number of releases.
type RepositoryReleasesTotalCountQuery struct {
	Repository struct {
		Releases struct {
			TotalCount int32
		} `graphql:"releases(first: 0)"`
	} `graphql:"repository(owner: $owner, name: $name)"`
}

type RepositoryReleasesTotalCountQueryVariables struct {
	RepoOwner string
	RepoName  string
}

// RepositoryReleasesQuery is a GraphQL query that retrieves the releases of a
// certain repository.
type RepositoryReleasesQuery struct {
	Repository struct {
		Releases struct {
			Edges []struct {
				Node struct {
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
			}
		} `graphql:"releases(orderBy: {field: CREATED_AT, direction: DESC}, first: $first)"`
	} `graphql:"repository(owner: $owner, name: $name)"`
}

type RepositoryReleasesQueryVariables struct {
	RepoOwner string
	RepoName  string
	Count     int
}

func (v *RepositoryReleasesTotalCountQueryVariables) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"owner": githubv4.String(v.RepoOwner),
		"name":  githubv4.String(v.RepoName),
	}
}

func (v *RepositoryReleasesQueryVariables) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"owner": githubv4.String(v.RepoOwner),
		"name":  githubv4.String(v.RepoName),
		"first": githubv4.Int(v.Count),
	}
}
