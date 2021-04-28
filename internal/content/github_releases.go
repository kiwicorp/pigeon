package content

import (
	"context"

	"github.com/selftechio/pigeon/internal/content/github_releases"
	"github.com/selftechio/pigeon/internal/db"
	"github.com/selftechio/pigeon/internal/model"
)

type GithubReleasesContentHandler struct {
	accessToken string
	RepoOwner   string
	RepoName    string
	recipient   string
}

func NewGithubReleasesContentHandler(accessToken, repoOwner, repoName, recipient string) GithubReleasesContentHandler {
	return GithubReleasesContentHandler{
		accessToken: accessToken,
		RepoOwner:   repoOwner,
		RepoName:    repoName,
		recipient:   recipient,
	}
}

func (ch *GithubReleasesContentHandler) Handle(ctx context.Context) (bool, error) {
	apiClient := github_releases.NewGithubV4ApiClient(ctx, ch.accessToken)

	totalCount, err := apiClient.ReleasesTotalCount(ctx, ch.RepoOwner, ch.RepoName)
	if err != nil {
		return false, err
	}

	item := model.NewGithubRepositoryReleases(ch.RepoOwner, ch.RepoName)
	err = db.GetItem(item)
	if err != nil {
		return false, err
	}
	if item.TotalCount == totalCount {
		return false, nil
	}

	// fixme 24/04/2021: some situations may lead to a request that asks for more than 100 items,
	// and that's prohibited by the github v4 api
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
