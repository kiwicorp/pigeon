package github

import "time"

// // RepositoryReleaseIsLatestQuery is a query for checking whether a release is the latest release,
// // before attempting to make more queries.
// type RepositoryReleaseIsLatestQuery struct {
// 	Repository struct {
// 		Release struct {
// 			IsLatest bool
// 		} `graphql:"release(tagName: $tagName)"`
// 	} `graphql:"repository(owner: $owner, name: $name)"`
// }

// // RepositoryReleasesQuery queries a repository for the latest releases.
// type RepositoryReleasesQuery struct {
// 	Repository struct {
// 		Releases struct {
// 			Nodes struct {
// 				Name   string
// 				Author struct {
// 					Login string
// 					Email string
// 				}
// 				IsDraft      bool
// 				Description  string
// 				IsPrerelease bool
// 				Url          string
// 			}
// 		} `graphql:"orderBy: {field: CREATED_AT, direction: DESC}, first: 10"`
// 	} `graphql:"repository(owner: $owner, name: $name)"`
// }

type RepositoryReleasesTotalCountQuery struct {
	Repository struct {
		Releases struct {
			TotalCount int32
		} `graphql:"releases(first: 0)"`
	} `graphql:"repository(owner: $owner, name: $name)"`
}

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
