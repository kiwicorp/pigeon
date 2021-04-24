package content

import (
	"context"
	"fmt"

	"github.com/selftechio/pigeon/internal/content/github_releases"
	"github.com/selftechio/pigeon/internal/db"
	"github.com/selftechio/pigeon/internal/model"
)

type GithubReleasesContentHandler struct {
	accessToken string
	RepoOwner   string
	RepoName    string
}

func NewGithubReleasesContentHandler(accessToken, repoOwner, repoName string) GithubReleasesContentHandler {
	return GithubReleasesContentHandler{
		accessToken: accessToken,
		RepoOwner:   repoOwner,
		RepoName:    repoName,
	}
}

func (ch *GithubReleasesContentHandler) Handle(ctx context.Context) (bool, error) {
	apiClient := github_releases.NewGithubV4ApiClient(ctx, ch.accessToken)

	totalCount, err := apiClient.ReleasesTotalCount(ctx, ch.RepoOwner, ch.RepoName)
	if err != nil {
		return false, err
	}

	item := &model.GithubRepositoryReleases{
		// fixme 24/04/2021: duplicated urn creation code
		Urn: fmt.Sprintf("urn:pigeon-selftech-io:content:github_releases:%s/%s", ch.RepoOwner, ch.RepoName),
	}
	err = db.GetItem(item)
	if err != nil {
		return false, err
	}
	if item.TotalCount == totalCount {
		return false, nil
	}

	releases, err := apiClient.LatestReleases(ctx, ch.RepoOwner, ch.RepoName, totalCount-item.TotalCount)
	if err != nil {
		return true, err
	}

	_, err = db.PutItem(releases)
	if err != nil {
		return true, err
	}

	return true, nil
}
