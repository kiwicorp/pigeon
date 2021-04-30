package github_releases

import (
	"github.com/selftechio/pigeon/internal/model"
)

// repositoryReleasesQuery is a GraphQL query that retrieves the releases of a
// certain repository.
type repositoryReleasesQuery struct {
	Repository struct {
		Releases struct {
			Edges []struct {
				Node model.GithubRepositoryReleasesData
			}
		} `graphql:"releases(orderBy: {field: CREATED_AT, direction: DESC}, first: $first)"`
	} `graphql:"repository(owner: $owner, name: $name)"`
}

// repositoryReleasesTotalCountQuery is a GraphQL query that retrieves the total
// number of releases.
type repositoryReleasesTotalCountQuery struct {
	Repository struct {
		Releases struct {
			TotalCount int32
		} `graphql:"releases(first: 0)"`
	} `graphql:"repository(owner: $owner, name: $name)"`
}
