package github_releases

import (
	"time"
)

// repositoryReleasesQuery is a GraphQL query that retrieves the releases of a
// certain repository.
type repositoryReleasesQuery struct {
	Repository struct {
		Releases struct {
			TotalCount int32
			Edges      []struct {
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
