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
	ReleasesTotalCount(ctx context.Context, owner, name string) (int, error)
	LatestReleases(ctx context.Context, owner, name string, count int) (*model.GithubRepositoryReleases, error)
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

func (c *v4ApiClient) LatestReleases(ctx context.Context, owner, name string, count int) (*model.GithubRepositoryReleases, error) {
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

	releases := model.NewGithubRepositoryReleases(owner, name)
	releases.TotalCount = int(qReleases.Repository.Releases.TotalCount)

	releases_data := make([]model.GithubRepositoryReleasesData, 0, len(qReleases.Repository.Releases.Edges))
	for _, release := range qReleases.Repository.Releases.Edges {
		releases_data = append(releases_data, model.GithubRepositoryReleasesData{
			Author: struct {
				Email string
				Login string
			}{
				Email: release.Node.Author.Email,
				Login: release.Node.Author.Login,
			},
			CreatedAt:       release.Node.CreatedAt,
			Description:     release.Node.Description,
			DescriptionHtml: release.Node.DescriptionHtml,
			IsDraft:         release.Node.IsDraft,
			IsPrerelease:    release.Node.IsPrerelease,
			Name:            release.Node.Name,
			PublishedAt:     release.Node.PublishedAt,
			Url:             release.Node.Url,
			TagName:         release.Node.TagName,
			IsLatest:        release.Node.IsLatest,
		})
	}

	releases.Data = releases_data
	return releases, nil
}
