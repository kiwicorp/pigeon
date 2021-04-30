package github_releases

import (
	"context"

	"github.com/selftechio/pigeon/internal/model"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

// GithubV4ApiClient defines a few abstraction methods that should pe provided
// by an api client.
type GithubV4ApiClient interface {
	// ReleasesTotalCount retrieves the number of releases.
	ReleasesTotalCount(ctx context.Context, owner, name string) (int, error)
	// LatestReleases retrieves a list of releases.
	LatestReleases(ctx context.Context, owner, name string, count int) ([]model.GithubRepositoryReleasesData, error)
}

type v4ApiClient struct {
	inner *githubv4.Client
}

// NewV4ApiClient creates a new GitHub V4 API Client.
func NewGithubV4ApiClient(ctx context.Context, accessToken string) GithubV4ApiClient {
	auth := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	httpClient := oauth2.NewClient(ctx, auth)
	return &v4ApiClient{inner: githubv4.NewClient(httpClient)}
}

func (c *v4ApiClient) ReleasesTotalCount(ctx context.Context, owner, name string) (int, error) {
	variables := map[string]interface{}{
		"owner": githubv4.String(owner),
		"name":  githubv4.String(name),
	}
	qTotalCount := repositoryReleasesTotalCountQuery{}
	err := c.inner.Query(ctx, &qTotalCount, variables)
	if err != nil {
		return -1, err
	}
	return int(qTotalCount.Repository.Releases.TotalCount), nil
}

func (c *v4ApiClient) LatestReleases(ctx context.Context, owner, name string, count int) ([]model.GithubRepositoryReleasesData, error) {
	qReleases := repositoryReleasesQuery{}
	variables := map[string]interface{}{
		"owner": githubv4.String(owner),
		"name":  githubv4.String(name),
		"first": githubv4.Int(count),
	}
	err := c.inner.Query(ctx, &qReleases, variables)
	if err != nil {
		return nil, err
	}
	releases := make([]model.GithubRepositoryReleasesData, len(qReleases.Repository.Releases.Edges))
	for index, release := range qReleases.Repository.Releases.Edges {
		releases[index] = release.Node
	}
	return releases, nil
}
