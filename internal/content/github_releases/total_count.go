package github_releases

// repositoryReleasesTotalCountQuery is a GraphQL query that retrieves the total
// number of releases.
type repositoryReleasesTotalCountQuery struct {
	Repository struct {
		Releases struct {
			TotalCount int32
		} `graphql:"releases(first: 0)"`
	} `graphql:"repository(owner: $owner, name: $name)"`
}
