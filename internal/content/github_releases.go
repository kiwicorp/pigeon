package content

import (
	"context"
	"log"

	"github.com/selftechio/pigeon/internal/content/github_releases"
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

	prevTotalCount := 0 // fixme 23/04/2021: get value from database
	if prevTotalCount == totalCount {
		return false, nil
	}

	releases, err := apiClient.LatestReleases(ctx, ch.RepoOwner, ch.RepoName, totalCount-prevTotalCount)
	if err != nil {
		return true, err
	}

	// todo 23/04/2021: save the changes in the database
	log.Printf("Latest releases: %#v", releases)

	return true, nil
}
