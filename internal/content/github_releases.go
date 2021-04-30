package content

import (
	"context"
	"fmt"
	"log"

	"github.com/selftechio/pigeon/internal/content/github_releases"
	"github.com/selftechio/pigeon/internal/db"
	"github.com/selftechio/pigeon/internal/mail"
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
	// get item from database
	item := model.NewGithubRepositoryReleases(ch.RepoOwner, ch.RepoName)
	err := db.GetItem(item)
	if err != nil {
		return false, err
	}

	// if the item has no total count, get the total count right now then return
	// fixme 30/04/2021: this will produce no mail on first run.
	//                   a better approach is to initialize the total count when the content urn is
	//                   created in the db or sum stuff
	if item.TotalCount == nil {
		apiClient := github_releases.NewGithubV4ApiClient(ctx, ch.accessToken)
		totalCount, err := apiClient.ReleasesTotalCount(ctx, ch.RepoOwner, ch.RepoName)
		if err != nil {
			return false, err
		}
		item.TotalCount = &totalCount
		_, err = db.PutItem(item)
		if err != nil {
			return false, err
		}
		return false, nil
	}

	apiClient := github_releases.NewGithubV4ApiClient(ctx, ch.accessToken)

	// check github releases count. if the same as before, return
	// otherwise, update current total count
	totalCount, err := apiClient.ReleasesTotalCount(ctx, ch.RepoOwner, ch.RepoName)
	if err != nil {
		return false, err
	}
	if totalCount == *item.TotalCount {
		return false, nil
	}
	item.TotalCount = &totalCount

	// get the diff of releases and cap it at 100, as allowed by the github api
	// fixme 30/04/2021: use multiple requests to grab all releases
	diff := totalCount - *item.TotalCount
	if diff > 100 {
		diff = 100
	}

	// get the new releases
	releases, err := apiClient.LatestReleases(ctx, ch.RepoOwner, ch.RepoName, diff)
	if err != nil {
		return true, err
	}

	// create the mail body
	templateData := github_releases.GithubReleasesTemplateData{
		Owner:    ch.RepoOwner,
		Name:     ch.RepoName,
		Releases: releases,
	}
	body, err := github_releases.CreateMailBody(&templateData)
	if err != nil {
		return true, err
	}

	// mail the new releases
	mailer := mail.NewMailer()
	err = mailer.SendMail([]string{ch.recipient}, fmt.Sprintf("New releases - %s/%s", ch.RepoOwner, ch.RepoName), body)
	if err != nil {
		log.Printf("email error: %#v", err)
		return false, err
	}

	// update the total count
	_, err = db.PutItem(item)
	if err != nil {
		return true, err
	}

	return true, nil
}
